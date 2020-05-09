package sse2

// ADDPDm128byte Add Packed Double-Precision Floating-Point Values
//go:noescape
func ADDPDm128byte(X1 []byte, X2 []byte)

// ADDPDm128float64 Add Packed Double-Precision Floating-Point Values
//go:noescape
func ADDPDm128float64(X1 []float64, X2 []float64)

// ADDSDm64byte Add Scalar Double-Precision Floating-Point Values
//go:noescape
func ADDSDm64byte(X1 []byte, X2 []byte)

// ADDSDm64float64 Add Scalar Double-Precision Floating-Point Values
//go:noescape
func ADDSDm64float64(X1 []float64, X2 []float64)

// ANDNPDm128byte Bitwise Logical AND NOT of Packed Double Precision Floating-Point Values
//go:noescape
func ANDNPDm128byte(X1 []byte, X2 []byte)

// ANDNPDm128float64 Bitwise Logical AND NOT of Packed Double Precision Floating-Point Values
//go:noescape
func ANDNPDm128float64(X1 []float64, X2 []float64)

// ANDPDm128byte Bitwise Logical AND of Packed Double Precision Floating-Point Values
//go:noescape
func ANDPDm128byte(X1 []byte, X2 []byte)

// ANDPDm128float64 Bitwise Logical AND of Packed Double Precision Floating-Point Values
//go:noescape
func ANDPDm128float64(X1 []float64, X2 []float64)

// COMISDm64byte Compare Scalar Ordered Double-Precision Floating-Point Values and Set EFLAGS
//go:noescape
func COMISDm64byte(X1 []byte, X2 []byte)

// COMISDm64float64 Compare Scalar Ordered Double-Precision Floating-Point Values and Set EFLAGS
//go:noescape
func COMISDm64float64(X1 []float64, X2 []float64)

// DIVPDm128byte Divide Packed Double-Precision Floating-Point Values
//go:noescape
func DIVPDm128byte(X1 []byte, X2 []byte)

// DIVPDm128float64 Divide Packed Double-Precision Floating-Point Values
//go:noescape
func DIVPDm128float64(X1 []float64, X2 []float64)

// DIVSDm64byte Divide Scalar Double-Precision Floating-Point Value
//go:noescape
func DIVSDm64byte(X1 []byte, X2 []byte)

// DIVSDm64float64 Divide Scalar Double-Precision Floating-Point Value
//go:noescape
func DIVSDm64float64(X1 []float64, X2 []float64)

// MASKMOVDQUbyte Store Selected Bytes of Double Quadword
//go:noescape
func MASKMOVDQUbyte(X1 []byte, X2 []byte)

// MAXPDm128byte Maximum of Packed Double-Precision Floating-Point Values
//go:noescape
func MAXPDm128byte(X1 []byte, X2 []byte)

// MAXPDm128float64 Maximum of Packed Double-Precision Floating-Point Values
//go:noescape
func MAXPDm128float64(X1 []float64, X2 []float64)

// MAXSDm64byte Return Maximum Scalar Double-Precision Floating-Point Value
//go:noescape
func MAXSDm64byte(X1 []byte, X2 []byte)

// MAXSDm64float64 Return Maximum Scalar Double-Precision Floating-Point Value
//go:noescape
func MAXSDm64float64(X1 []float64, X2 []float64)

// MINPDm128byte Minimum of Packed Double-Precision Floating-Point Values
//go:noescape
func MINPDm128byte(X1 []byte, X2 []byte)

// MINPDm128float64 Minimum of Packed Double-Precision Floating-Point Values
//go:noescape
func MINPDm128float64(X1 []float64, X2 []float64)

// MINSDm64byte Return Minimum Scalar Double-Precision Floating-Point Value
//go:noescape
func MINSDm64byte(X1 []byte, X2 []byte)

// MINSDm64float64 Return Minimum Scalar Double-Precision Floating-Point Value
//go:noescape
func MINSDm64float64(X1 []float64, X2 []float64)

// MULPDm128byte Multiply Packed Double-Precision Floating-Point Values
//go:noescape
func MULPDm128byte(X1 []byte, X2 []byte)

// MULPDm128float64 Multiply Packed Double-Precision Floating-Point Values
//go:noescape
func MULPDm128float64(X1 []float64, X2 []float64)

// MULSDm64byte Multiply Scalar Double-Precision Floating-Point Value
//go:noescape
func MULSDm64byte(X1 []byte, X2 []byte)

// MULSDm64float64 Multiply Scalar Double-Precision Floating-Point Value
//go:noescape
func MULSDm64float64(X1 []float64, X2 []float64)

// ORPDm128byte Bitwise Logical OR of Packed Double Precision Floating-Point Values
//go:noescape
func ORPDm128byte(X1 []byte, X2 []byte)

// ORPDm128float64 Bitwise Logical OR of Packed Double Precision Floating-Point Values
//go:noescape
func ORPDm128float64(X1 []float64, X2 []float64)

// PACKSSDWm128byte Pack with Signed Saturation
//go:noescape
func PACKSSDWm128byte(X1 []byte, X2 []byte)

// PACKSSWBm128byte Pack with Signed Saturation
//go:noescape
func PACKSSWBm128byte(X1 []byte, X2 []byte)

// PACKUSWBm128byte Pack with Unsigned Saturation
//go:noescape
func PACKUSWBm128byte(X1 []byte, X2 []byte)

// PADDBm128byte Add Packed Integers
//go:noescape
func PADDBm128byte(X1 []byte, X2 []byte)

// PADDBm128int8 Add Packed Integers
//go:noescape
func PADDBm128int8(X1 []int8, X2 []int8)

// PADDDm128byte Add Packed Integers
//go:noescape
func PADDDm128byte(X1 []byte, X2 []byte)

// PADDDm128int32 Add Packed Integers
//go:noescape
func PADDDm128int32(X1 []int32, X2 []int32)

// PADDQm128byte Add Packed Integers
//go:noescape
func PADDQm128byte(X1 []byte, X2 []byte)

// PADDQm128int64 Add Packed Integers
//go:noescape
func PADDQm128int64(X1 []int64, X2 []int64)

// PADDSBm128byte Add Packed Signed Integers with Signed Saturation
//go:noescape
func PADDSBm128byte(X1 []byte, X2 []byte)

// PADDSBm128int8 Add Packed Signed Integers with Signed Saturation
//go:noescape
func PADDSBm128int8(X1 []int8, X2 []int8)

// PADDSWm128byte Add Packed Signed Integers with Signed Saturation
//go:noescape
func PADDSWm128byte(X1 []byte, X2 []byte)

// PADDSWm128int16 Add Packed Signed Integers with Signed Saturation
//go:noescape
func PADDSWm128int16(X1 []int16, X2 []int16)

// PADDUSBm128byte Add Packed Unsigned Integers with Unsigned Saturation
//go:noescape
func PADDUSBm128byte(X1 []byte, X2 []byte)

// PADDUSBm128uint8 Add Packed Unsigned Integers with Unsigned Saturation
//go:noescape
func PADDUSBm128uint8(X1 []uint8, X2 []uint8)

// PADDUSWm128byte Add Packed Unsigned Integers with Unsigned Saturation
//go:noescape
func PADDUSWm128byte(X1 []byte, X2 []byte)

// PADDUSWm128uint16 Add Packed Unsigned Integers with Unsigned Saturation
//go:noescape
func PADDUSWm128uint16(X1 []uint16, X2 []uint16)

// PADDWm128byte Add Packed Integers
//go:noescape
func PADDWm128byte(X1 []byte, X2 []byte)

// PADDWm128int16 Add Packed Integers
//go:noescape
func PADDWm128int16(X1 []int16, X2 []int16)

// PANDm128byte Logical AND
//go:noescape
func PANDm128byte(X1 []byte, X2 []byte)

// PANDNm128byte Logical AND NOT
//go:noescape
func PANDNm128byte(X1 []byte, X2 []byte)

// PAVGBm128byte Average Packed Integers
//go:noescape
func PAVGBm128byte(X1 []byte, X2 []byte)

// PAVGBm128int8 Average Packed Integers
//go:noescape
func PAVGBm128int8(X1 []int8, X2 []int8)

// PAVGWm128byte Average Packed Integers
//go:noescape
func PAVGWm128byte(X1 []byte, X2 []byte)

// PAVGWm128int16 Average Packed Integers
//go:noescape
func PAVGWm128int16(X1 []int16, X2 []int16)

// PCMPEQBm128byte Compare Packed Data for Equal
//go:noescape
func PCMPEQBm128byte(X1 []byte, X2 []byte)

// PCMPEQDm128byte Compare Packed Data for Equal
//go:noescape
func PCMPEQDm128byte(X1 []byte, X2 []byte)

// PCMPEQWm128byte Compare Packed Data for Equal
//go:noescape
func PCMPEQWm128byte(X1 []byte, X2 []byte)

// PCMPGTBm128byte Compare Packed Signed Integers for Greater Than
//go:noescape
func PCMPGTBm128byte(X1 []byte, X2 []byte)

// PCMPGTBm128int8 Compare Packed Signed Integers for Greater Than
//go:noescape
func PCMPGTBm128int8(X1 []int8, X2 []int8)

// PCMPGTDm128byte Compare Packed Signed Integers for Greater Than
//go:noescape
func PCMPGTDm128byte(X1 []byte, X2 []byte)

// PCMPGTDm128int32 Compare Packed Signed Integers for Greater Than
//go:noescape
func PCMPGTDm128int32(X1 []int32, X2 []int32)

// PCMPGTWm128byte Compare Packed Signed Integers for Greater Than
//go:noescape
func PCMPGTWm128byte(X1 []byte, X2 []byte)

// PCMPGTWm128int16 Compare Packed Signed Integers for Greater Than
//go:noescape
func PCMPGTWm128int16(X1 []int16, X2 []int16)

// PMADDWDm128byte Multiply and Add Packed Integers
//go:noescape
func PMADDWDm128byte(X1 []byte, X2 []byte)

// PMADDWDm128int32 Multiply and Add Packed Integers
//go:noescape
func PMADDWDm128int32(X1 []int32, X2 []int32)

// PMAXSWm128byte Maximum of Packed Signed Integers
//go:noescape
func PMAXSWm128byte(X1 []byte, X2 []byte)

// PMAXSWm128int16 Maximum of Packed Signed Integers
//go:noescape
func PMAXSWm128int16(X1 []int16, X2 []int16)

// PMAXUBm128byte Maximum of Packed Unsigned Integers
//go:noescape
func PMAXUBm128byte(X1 []byte, X2 []byte)

// PMAXUBm128uint8 Maximum of Packed Unsigned Integers
//go:noescape
func PMAXUBm128uint8(X1 []uint8, X2 []uint8)

// PMINSWm128byte Minimum of Packed Signed Integers
//go:noescape
func PMINSWm128byte(X1 []byte, X2 []byte)

// PMINSWm128int16 Minimum of Packed Signed Integers
//go:noescape
func PMINSWm128int16(X1 []int16, X2 []int16)

// PMINUBm128byte Minimum of Packed Unsigned Integers
func PMINUBm128byte(X1 []byte, X2 []byte)

//go:noescape
// PMINUBm128uint8 Minimum of Packed Unsigned Integers
func PMINUBm128uint8(X1 []uint8, X2 []uint8)

//go:noescape
// PMULHUWm128byte Multiply Packed Unsigned Integers and Store High Result
func PMULHUWm128byte(X1 []byte, X2 []byte)

//go:noescape
// PMULHUWm128uint16 Multiply Packed Unsigned Integers and Store High Result
func PMULHUWm128uint16(X1 []uint16, X2 []uint16)

//go:noescape
// PMULHWm128byte Multiply Packed Signed Integers and Store High Result
func PMULHWm128byte(X1 []byte, X2 []byte)

//go:noescape
// PMULHWm128int16 Multiply Packed Signed Integers and Store High Result
//go:noescape
func PMULHWm128int16(X1 []int16, X2 []int16)

// PMULLWm128byte Multiply Packed Signed Integers and Store Low Result
//go:noescape
func PMULLWm128byte(X1 []byte, X2 []byte)

// PMULLWm128int16 Multiply Packed Signed Integers and Store Low Result
//go:noescape
func PMULLWm128int16(X1 []int16, X2 []int16)

// PMULUDQm128byte Multiply Packed Unsigned Doubleword Integers
//go:noescape
func PMULUDQm128byte(X1 []byte, X2 []byte)

// PMULUDQm128int64 Multiply Packed Unsigned Doubleword Integers
//go:noescape
func PMULUDQm128int64(X1 []int64, X2 []int64)

// PORm128byte Bitwise Logical OR
//go:noescape
func PORm128byte(X1 []byte, X2 []byte)

// PSADBWm128byte Compute Sum of Absolute Differences
//go:noescape
func PSADBWm128byte(X1 []byte, X2 []byte)

// PSLLDm128byte Shift Packed Data Left Logical
//go:noescape
func PSLLDm128byte(X1 []byte, X2 []byte)

// PSLLQm128byte Shift Packed Data Left Logical
//go:noescape
func PSLLQm128byte(X1 []byte, X2 []byte)

// PSLLWm128byte Shift Packed Data Left Logical
//go:noescape
func PSLLWm128byte(X1 []byte, X2 []byte)

// PSRADm128byte Shift Packed Data Right Arithmetic
//go:noescape
func PSRADm128byte(X1 []byte, X2 []byte)

// PSRAWm128byte Shift Packed Data Right Arithmetic
//go:noescape
func PSRAWm128byte(X1 []byte, X2 []byte)

// PSRLDm128byte Shift Packed Data Right Logical
//go:noescape
func PSRLDm128byte(X1 []byte, X2 []byte)

// PSRLQm128byte Shift Packed Data Right Logical
//go:noescape
func PSRLQm128byte(X1 []byte, X2 []byte)

// PSRLWm128byte Shift Packed Data Right Logical
//go:noescape
func PSRLWm128byte(X1 []byte, X2 []byte)

// PSUBBm128byte Subtract Packed Integers
//go:noescape
func PSUBBm128byte(X1 []byte, X2 []byte)

// PSUBBm128int8 Subtract Packed Integers
//go:noescape
func PSUBBm128int8(X1 []int8, X2 []int8)

// PSUBDm128byte Subtract Packed Integers
//go:noescape
func PSUBDm128byte(X1 []byte, X2 []byte)

// PSUBDm128int32 Subtract Packed Integers
//go:noescape
func PSUBDm128int32(X1 []int32, X2 []int32)

// PSUBQm128byte Subtract Packed Quadword Integers
//go:noescape
func PSUBQm128byte(X1 []byte, X2 []byte)

// PSUBQm128int64 Subtract Packed Quadword Integers
//go:noescape
func PSUBQm128int64(X1 []int64, X2 []int64)

// PSUBSBm128byte Subtract Packed Signed Integers with Signed Saturation
//go:noescape
func PSUBSBm128byte(X1 []byte, X2 []byte)

// PSUBSBm128int8 Subtract Packed Signed Integers with Signed Saturation
//go:noescape
func PSUBSBm128int8(X1 []int8, X2 []int8)

// PSUBSWm128byte Subtract Packed Signed Integers with Signed Saturation
//go:noescape
func PSUBSWm128byte(X1 []byte, X2 []byte)

// PSUBSWm128int16 Subtract Packed Signed Integers with Signed Saturation
//go:noescape
func PSUBSWm128int16(X1 []int16, X2 []int16)

// PSUBUSBm128byte Subtract Packed Unsigned Integers with Unsigned Saturation
//go:noescape
func PSUBUSBm128byte(X1 []byte, X2 []byte)

// PSUBUSBm128uint8 Subtract Packed Unsigned Integers with Unsigned Saturation
//go:noescape
func PSUBUSBm128uint8(X1 []uint8, X2 []uint8)

// PSUBUSWm128byte Subtract Packed Unsigned Integers with Unsigned Saturation
//go:noescape
func PSUBUSWm128byte(X1 []byte, X2 []byte)

// PSUBUSWm128uint16 Subtract Packed Unsigned Integers with Unsigned Saturation
//go:noescape
func PSUBUSWm128uint16(X1 []uint16, X2 []uint16)

// PSUBWm128byte Subtract Packed Integers
//go:noescape
func PSUBWm128byte(X1 []byte, X2 []byte)

// PSUBWm128int16 Subtract Packed Integers
//go:noescape
func PSUBWm128int16(X1 []int16, X2 []int16)

// PUNPCKHBWm128byte Unpack High Data
//go:noescape
func PUNPCKHBWm128byte(X1 []byte, X2 []byte)

// PUNPCKHDQm128byte Unpack High Data
//go:noescape
func PUNPCKHDQm128byte(X1 []byte, X2 []byte)

// PUNPCKHQDQm128byte Unpack High Data
//go:noescape
func PUNPCKHQDQm128byte(X1 []byte, X2 []byte)

// PUNPCKHWDm128byte Unpack High Data
//go:noescape
func PUNPCKHWDm128byte(X1 []byte, X2 []byte)

// PUNPCKLBWm128byte Unpack Low Data
//go:noescape
func PUNPCKLBWm128byte(X1 []byte, X2 []byte)

// PUNPCKLDQm128byte Unpack Low Data
//go:noescape
func PUNPCKLDQm128byte(X1 []byte, X2 []byte)

// PUNPCKLQDQm128byte Unpack Low Data
//go:noescape
func PUNPCKLQDQm128byte(X1 []byte, X2 []byte)

// PUNPCKLWDm128byte Unpack Low Data
//go:noescape
func PUNPCKLWDm128byte(X1 []byte, X2 []byte)

// PXORm128byte Logical Exclusive OR
//go:noescape
func PXORm128byte(X1 []byte, X2 []byte)

// SQRTPDm128byte Square Root of Double-Precision Floating-Point Values
//go:noescape
func SQRTPDm128byte(X1 []byte, X2 []byte)

// SQRTPDm128float64 Square Root of Double-Precision Floating-Point Values
//go:noescape
func SQRTPDm128float64(X1 []float64, X2 []float64)

// SQRTSDm64byte Compute Square Root of Scalar Double-Precision Floating-Point Value
//go:noescape
func SQRTSDm64byte(X1 []byte, X2 []byte)

// SQRTSDm64float64 Compute Square Root of Scalar Double-Precision Floating-Point Value
//go:noescape
func SQRTSDm64float64(X1 []float64, X2 []float64)

// SUBPDm128byte Subtract Packed Double-Precision Floating-Point Values
//go:noescape
func SUBPDm128byte(X1 []byte, X2 []byte)

// SUBPDm128float64 Subtract Packed Double-Precision Floating-Point Values
//go:noescape
func SUBPDm128float64(X1 []float64, X2 []float64)

// SUBSDm64byte Subtract Scalar Double-Precision Floating-Point Value
//go:noescape
func SUBSDm64byte(X1 []byte, X2 []byte)

// SUBSDm64float64 Subtract Scalar Double-Precision Floating-Point Value
//go:noescape
func SUBSDm64float64(X1 []float64, X2 []float64)

// UCOMISDm64byte Unordered Compare Scalar Double-Precision Floating-Point Values and Set EFLAGS
//go:noescape
func UCOMISDm64byte(X1 []byte, X2 []byte)

// UCOMISDm64float64 Unordered Compare Scalar Double-Precision Floating-Point Values and Set EFLAGS
//go:noescape
func UCOMISDm64float64(X1 []float64, X2 []float64)

// UNPCKHPDm128byte Unpack and Interleave High Packed Double-Precision Floating-Point Values
//go:noescape
func UNPCKHPDm128byte(X1 []byte, X2 []byte)

// UNPCKHPDm128float64 Unpack and Interleave High Packed Double-Precision Floating-Point Values
//go:noescape
func UNPCKHPDm128float64(X1 []float64, X2 []float64)

// UNPCKLPDm128byte Unpack and Interleave Low Packed Double-Precision Floating-Point Values
//go:noescape
func UNPCKLPDm128byte(X1 []byte, X2 []byte)

// UNPCKLPDm128float64 Unpack and Interleave Low Packed Double-Precision Floating-Point Values
//go:noescape
func UNPCKLPDm128float64(X1 []float64, X2 []float64)

// XORPDm128byte Bitwise Logical XOR of Packed Double Precision Floating-Point Values
//go:noescape
func XORPDm128byte(X1 []byte, X2 []byte)

// XORPDm128float64 Bitwise Logical XOR of Packed Double Precision Floating-Point Values
//go:noescape
func XORPDm128float64(X1 []float64, X2 []float64)
