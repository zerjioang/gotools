package _defer

func doNoDefer(t *int) {
	func() {
		*t++
	}()
}

func doDefer(t *int) {
	defer func() {
		*t++
	}()
}
