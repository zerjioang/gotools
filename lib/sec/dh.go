package sec

import (
	"fmt"
	"math/big"
)

//check if DF param is safe or not
func isSafePrime(modulusInt *big.Int) bool {
	// q2 = p - 1
	b1 := new(big.Int)
	fmt.Sscan("1", b1)

	var q2 = new(big.Int)

	q2.Sub(modulusInt, b1)

	//fmt.Println("p-1", q2.String())

	// q2 % b2 == 0?
	b2 := new(big.Int)
	fmt.Sscan("2", b2)
	b0 := new(big.Int)
	fmt.Sscan("0", b0)
	mod := new(big.Int)

	if b0.Cmp(mod.Mod(q2, b2)) != 0 {
		return false
	}

	// q2 / 2 prime?
	q2.Div(q2, b2)

	if q2.ProbablyPrime(500) {
		return true
	} else {
		return false
	}

}
