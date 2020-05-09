package secure

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

const (
	charset string = "abcdefghijklmnopqrstuvwxyzABCDEFHFGHIJKLMNOPQRSTUVWXYZ1234567890"
)

func Keygen256() string {
	b := strings.Builder{}
	size := 256 / 8
	for key := 0; key < size; key++ {
		res, _ := rand.Int(rand.Reader, big.NewInt(64))
		keyGen := charset[res.Int64()]
		stringGen := fmt.Sprintf("%c", keyGen)
		b.WriteString(stringGen)
	}
	return b.String()
}
