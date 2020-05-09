// +build cgo

package cgo

// const char* compilerVersion() {
// #if defined(__clang__)
// 	return __VERSION__;
// #elif defined(__GNUC__) || defined(__GNUG__)
// 	return "gcc " __VERSION__;
// #else
// 	return "non-gcc, non-clang (or an unrecognized version)";
// #endif
// }
import "C"

var version = "unknown"

func init() {
	FlushCache()
}

func CgoVersion() string {
	//return version
	return version
}

func FlushCache() {
	version = C.GoString(C.compilerVersion())
}
