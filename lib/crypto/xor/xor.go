package xor

// EncryptDecrypt runs a XOR encryption on the input string, encrypting it if it hasn't already been,
// and decrypting it if it has, using the key provided.
func EncryptDecrypt(input, key string) (output string) {
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}
	return output
}

func EncryptDecryptBytes(input, key []byte) []byte {
	l := len(input)
	lk := len(key)
	output := make([]byte, l)
	for i := l - 1; i >= 0; i-- {
		output[i] = input[i] ^ key[i%lk]
	}
	return output
}
