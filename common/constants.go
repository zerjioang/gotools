package common

const (
	//new line character
	NewLine = "\n"
	// set system pointer size
	PointerSize = 32 + int(^uintptr(0)>>63<<5)
)
