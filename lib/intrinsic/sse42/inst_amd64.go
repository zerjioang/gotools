package sse42

// PCMPGTQm128byte Compare Packed Signed Integers for Greater Than
//go:noescape
func PCMPGTQm128byte(X1 []byte, X2 []byte)

// PCMPGTQm128int64 Compare Packed Signed Integers for Greater Than
//go:noescape
func PCMPGTQm128int64(X1 []int64, X2 []int64)
