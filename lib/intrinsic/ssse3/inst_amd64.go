package ssse3

//go:noescape
// PABSBm128byte Packed Absolute Value Integers
func PABSBm128byte(X1 []byte, X2 []byte)

// PABSBm128int8 Packed Absolute Value Integers
//go:noescape
func PABSBm128int8(X1 []int8, X2 []int8)

// PABSDm128byte Packed Absolute Value Integers
//go:noescape
func PABSDm128byte(X1 []byte, X2 []byte)

//go:noescape
// PABSDm128int32 Packed Absolute Value Integers
//go:noescape
func PABSDm128int32(X1 []int32, X2 []int32)

// PABSWm128byte Packed Absolute Value Integers
//go:noescape
func PABSWm128byte(X1 []byte, X2 []byte)

// PABSWm128int16 Packed Absolute Value Integers
//go:noescape
func PABSWm128int16(X1 []int16, X2 []int16)

// PHADDDm128byte Packed Horizontal Add
//go:noescape
func PHADDDm128byte(X1 []byte, X2 []byte)

// PHADDSWm128byte Packed Horizontal Add and Saturate
//go:noescape
func PHADDSWm128byte(X1 []byte, X2 []byte)

// PHADDWm128byte Packed Horizontal Add
//go:noescape
func PHADDWm128byte(X1 []byte, X2 []byte)

// PHSUBDm128byte Packed Horizontal Subtract
//go:noescape
func PHSUBDm128byte(X1 []byte, X2 []byte)

// PHSUBSWm128byte Packed Horizontal Subtract and Saturate
//go:noescape
func PHSUBSWm128byte(X1 []byte, X2 []byte)

// PHSUBWm128byte Packed Horizontal Subtract
//go:noescape
func PHSUBWm128byte(X1 []byte, X2 []byte)

// PMADDUBSWm128byte Multiply and Add Packed Signed and Unsigned Bytes
//go:noescape
func PMADDUBSWm128byte(X1 []byte, X2 []byte)

// PMULHRSWm128byte Packed Multiply High with Round and Scale
//go:noescape
func PMULHRSWm128byte(X1 []byte, X2 []byte)

// PSHUFBm128byte Packed Shuffle Bytes
//go:noescape
func PSHUFBm128byte(X1 []byte, X2 []byte)

// PSIGNBm128byte Packed SIGN
//go:noescape
func PSIGNBm128byte(X1 []byte, X2 []byte)

// PSIGNDm128byte Packed SIGN
//go:noescape
func PSIGNDm128byte(X1 []byte, X2 []byte)

// PSIGNWm128byte Packed SIGN
//go:noescape
func PSIGNWm128byte(X1 []byte, X2 []byte)
