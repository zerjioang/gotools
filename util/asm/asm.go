// +build !noasm
// +build !appengine

package hex

//go:noescape
func _add(a, b int) (c int)

func Add(a, b int) int {
	return _add(a, b)
}
