package pgp

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

func TestAsciiArmor(t *testing.T) {
	encryptionPassphrase := []byte("golang")
	encryptionText := "Hello world. Encryption and Decryption testing.\n"
	message := []byte(encryptionText)

	encryptionType := "PGP SIGNATURE"

	encbuf := bytes.NewBuffer(nil)
	w, err := armor.Encode(encbuf, encryptionType, nil)
	if err != nil {
		log.Fatal(err)
	}

	hints := &openpgp.FileHints{
		IsBinary: true,
	}

	packetConfig := &packet.Config{
		DefaultCipher: packet.CipherAES256,
	}

	// entities := []*openpgp.Entity{}
	//plaintext, err := openpgp.Encrypt(w, entities, nil, hints, packetConfig)
	plaintext, err := openpgp.SymmetricallyEncrypt(w, encryptionPassphrase, hints, packetConfig)
	if err != nil {
		log.Fatal(err)
	}
	_, err = plaintext.Write(message)

	defer plaintext.Close()
	defer w.Close()

	fmt.Printf("Encrypted:\n%s\n", encbuf)

	decbuf := bytes.NewBuffer([]byte(encbuf.String()))
	result, err := armor.Decode(decbuf)
	if err != nil {
		log.Fatal(err)
	}

	alreadyPrompted := false
	md, err := openpgp.ReadMessage(result.Body, nil, func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		// from openpgp docs: https://godoc.org/golang.org/x/crypto/openpgp#PromptFunction:
		// If the decrypted private key or given passphrase isn't correct, the function will be called again, forever.
		if alreadyPrompted {
			return nil, errors.New("could not decrypt data using supplied passphrase")
		} else {
			alreadyPrompted = true
		}
		return encryptionPassphrase, nil
	}, nil)

	if err != nil {
		log.Fatal("Could not decrypt data: ", err)
	}

	decryptedBytes, err := ioutil.ReadAll(md.UnverifiedBody)
	fmt.Printf("Decrypted:\n%s\n", string(decryptedBytes))
}
