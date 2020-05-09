package base58

import (
	"testing"

	"github.com/pkg/profile"
	"github.com/stretchr/testify/assert"
)

func TestBase58_encode_decode_table(t *testing.T) {
	t.Run("encode-decode-fast", func(t *testing.T) {
		testAddr := []string{
			"1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq",
			"1DhRmSGnhPjUaVPAj48zgPV9e2oRhAQFUb",
			"17LN2oPYRYsXS9TdYdXCCDvF2FegshLDU2",
			"14h2bDLZSuvRFhUL45VjPHJcW667mmRAAn",
		}
		for ii, vv := range testAddr {
			// num := Base58Decode([]byte(vv))
			// chk := Base58Encode(num)
			num, err := FastBase58Decoding(vv)
			if err != nil {
				t.Errorf("Test %d, expected success, got error %s\n", ii, err)
			}
			chk := FastBase58Encoding(num)
			if vv != string(chk) {
				t.Errorf("Test %d, expected=%s got=%s Address did base58 encode/decode correctly.", ii, vv, chk)
			}
		}
	})
	t.Run("encode-decode-trivial", func(t *testing.T) {
		encoded := TrivialBase58EncodingAlphabet([]byte("hello world from base58"), BTCAlphabet)
		assert.Equal(t, encoded, "3A4aJRHr8F4zh2qnGZB9uZctijAxYtPh")
		t.Log(encoded)
	})
	// -gcflags="-d=ssa/check_bce/debug=1"
	t.Run("encode-decode-fast", func(t *testing.T) {
		encoded := FastBase58Encoding([]byte("hello world from base58"))
		assert.Equal(t, encoded, "3A4aJRHr8F4zh2qnGZB9uZctijAxYtPh")
		t.Log(encoded)
	})
	t.Run("encode-cpu-profile", func(t *testing.T) {
		// CPU profiling by default
		defer profile.Start().Stop()
		src := "3A4aJRHr8F4zh2qnGZB9uZctijAxYtPh"
		msg := []byte("hello world from base58")
		for i := 0; i < 10000000; i++ {
			encoded := FastBase58Encoding(msg)
			if encoded == src {

			}
		}
	})
	t.Run("encode-cpu-profile-2", func(t *testing.T) {
		// CPU profiling by default
		defer profile.Start().Stop()
		msg := []byte("hello world from base58")
		for i := 0; i < 10000000; i++ {
			_ = FastBase58Encoding2(msg)
		}
	})
	t.Run("encode-mem-profile", func(t *testing.T) {
		// CPU profiling by default
		defer profile.Start(profile.MemProfile, profile.MemProfileRate(1000)).Stop()
		src := "3A4aJRHr8F4zh2qnGZB9uZctijAxYtPh"
		msg := []byte("hello world from base58")
		for i := 0; i < 10000000; i++ {
			encoded := FastBase58Encoding(msg)
			if encoded == src {

			}
		}
	})
}
