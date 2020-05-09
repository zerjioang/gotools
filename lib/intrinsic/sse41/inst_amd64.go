package sse41

// PACKUSDWm128byte Pack with Unsigned Saturation
//go:noescape
func PACKUSDWm128byte(X1 []byte, X2 []byte)

// PCMPEQQm128byte Compare Packed Qword Data for Equal
//go:noescape
func PCMPEQQm128byte(X1 []byte, X2 []byte)

// PHMINPOSUWm128byte Packed Horizontal Word Minimum
//go:noescape
func PHMINPOSUWm128byte(X1 []byte, X2 []byte)

// PMAXSBm128byte Maximum of Packed Signed Integers
//go:noescape
func PMAXSBm128byte(X1 []byte, X2 []byte)

// PMAXSBm128int8 Maximum of Packed Signed Integers
//go:noescape
func PMAXSBm128int8(X1 []int8, X2 []int8)

// PMAXSDm128byte Maximum of Packed Signed Integers
//go:noescape
func PMAXSDm128byte(X1 []byte, X2 []byte)

// PMAXSDm128int32 Maximum of Packed Signed Integers
//go:noescape
func PMAXSDm128int32(X1 []int32, X2 []int32)

// PMAXUDm128byte Maximum of Packed Unsigned Integers
//go:noescape
func PMAXUDm128byte(X1 []byte, X2 []byte)

// PMAXUDm128uint32 Maximum of Packed Unsigned Integers
//go:noescape
func PMAXUDm128uint32(X1 []uint32, X2 []uint32)

// PMAXUWm128byte Maximum of Packed Unsigned Integers
//go:noescape
func PMAXUWm128byte(X1 []byte, X2 []byte)

// PMAXUWm128uint16 Maximum of Packed Unsigned Integers
//go:noescape
func PMAXUWm128uint16(X1 []uint16, X2 []uint16)

// PMINSBm128byte Minimum of Packed Signed Integers
//go:noescape
func PMINSBm128byte(X1 []byte, X2 []byte)

// PMINSBm128int8 Minimum of Packed Signed Integers
//go:noescape
func PMINSBm128int8(X1 []int8, X2 []int8)

// PMINSDm128byte Minimum of Packed Signed Integers
//go:noescape
func PMINSDm128byte(X1 []byte, X2 []byte)

// PMINSDm128int32 Minimum of Packed Signed Integers
//go:noescape
func PMINSDm128int32(X1 []int32, X2 []int32)

// PMINUDm128byte Minimum of Packed Unsigned Integers
//go:noescape
func PMINUDm128byte(X1 []byte, X2 []byte)

// PMINUDm128uint32 Minimum of Packed Unsigned Integers
//go:noescape
func PMINUDm128uint32(X1 []uint32, X2 []uint32)

// PMINUWm128byte Minimum of Packed Unsigned Integers
//go:noescape
func PMINUWm128byte(X1 []byte, X2 []byte)

// PMINUWm128uint16 Minimum of Packed Unsigned Integers
//go:noescape
func PMINUWm128uint16(X1 []uint16, X2 []uint16)

// PMOVSXBDm32byte Packed Move with Sign Extend
//go:noescape
func PMOVSXBDm32byte(X1 []byte, X2 []byte)

// PMOVSXBQm16byte Packed Move with Sign Extend
//go:noescape
func PMOVSXBQm16byte(X1 []byte, X2 []byte)

// PMOVSXBWm64byte Packed Move with Sign Extend
//go:noescape
func PMOVSXBWm64byte(X1 []byte, X2 []byte)

// PMOVSXDQm64byte Packed Move with Sign Extend
//go:noescape
func PMOVSXDQm64byte(X1 []byte, X2 []byte)

// PMOVSXWDm64byte Packed Move with Sign Extend
//go:noescape
func PMOVSXWDm64byte(X1 []byte, X2 []byte)

// PMOVSXWQm32byte Packed Move with Sign Extend
//go:noescape
func PMOVSXWQm32byte(X1 []byte, X2 []byte)

// PMOVZXBDm32byte Packed Move with Zero Extend
//go:noescape
func PMOVZXBDm32byte(X1 []byte, X2 []byte)

// PMOVZXBQm16byte Packed Move with Zero Extend
//go:noescape
func PMOVZXBQm16byte(X1 []byte, X2 []byte)

// PMOVZXBWm64byte Packed Move with Zero Extend
//go:noescape
func PMOVZXBWm64byte(X1 []byte, X2 []byte)

// PMOVZXDQm64byte Packed Move with Zero Extend
//go:noescape
func PMOVZXDQm64byte(X1 []byte, X2 []byte)

// PMOVZXWDm64byte Packed Move with Zero Extend
//go:noescape
func PMOVZXWDm64byte(X1 []byte, X2 []byte)

// PMOVZXWQm32byte Packed Move with Zero Extend
//go:noescape
func PMOVZXWQm32byte(X1 []byte, X2 []byte)

// PMULDQm128byte Multiply Packed Doubleword Integers
//go:noescape
func PMULDQm128byte(X1 []byte, X2 []byte)

// PMULDQm128int64 Multiply Packed Doubleword Integers
//go:noescape
func PMULDQm128int64(X1 []int64, X2 []int64)

// PMULLDm128byte Multiply Packed Integers and Store Low Result
//go:noescape
func PMULLDm128byte(X1 []byte, X2 []byte)

// PMULLDm128int32 Multiply Packed Integers and Store Low Result
//go:noescape
func PMULLDm128int32(X1 []int32, X2 []int32)

// PTESTm128byte PTEST- Logical Compare
//go:noescape
func PTESTm128byte(X1 []byte, X2 []byte)
