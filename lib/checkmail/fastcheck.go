package checkmail

import "strings"

const (
	// minimum email size is a@b.c (4)
	minEmailSize = 4
	arroba       = "@"
)

func FastEmailCheck(email string) bool {
	size := len(email)
	if size < minEmailSize {
		return false
	}
	//lets find the @
	arrobaIdx := strings.Index(email, arroba)
	leftSide := email[0:arrobaIdx]
	rightSide := email[arrobaIdx:]
	if len(leftSide) > 0 {
		// we have some email name
		// todo validate email name
		if len(rightSide) > 0 {
			// we have some domain
			// todo validate email domain
			return true
		}
	}
	return false
}
