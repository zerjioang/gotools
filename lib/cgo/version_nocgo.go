// +build !cgo

package cgo

var version = "unknown"

func init() {
}

func CgoVersion() string {
	//return version
	return version
}

func FlushCache() {
}
