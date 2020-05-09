package pearson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPearsonHasher(t *testing.T) {
	tests := map[string]struct {
		salt             []byte
		message          []byte
		expectedHashDo   []byte
		expectedHashSalt []byte
	}{
		"Pearson 0x00": {[]byte{0xff}, []byte{0x00}, []byte{0xd2}, []byte{0xeb}},
	}

	hasher := NewPearsonHasher()
	for testname, test := range tests {
		hashSalt := hasher.Salted(test.salt, test.message)
		hashDo := hasher.Do(test.message)
		assert.Equalf(t, test.expectedHashDo, hashDo, "Hash Do don't match in test: %s", testname)
		assert.Equalf(t, test.expectedHashSalt, hashSalt, "Hash Salt don't match in test: %s", testname)
		assert.NotEqual(t, hashDo, hashSalt, "Do and Salted hashes should NOT match in test: %s", testname)
	}
}
