package chacha20

type ChachaEncoder struct {
	key []byte
}

func NewChachaEncoderParams(key []byte) *ChachaEncoder {
	return &ChachaEncoder{
		key: key,
	}
}

func (enc *ChachaEncoder) Encrypt(message []byte) ([]byte, []byte, error) {
	return Encrypt(enc.key, message)
}

func (enc *ChachaEncoder) EncryptWithKey(key []byte, message []byte) ([]byte, []byte, error) {
	return Encrypt(key, message)
}

func (enc *ChachaEncoder) Decrypt(ciphertext []byte, nonce []byte) ([]byte, error) {
	return Decrypt(enc.key, nonce, ciphertext)
}
