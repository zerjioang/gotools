// Copyright 2012-2013 Huan Truong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package md2 implements the MD2 hash algorithm as defined in RFC 1319.
package md2

import (
	"hash"
)

func init() {
	//crypto.RegisterHash(crypto.MD2, New)
}

const (
	// The size of an MD2 checksum in bytes.
	Size = 16
	// The blocksize of MD2 in bytes.
	BlockSize = 16
	_Chunk    = 16
)

var PiSubst = []uint8{
	41, 46, 67, 201, 162, 216, 124, 1, 61, 54, 84, 161, 236, 240, 6,
	19, 98, 167, 5, 243, 192, 199, 115, 140, 152, 147, 43, 217, 188,
	76, 130, 202, 30, 155, 87, 60, 253, 212, 224, 22, 103, 66, 111, 24,
	138, 23, 229, 18, 190, 78, 196, 214, 218, 158, 222, 73, 160, 251,
	245, 142, 187, 47, 238, 122, 169, 104, 121, 145, 21, 178, 7, 63,
	148, 194, 16, 137, 11, 34, 95, 33, 128, 127, 93, 154, 90, 144, 50,
	39, 53, 62, 204, 231, 191, 247, 151, 3, 255, 25, 48, 179, 72, 165,
	181, 209, 215, 94, 146, 42, 172, 86, 170, 198, 79, 184, 56, 210,
	150, 164, 125, 182, 118, 252, 107, 226, 156, 116, 4, 241, 69, 157,
	112, 89, 100, 113, 135, 32, 134, 91, 207, 101, 230, 45, 168, 2, 27,
	96, 37, 173, 174, 176, 185, 246, 28, 70, 97, 105, 52, 64, 126, 15,
	85, 71, 163, 35, 221, 81, 175, 58, 195, 92, 249, 206, 186, 197,
	234, 38, 44, 83, 13, 110, 133, 40, 132, 9, 211, 223, 205, 244, 65,
	129, 77, 82, 106, 220, 55, 200, 108, 193, 171, 250, 36, 225, 123,
	8, 12, 189, 177, 74, 120, 136, 149, 139, 227, 99, 232, 109, 233,
	203, 213, 254, 59, 0, 29, 57, 242, 239, 183, 14, 102, 88, 208, 228,
	166, 119, 114, 248, 235, 117, 75, 10, 49, 68, 80, 180, 143, 237,
	31, 26, 219, 153, 141, 51, 159, 17, 131, 20,
}

// digest represents the partial evaluation of a checksum.
type Md2Digest struct {
	digest [Size]byte   // the digest, Size
	state  [48]byte     // state, 48 ints
	x      [_Chunk]byte // temp storage buffer, 16 bytes, _Chunk
	nx     uint8        // how many bytes are there in the buffer
}

func (dg *Md2Digest) Reset() {
	/*for i := range dg.digest {
		fmt.Printf("dg.digest[%d] = 0\n", i)
		dg.digest[i] = 0
	}*/
	dg.digest[15] = 0
	dg.digest[14] = 0
	dg.digest[13] = 0
	dg.digest[12] = 0
	dg.digest[11] = 0
	dg.digest[10] = 0
	dg.digest[9] = 0
	dg.digest[8] = 0
	dg.digest[7] = 0
	dg.digest[6] = 0
	dg.digest[5] = 0
	dg.digest[4] = 0
	dg.digest[3] = 0
	dg.digest[2] = 0
	dg.digest[1] = 0
	dg.digest[0] = 0
	/*for i := len(dg.state)-1; i >= 0; i-- {
		fmt.Printf("dg.state[%d] = 0\n", i)
		dg.state[i] = 0
	}*/
	dg.state[47] = 0
	dg.state[46] = 0
	dg.state[45] = 0
	dg.state[44] = 0
	dg.state[43] = 0
	dg.state[42] = 0
	dg.state[41] = 0
	dg.state[40] = 0
	dg.state[39] = 0
	dg.state[38] = 0
	dg.state[37] = 0
	dg.state[36] = 0
	dg.state[35] = 0
	dg.state[34] = 0
	dg.state[33] = 0
	dg.state[32] = 0
	dg.state[31] = 0
	dg.state[30] = 0
	dg.state[29] = 0
	dg.state[28] = 0
	dg.state[27] = 0
	dg.state[26] = 0
	dg.state[25] = 0
	dg.state[24] = 0
	dg.state[23] = 0
	dg.state[22] = 0
	dg.state[21] = 0
	dg.state[20] = 0
	dg.state[19] = 0
	dg.state[18] = 0
	dg.state[17] = 0
	dg.state[16] = 0
	dg.state[15] = 0
	dg.state[14] = 0
	dg.state[13] = 0
	dg.state[12] = 0
	dg.state[11] = 0
	dg.state[10] = 0
	dg.state[9] = 0
	dg.state[8] = 0
	dg.state[7] = 0
	dg.state[6] = 0
	dg.state[5] = 0
	dg.state[4] = 0
	dg.state[3] = 0
	dg.state[2] = 0
	dg.state[1] = 0
	dg.state[0] = 0
	/*for i := range dg.x {
		fmt.Printf("dg.x[%d] = 0\n", i)
		dg.x[i] = 0
	}*/
	dg.x[15] = 0
	dg.x[14] = 0
	dg.x[13] = 0
	dg.x[12] = 0
	dg.x[11] = 0
	dg.x[10] = 0
	dg.x[9] = 0
	dg.x[8] = 0
	dg.x[7] = 0
	dg.x[6] = 0
	dg.x[5] = 0
	dg.x[4] = 0
	dg.x[3] = 0
	dg.x[2] = 0
	dg.x[1] = 0
	dg.x[0] = 0
	dg.nx = 0
}

// New returns a new hash.Hash computing the MD2 checksum.
func New() hash.Hash {
	d := new(Md2Digest)
	d.Reset()
	return d
}

// New returns a new hash.Hash computing the MD2 checksum.
func NewMd2() *Md2Digest {
	d := new(Md2Digest)
	d.Reset()
	return d
}

func (dg *Md2Digest) Size() int { return Size }

func (dg *Md2Digest) BlockSize() int { return BlockSize }

// Write is the interface for IO Writer
func (dg *Md2Digest) Write(p []byte) (nn int, err error) {
	nn = len(p)
	//d.len += uint64(nn)
	// If we have something left in the buffer
	if dg.nx > 0 {
		n := uint8(len(p))
		var i uint8
		// try to copy the rest n bytes free of the buffer into the buffer
		// then hash the buffer
		if (n + dg.nx) > _Chunk {
			n = _Chunk - dg.nx
		}
		for i = 0; i < n; i++ {
			// copy n bytes to the buffer
			dg.x[dg.nx+i] = p[i]
		}
		dg.nx += n
		// if we have exactly 1 block in the buffer then hash that block
		if dg.nx == _Chunk {
			block(dg, dg.x[0:_Chunk])
			dg.nx = 0
		}
		p = p[n:]
	}
	imax := len(p) / _Chunk
	// For the rest, try hashing by the blocksize
	for i := 0; i < imax; i++ {
		block(dg, p[:_Chunk])
		p = p[_Chunk:]
	}
	// Then stuff the rest that doesn't add up to a block to the buffer
	if len(p) > 0 {
		dg.nx = uint8(copy(dg.x[:], p))
	}
	return
}

func (dg *Md2Digest) Sum(in []byte) []byte {
	// Make a copy of d0 so that caller can keep writing and summing.
	d := *dg
	// Padding.
	var tmp [_Chunk]byte
	//tmp := make([]byte, _Chunk, _Chunk)
	lsize := d.nx
	/*for i := range tmp {
		fmt.Printf("tmp[%d] = _Chunk - lsize\n", i)
		tmp[i] = _Chunk - lsize
	}*/
	m := _Chunk - lsize
	tmp[15] = m
	tmp[14] = m
	tmp[13] = m
	tmp[12] = m
	tmp[11] = m
	tmp[10] = m
	tmp[9] = m
	tmp[8] = m
	tmp[7] = m
	tmp[6] = m
	tmp[5] = m
	tmp[4] = m
	tmp[3] = m
	tmp[2] = m
	tmp[1] = m
	tmp[0] = m
	//fmt.Printf("Calling block(%x) size %d for last block...\n", tmp, _Chunk-len)
	_, _ = d.Write(tmp[0:m])
	// At this state we should have nothing left in buffer
	if d.nx != 0 {
		panic("d.nx != 0")
	}
	_, _ = d.Write(d.digest[0:16])
	// At this state we should have nothing left in buffer
	if d.nx != 0 {
		panic("d.nx != 0")
	}
	return append(in, d.state[0:16]...)
}

func (dg *Md2Digest) Resolve() []byte {
	// Make a copy of d0 so that caller can keep writing and summing.
	d := *dg
	// Padding.
	var tmp [_Chunk]byte
	//tmp := make([]byte, _Chunk, _Chunk)
	lsize := d.nx
	/*for i := range tmp {
		tmp[i] = _Chunk - lsize
	}*/
	m := _Chunk - lsize
	tmp[15] = m
	tmp[14] = m
	tmp[13] = m
	tmp[12] = m
	tmp[11] = m
	tmp[10] = m
	tmp[9] = m
	tmp[8] = m
	tmp[7] = m
	tmp[6] = m
	tmp[5] = m
	tmp[4] = m
	tmp[3] = m
	tmp[2] = m
	tmp[1] = m
	tmp[0] = m
	//fmt.Printf("Calling block(%x) size %d for last block...\n", tmp, _Chunk-len)
	_, _ = d.Write(tmp[0:m])
	// At this state we should have nothing left in buffer
	if d.nx != 0 {
		panic("d.nx != 0")
	}
	_, _ = d.Write(d.digest[0:16])
	// At this state we should have nothing left in buffer
	if d.nx != 0 {
		panic("d.nx != 0")
	}
	return d.state[0:16]
}

func block(dig *Md2Digest, p []byte) {
	var t, i, j uint8
	t = 0
	for i = 0; i < 16; i++ {
		dig.state[i+16] = p[i]
		dig.state[i+32] = p[i] ^ dig.state[i]
	}
	for i = 0; i < 18; i++ {
		for j = 0; j < 48; j++ {
			dig.state[j] = dig.state[j] ^ PiSubst[t]
			t = dig.state[j]
		}
		t = t + i
	}
	t = dig.digest[15]
	for i = 0; i < 16; i++ {
		dig.digest[i] ^= PiSubst[p[i]^t]
		t = dig.digest[i]
	}
}
