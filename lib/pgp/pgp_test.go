// +build ignore

package pgp

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	"crypto"
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/sha256"

	_ "golang.org/x/crypto/ripemd160"

	"compress/gzip"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

var (
	// Goencrypt app
	bits          = 4096
	privateKey    = ""
	publicKey     = ""
	signatureFile = ""

	// Generates new public and private keys
	keyGenCmd       = "keygen"
	keyOutputPrefix = "prefix"
	keyOutputDir    = "d"

	// Encrypts a file with a public key
	encryptionCmd = "encrypt"

	// Signs a file with a private key
	signCmd = "sign"

	// Verifies a file was signed with the public key
	verifyCmd = "verify"

	// Decrypts a file with a private key
	decryptionCmd = "decrypt"
)

func encodePrivateKey(out io.Writer, key *rsa.PrivateKey) error {
	w, err := armor.Encode(out, openpgp.PrivateKeyType, make(map[string]string))
	if err != nil {
		// kingpin.FatalIfError(err, "Error creating OpenPGP Armor: %s", err)
		return err
	}

	pgpKey := packet.NewRSAPrivateKey(time.Now(), key)
	sErr := pgpKey.Serialize(w)
	if sErr != nil {
		//"Error serializing private key: %s", err)
		return sErr
	}
	clErr := w.Close()
	if clErr != nil {
		// "Error serializing private key: %s", err)
		return clErr
	}
	return nil
}

func decodePrivateKey(filename string) (*packet.PrivateKey, error) {

	// open ascii armored private key
	in, oErr := os.Open(filename)
	if oErr != nil {
		// "Error opening private key: %s", err
		return nil, oErr
	}
	defer in.Close()

	block, dErr := armor.Decode(in)
	if dErr != nil {
		// dErr, "Error decoding OpenPGP Armor: %s", err)
		return nil, dErr
	}

	if block.Type != openpgp.PrivateKeyType {
		return nil, errors.New("invalid private key file - error decoding private key")
	}

	reader := packet.NewReader(block.Body)
	pkt, rErr := reader.Next()
	if rErr != nil {
		// "error reading private key"
		return nil, rErr
	}

	key, ok := pkt.(*packet.PrivateKey)
	if !ok {
		return nil, errors.New("invalid private key, error parsing private key")
	}
	return key, nil
}

func encodePublicKey(out io.Writer, key *rsa.PrivateKey) error {
	w, eErr := armor.Encode(out, openpgp.PublicKeyType, make(map[string]string))
	if eErr != nil {
		// "Error creating OpenPGP Armor: %s", err
		return eErr
	}

	pgpKey := packet.NewRSAPublicKey(time.Now(), &key.PublicKey)
	sErr := pgpKey.Serialize(w)
	if sErr != nil {
		// "Error serializing public key: %s", sErr
		return sErr
	}
	clErr := w.Close()
	if clErr != nil {
		// "Error serializing public key: %s", err
		return clErr
	}
	return nil
}

func decodePublicKey(filename string) *packet.PublicKey {

	// open ascii armored public key
	in, err := os.Open(filename)
	kingpin.FatalIfError(err, "Error opening public key: %s", err)
	defer in.Close()

	block, err := armor.Decode(in)
	kingpin.FatalIfError(err, "Error decoding OpenPGP Armor: %s", err)

	if block.Type != openpgp.PublicKeyType {
		kingpin.FatalIfError(errors.New("Invalid private key file"), "Error decoding private key")
	}

	reader := packet.NewReader(block.Body)
	pkt, err := reader.Next()
	kingpin.FatalIfError(err, "Error reading private key")

	key, ok := pkt.(*packet.PublicKey)
	if !ok {
		kingpin.FatalIfError(errors.New("Invalid public key"), "Error parsing public key")
	}
	return key
}

func decodeSignature(filename string) *packet.Signature {

	// open ascii armored public key
	in, err := os.Open(filename)
	kingpin.FatalIfError(err, "Error opening public key: %s", err)
	defer in.Close()

	block, err := armor.Decode(in)
	kingpin.FatalIfError(err, "Error decoding OpenPGP Armor: %s", err)

	if block.Type != openpgp.SignatureType {
		kingpin.FatalIfError(errors.New("Invalid signature file"), "Error decoding signature")
	}

	reader := packet.NewReader(block.Body)
	pkt, err := reader.Next()
	kingpin.FatalIfError(err, "Error reading signature")

	sig, ok := pkt.(*packet.Signature)
	if !ok {
		kingpin.FatalIfError(errors.New("Invalid signature"), "Error parsing signature")
	}
	return sig
}

func encryptFile() {
	pubKey := decodePublicKey(publicKey)
	privKey := decodePrivateKey(privateKey)

	to := createEntityFromKeys(pubKey, privKey)

	w, err := armor.Encode(os.Stdout, "Message", make(map[string]string))
	kingpin.FatalIfError(err, "Error creating OpenPGP Armor: %s", err)
	defer w.Close()

	plain, err := openpgp.Encrypt(w, []*openpgp.Entity{to}, nil, nil, nil)
	kingpin.FatalIfError(err, "Error creating entity for encryption")
	defer plain.Close()

	compressed, err := gzip.NewWriterLevel(plain, gzip.BestCompression)
	kingpin.FatalIfError(err, "Invalid compression level")

	n, err := io.Copy(compressed, os.Stdin)
	kingpin.FatalIfError(err, "Error writing encrypted file")
	kingpin.Errorf("Encrypted %d bytes", n)

	compressed.Close()
}

func decryptFile() {
	pubKey := decodePublicKey(publicKey)
	privKey := decodePrivateKey(privateKey)

	entity := createEntityFromKeys(pubKey, privKey)

	block, err := armor.Decode(os.Stdin)
	kingpin.FatalIfError(err, "Error reading OpenPGP Armor: %s", err)

	if block.Type != "Message" {
		kingpin.FatalIfError(err, "Invalid message type")
	}

	var entityList openpgp.EntityList
	entityList = append(entityList, entity)

	md, err := openpgp.ReadMessage(block.Body, entityList, nil, nil)
	kingpin.FatalIfError(err, "Error reading message")

	compressed, err := gzip.NewReader(md.UnverifiedBody)
	kingpin.FatalIfError(err, "Invalid compression level")
	defer compressed.Close()

	n, err := io.Copy(os.Stdout, compressed)
	kingpin.FatalIfError(err, "Error reading encrypted file")
	kingpin.Errorf("Decrypted %d bytes", n)
}

func signFile() {
	pubKey := decodePublicKey(publicKey)
	privKey := decodePrivateKey(privateKey)

	signer := createEntityFromKeys(pubKey, privKey)

	err := openpgp.ArmoredDetachSign(os.Stdout, signer, os.Stdin, nil)
	kingpin.FatalIfError(err, "Error signing input")
}

func verifyFile() {
	pubKey := decodePublicKey(publicKey)
	sig := decodeSignature(signatureFile)

	hash := sig.Hash.New()
	io.Copy(hash, os.Stdin)

	err := pubKey.VerifySignature(hash, sig)
	kingpin.FatalIfError(err, "Error signing input")
	kingpin.Errorf("Verified signature")
}

func createEntityFromKeys(pubKey *packet.PublicKey, privKey *packet.PrivateKey) *openpgp.Entity {
	config := packet.Config{
		DefaultHash:            crypto.SHA256,
		DefaultCipher:          packet.CipherAES256,
		DefaultCompressionAlgo: packet.CompressionZLIB,
		CompressionConfig: &packet.CompressionConfig{
			Level: 9,
		},
		RSABits: bits,
	}
	currentTime := config.Now()
	uid := packet.NewUserId("", "", "")

	e := openpgp.Entity{
		PrimaryKey: pubKey,
		PrivateKey: privKey,
		Identities: make(map[string]*openpgp.Identity),
	}
	isPrimaryId := false

	e.Identities[uid.Id] = &openpgp.Identity{
		Name:   uid.Name,
		UserId: uid,
		SelfSignature: &packet.Signature{
			CreationTime: currentTime,
			SigType:      packet.SigTypePositiveCert,
			PubKeyAlgo:   packet.PubKeyAlgoRSA,
			Hash:         config.Hash(),
			IsPrimaryId:  &isPrimaryId,
			FlagsValid:   true,
			FlagSign:     true,
			FlagCertify:  true,
			IssuerKeyId:  &e.PrimaryKey.KeyId,
		},
	}

	keyLifetimeSecs := uint32(86400 * 365)

	e.Subkeys = make([]openpgp.Subkey, 1)
	e.Subkeys[0] = openpgp.Subkey{
		PublicKey:  pubKey,
		PrivateKey: privKey,
		Sig: &packet.Signature{
			CreationTime:              currentTime,
			SigType:                   packet.SigTypeSubkeyBinding,
			PubKeyAlgo:                packet.PubKeyAlgoRSA,
			Hash:                      config.Hash(),
			PreferredHash:             []uint8{8}, // SHA-256
			FlagsValid:                true,
			FlagEncryptStorage:        true,
			FlagEncryptCommunications: true,
			IssuerKeyId:               &e.PrimaryKey.KeyId,
			KeyLifetimeSecs:           &keyLifetimeSecs,
		},
	}
	return &e
}

func generateKeys() error {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		// Error generating RSA key: %s"
		return err
	}

	priv, err := os.Create(filepath.Join(keyOutputDir, keyOutputPrefix+".privkey"))
	if err != nil {
		// Error writing private key to file
		return err
	}
	defer priv.Close()

	pub, err := os.Create(filepath.Join(keyOutputDir, keyOutputPrefix+".pubkey"))
	if err != nil {
		// Error writing public key to file
		return err
	}
	defer pub.Close()

	_ = encodePrivateKey(priv, key)
	_ = encodePublicKey(pub, key)
	return nil
}

func TestPgp(t *testing.T) {
	t.Run("generate", func(t *testing.T) {
		generateKeys()
	})
	t.Run("encrypt", func(t *testing.T) {
		encryptFile()
	})
	t.Run("sign", func(t *testing.T) {
		signFile()
	})
	t.Run("verify", func(t *testing.T) {
		verifyFile()
	})
	t.Run("decrypt", func(t *testing.T) {
		decryptFile()
	})
}
