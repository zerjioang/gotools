package atomicvar

import "sync/atomic"

type AtomicBoolUint32 uint32

func (b *AtomicBoolUint32) IsFalse() bool {
	return atomic.LoadUint32((*uint32)(b)) == 0
}

func (b *AtomicBoolUint32) IsTrue() bool {
	return atomic.LoadUint32((*uint32)(b)) == 1
}

func (b *AtomicBoolUint32) SetTrue() {
	atomic.StoreUint32((*uint32)(b), 1)
}

func (b *AtomicBoolUint32) SetFalse() {
	atomic.StoreUint32((*uint32)(b), 0)
}
