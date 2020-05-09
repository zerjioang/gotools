// +build !noasm
// +build !appengine

package convert

import "unsafe"

/*
//go:noescape
func _ip_to_int(buf unsafe.Pointer) (result uint32)
func IpToInt(ip []byte) uint32 {
	r := _ip_to_int(
		unsafe.Pointer(&ip[0]),
	)
	return r
}
*/

//go:noescape
func _ip_to_int2(buf, size unsafe.Pointer) (result uint32)

func IpToIntAssemblyAmd64(ip []byte) uint32 {
	r := _ip_to_int2(
		unsafe.Pointer(&ip[0]),
		unsafe.Pointer(uintptr(len(ip))),
	)
	return r
}
