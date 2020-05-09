package arena

// One of the weaknesses of Go’s runtime today is the relatively naive GC implementation.
// This is evident from go performing consistently worse than most other languages in
// the binary trees benchmark. However, the language can make designing programs that
// reduce GC cost fairly straightforward.
//
// On my laptop, the Go program runs in 32.32s and consumes 348.3 MB of resident memory.
// In comparison, the Haskell version (with “+RTS -N4 -K128M -H -RTS”) runs in 16.06s
// and consumes 772.6MB!
//
// The code used in the benchmark creates and discards several binary trees by allocating
// each node individually. A small change in the code to allocate several nodes whenever
// allocation is required, instead of just one, can significantly reduce the amount of
// work performed by GC. This technique is called arena allocation.
//
// The NewChannelArena type’s Pop method returns a pointer to an element in the slice and allocates
// in chunks of 10,000 when the slice becomes empty.

// ChannelArena is a free list that provides quick access to pre-allocated byte
// slices, greatly reducing memory churn and effectively disabling GC for these
// allocations. After the ChannelArena is created, a slice of bytes can be requested by
// calling Pop(). The caller is responsible for calling Push(), which puts the
// blocks back in the queue for later usage. The bytes given by Pop() are *not*
// zeroed, so the caller should only read positions that it knows to have been
// overwitten. That can be done by shortening the slice at the right place,
// based on the count of bytes returned by Write() and similar functions.
type ChannelArena chan []byte

func NewChannelArena(numBlocks int, blockSize int) ChannelArena {
	// blocks: is a list of ChannelArena
	blocks := make(ChannelArena, numBlocks)
	for i := 0; i < numBlocks; i++ {
		blocks <- make([]byte, blockSize)
	}
	return blocks
}

func (a ChannelArena) Pop() (x []byte) {
	return <-a
}

func (a ChannelArena) Push(x []byte) {
	x = x[:cap(x)]
	a <- x
}

func (a ChannelArena) PushByte(x byte) {
	a <- []byte{x}
}
