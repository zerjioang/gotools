package exchange

import (
	"math/big"
)

type DHKey struct {
	x *big.Int
	y *big.Int

	group *DHGroup
}

func (k *DHKey) Bytes() []byte {
	if k.y == nil {
		return nil
	}
	if k.group != nil {
		// len = ceil(bitLen(y) / 8)
		blen := (k.group.p.BitLen() + 7) / 8
		ret := make([]byte, blen)
		copyWithLeftPad(ret, k.y.Bytes())
		return ret
	}
	return k.y.Bytes()
}

func (k *DHKey) String() string {
	if k.y == nil {
		return ""
	}
	return k.y.String()
}

func (k *DHKey) IsPrivateKey() bool {
	return k.x != nil
}

func NewPublicKey(s []byte) *DHKey {
	key := new(DHKey)
	key.y = new(big.Int).SetBytes(s)
	return key
}

// copyWithLeftPad copies src to the end of dest, padding with zero bytes as
// needed.
func copyWithLeftPad(dest, src []byte) {
	numPaddingBytes := len(dest) - len(src)
	for i := 0; i < numPaddingBytes; i++ {
		dest[i] = 0
	}
	copy(dest[numPaddingBytes:], src)
}
