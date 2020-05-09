package exchange

import (
	"math/big"
	"testing"
)

func TestKeyExchange(t *testing.T) {
	group, _ := GetGroup(14)
	p1 := newPeer(group)
	p2 := newPeer(group)

	err := exchangeKey(p1, p2)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestCustomGroupKeyExchange(t *testing.T) {
	p, _ := new(big.Int).SetString("FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A63A36210000000000090563", 16)
	g := new(big.Int).SetInt64(2)
	group := CreateGroup(p, g)
	p1 := newPeer(group)
	p2 := newPeer(group)

	err := exchangeKey(p1, p2)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestPIsNotMutable(t *testing.T) {
	d, _ := GetGroup(0)
	p := d.p.String()
	d.P().Set(big.NewInt(1))
	if p != d.p.String() {
		t.Errorf("group's prime mutated externally, should be %s, was changed to %s", p, d.p.String())
	}
}

func TestGIsNotMutable(t *testing.T) {
	d, _ := GetGroup(0)
	g := d.g.String()
	d.G().Set(big.NewInt(0))
	if g != d.g.String() {
		t.Errorf("group's generator mutated externally, should be %s, was changed to %s", g, d.g.String())
	}
}
