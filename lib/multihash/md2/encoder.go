package md2

func Encoder(data []byte) ([]byte, error) {
	h := new(Md2Digest)
	h.Reset()
	_, _ = h.Write(data)
	return h.Resolve(), nil
}
