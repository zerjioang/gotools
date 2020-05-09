package jwt

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/zerjioang/gotools/util/str"
)

const (
	open      = "{"
	close     = "}"
	separator = ","
	equal     = "="
	dot       = "."
)

var (
	openraw      = []byte(open)
	closeraw     = []byte(close)
	separatorraw = []byte(separator)
	equalraw     = []byte(equal)
	dotraw       = []byte(dot)
)
var (
	defaultTokenHeaders = map[string]interface{}{
		"typ": "JWT",
		"alg": "HS256",
	}
	//default token header as bytes
	defaultTokenHeaderBytes = []byte(`{"alg":"HS256","typ":"JWT"}`)
	//token header encoded as base64
	b64TokenHeader    = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
	b64TokenHeaderRaw = []byte(b64TokenHeader)
)

type LowClaim struct {
	Key   string
	Value interface{}
}

func (c LowClaim) Json() []byte {
	// calculate json value
	var jsonValue string
	switch c.Value.(type) {
	case int:
		jsonValue = strconv.FormatInt(int64(c.Value.(int)), 10)
	case int64:
		jsonValue = strconv.FormatInt(c.Value.(int64), 10)
	case string:
		jsonValue = c.Value.(string)
	default:
		jsonValue = "undefined"
	}
	content := `"` + c.Key + `": "` + jsonValue + `"`
	return str.UnsafeBytes(content)
}

type LowClaims []LowClaim

func (c LowClaims) Json() []byte {
	var builder bytes.Buffer
	//builder := strings.Builder{}
	builder.Write(openraw)
	if len(c) > 0 {
		builder.Write(c[0].Json())
		if len(c) > 1 {
			builder.Write(separatorraw)
			_ = c[len(c)-1]
			for i := 1; i < len(c); i++ {
				builder.Write(c[i].Json())
			}
		}
	}
	builder.Write(closeraw)
	return builder.Bytes()
}

func GenerateHS256Jwt(claims LowClaims, key []byte) (string, error) {
	// Sign and get the complete encoded token as a string using the secret
	// Generate the signing string.  This is the
	// most expensive part of the whole deal.  Unless you
	// need this for something special, just go straight for
	// the SignedString.
	if key == nil || len(key) == 0 {
		return "", errors.New("invalid key provided")
	}
	//var parts = [2]string{b64TokenHeader, ""}
	// a, _ := json.Marshal(defaultTokenHeaders)
	b := claims.Json()

	// bytes buffer
	//define a byte buffer for string concatenations
	var buffer bytes.Buffer
	//encoder
	var encoder = base64.URLEncoding

	// parts[0] = EncodeSegment(defaultTokenHeaderBytes)
	// Encode segment
	buf := make([]byte, encoder.EncodedLen(len(b)))
	encoder.Encode(buf, b)
	raw := str.UnsafeString(buf)
	//parts[1] = strings.TrimRight(raw, "=")
	claimsRaw := raw[:strings.Index(raw, "=")]

	// concatenate strings into a message
	buffer.Write(b64TokenHeaderRaw)
	buffer.Write(dotraw)
	buffer.Write(str.UnsafeBytes(claimsRaw))
	message := buffer.Bytes()

	// hasher
	hasher := hmac.New(crypto.SHA256.New, key)
	_, _ = hasher.Write(message)
	hashedResult := hasher.Sum(nil)
	// Encode JWT specific base64url encoding with padding stripped
	// Encode segment
	buf = make([]byte, encoder.EncodedLen(len(hashedResult)))
	encoder.Encode(buf, hashedResult)

	//signature := strings.TrimRight(raw, "=")
	equalIdx := bytes.Index(buf, equalraw)
	// check if equalIdx is last character
	// if it is, we make a hack and replace the equal with a dot
	// in this way, we save an array copy and array write to the buffer
	//trim the content
	signature := buf[:equalIdx]
	// return full JWT using buffer for concats
	buffer.Reset()
	buffer.Write(message)
	buffer.Write(dotraw)
	buffer.Write(signature)
	token := buffer.String()
	return token, nil
}
