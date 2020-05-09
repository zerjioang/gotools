package base58

const base58Size = 58
const alphabetIdx0 = '1'

// Alphabet is a a b58 alphabet.
type Alphabet struct {
	decode [128]int8
	encode [base58Size]byte
}

// NewAlphabet creates a new alphabet from the passed string.
//
// It panics if the passed string is not 58 bytes long or isn't valid ASCII.
func NewAlphabet(s string) Alphabet {
	ret := Alphabet{}
	if len(s) == base58Size {
		copy(ret.encode[:], s)
		for i := 0; i < len(ret.encode); i++ {
			b := ret.encode[i]
			ret.encode[i] = s[i]
			ret.decode[b] = int8(i)
		}
	}
	return ret
}

// BTCAlphabet is the bitcoin base58 alphabet.
// alphabet is the modified base58 alphabet used by Bitcoin.
var BTCAlphabet = NewAlphabet("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// FlickrAlphabet is the flickr base58 alphabet.
var FlickrAlphabet = NewAlphabet("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")
