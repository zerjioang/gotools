// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package longpool

// LongPool holds interface{}.
type LongPool struct {
	New  func() interface{}
	pool chan interface{}
}

// NewPool creates a new pool of interface{}.
func NewLongPool(max int) *LongPool {
	return &LongPool{
		pool: make(chan interface{}, max),
	}
}

// Borrow an object from the pool.
func (p *LongPool) Get() interface{} {
	var c interface{}
	select {
	case c = <-p.pool:
	default:
		c = p.New()
	}
	return c
}

// Return returns an object to the pool.
func (p *LongPool) Put(c interface{}) {
	select {
	case p.pool <- c:
	default:
		// let it go, let it go...
	}
}
