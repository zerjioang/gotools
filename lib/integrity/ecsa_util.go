// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package integrity

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/hex"
	"math/big"

	"github.com/zerjioang/gotools/lib/logger"
)

const (
	// private key json marshaled
	keyData = `{
		"Curve": {
			"P": 26959946667150639794667015087019630673557916260026308143510066298881,
			"N": 26959946667150639794667015087019625940457807714424391721682722368061,
			"B": 18958286285566608000408668544493926415504680968679321075787234672564,
			"Gx": 19277929113566293071110308034699488026831934219452440156649784352033,
			"Gy": 19926808758034470970197974370888749184205991990603949537637343198772,
			"BitSize": 224,
			"name": "P-224"
		},
		"X": 9924104907269953551211951017586279157096504725630884119809217880060,
		"Y": 2305888103526496946328786469202700929478437529547546609149473741759,
		"D": 4901073595975219053801381639479282785126445966884851835499842583479
	}`
)

var (
	zero = big.NewInt(0)
	// sha256 hash
	h = sha256.New()
	//decode private key
	integrityPrivKey *ecdsa.PrivateKey
	//private integrity key bytes
	//privateBytes = []byte(constants.IntegrityPrivateKeyPem)
	//publicBytes  = []byte(constants.IntegrityPublicKeyPem)
)

func newBig(value string) *big.Int {
	n := new(big.Int)
	n, ok := n.SetString(value, 10)
	if !ok {
		logger.Error("failed to set big int value")
	}
	return n
}

func init() {
	//integrityPrivKey, _ = decode(privateBytes, publicBytes)
	//raw, err := json.Marshal(integrityPrivKey)
	//logger.Debug(string(raw), err)
	/*
		{
			"Curve": {
				"P": 26959946667150639794667015087019630673557916260026308143510066298881,
				"N": 26959946667150639794667015087019625940457807714424391721682722368061,
				"B": 18958286285566608000408668544493926415504680968679321075787234672564,
				"Gx": 19277929113566293071110308034699488026831934219452440156649784352033,
				"Gy": 19926808758034470970197974370888749184205991990603949537637343198772,
				"BitSize": 224,
				"name": "P-224"
			},
			"X": 9924104907269953551211951017586279157096504725630884119809217880060,
			"Y": 2305888103526496946328786469202700929478437529547546609149473741759,
			"D": 4901073595975219053801381639479282785126445966884851835499842583479
		}
	*/
	integrityPrivKey = new(ecdsa.PrivateKey)
	integrityPrivKey.Curve = elliptic.P224()
	integrityPrivKey.Curve.Params().P = newBig("26959946667150639794667015087019630673557916260026308143510066298881")
	integrityPrivKey.Curve.Params().N = newBig("26959946667150639794667015087019625940457807714424391721682722368061")
	integrityPrivKey.Curve.Params().B = newBig("18958286285566608000408668544493926415504680968679321075787234672564")
	integrityPrivKey.Curve.Params().Gx = newBig("19277929113566293071110308034699488026831934219452440156649784352033")
	integrityPrivKey.Curve.Params().Gy = newBig("19926808758034470970197974370888749184205991990603949537637343198772")
	integrityPrivKey.X = newBig("9924104907269953551211951017586279157096504725630884119809217880060")
	integrityPrivKey.Y = newBig("2305888103526496946328786469202700929478437529547546609149473741759")
	integrityPrivKey.D = newBig("4901073595975219053801381639479282785126445966884851835499842583479")
}

func SignMsgWithIntegrity(message string) (string, string) {
	//create test message
	//str := util.UnsafeBytes(message)
	// hash test message
	h.Reset()
	h.Write([]byte(message))
	signhash := h.Sum(nil)
	hexhash := hex.EncodeToString(signhash)
	r, s, _ := ecdsaSign(signhash, integrityPrivKey)
	signature := PointsToDER(r, s)
	return hexhash, signature
}

// create a ecdsa signature
func ecdsaSign(message []byte, priv *ecdsa.PrivateKey) (r, s *big.Int, err error) {
	r = zero
	s = zero
	r, s, err = ecdsa.Sign(rand.Reader, priv, message)
	return r, s, err
}

// verify given ecdsa signature
func ecdsaVerify(hash []byte, r *big.Int, s *big.Int, pub *ecdsa.PublicKey) bool {
	return ecdsa.Verify(pub, hash, r, s)
}

// Convert an ECDSA signature (points R and S) to a byte array using ASN.1 DER encoding.
func PointsToDER(r, s *big.Int) string {
	type ecdsaWrapper struct {
		R, S *big.Int
	}
	sequence := ecdsaWrapper{r, s}
	encoding, _ := asn1.Marshal(sequence)
	return hex.EncodeToString(encoding)
}
