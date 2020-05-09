package exchange

import "fmt"

type peer struct {
	priv  *DHKey
	group *DHGroup
	pub   *DHKey
}

func newPeer(g *DHGroup) *peer {
	ret := new(peer)
	ret.priv, _ = g.GeneratePrivateKey(nil)
	ret.group = g
	return ret
}

func (p *peer) getPubKey() []byte {
	return p.priv.Bytes()
}

func (p *peer) recvPeerPubKey(pub []byte) {
	pubKey := NewPublicKey(pub)
	p.pub = pubKey
}

func (p *peer) getKey() []byte {
	k, err := p.group.ComputeKey(p.pub, p.priv)
	if err != nil {
		return nil
	}
	return k.Bytes()
}

func exchangeKey(p1, p2 *peer) error {
	pub1 := p1.getPubKey()
	pub2 := p2.getPubKey()

	p1.recvPeerPubKey(pub2)
	p2.recvPeerPubKey(pub1)

	key1 := p1.getKey()
	key2 := p2.getKey()

	if key1 == nil {
		return fmt.Errorf("p1 has nil key")
	}
	if key2 == nil {
		return fmt.Errorf("p2 has nil key")
	}

	for i, k := range key1 {
		if key2[i] != k {
			return fmt.Errorf("th byte does not same")
		}
	}
	return nil
}
