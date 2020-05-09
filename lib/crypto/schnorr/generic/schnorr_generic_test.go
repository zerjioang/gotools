package generic

import (
	"math/big"
	"testing"

	"github.com/zerjioang/gotools/lib/crypto/util"
)

func TestSchnorr(t *testing.T) {
	t.Run("create-schnorr-group", func(t *testing.T) {
		p, _ := util.RandomPrime(50)
		q, _ := util.RandomPrime(50)
		// r = (p-1)/q
		one := big.NewInt(1)
		p1 := big.NewInt(0).Sub(p, one)
		r := big.NewInt(0).Div(p1, q)
		t.Log(p, q, r)
	})
}
