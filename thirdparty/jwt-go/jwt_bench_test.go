package jwt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/pkg/profile"
	"github.com/stretchr/testify/assert"
	"github.com/zerjioang/gotools/thirdparty/jwt-go"
)

var (
	hmacSampleSecret = []byte("foo-bar")
)

func TestJwt(t *testing.T) {
	t.Run("standard-implementation", func(t *testing.T) {
		// Create a new token object, specifying signing method and the claims
		// you would like it to contain.
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"foo": "bar",
			"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(hmacSampleSecret)
		assert.NotNil(t, tokenString)
		assert.NoError(t, err)

		fmt.Println(tokenString, err)
	})
	t.Run("low-level-implementation", func(t *testing.T) {
		//simulate custom claims
		claims := jwt.LowClaims{
			jwt.LowClaim{Key: "foo", Value: "bar"},
			jwt.LowClaim{Key: "nbf", Value: time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix()},
		}
		tokenString, err := jwt.GenerateHS256Jwt(claims, hmacSampleSecret)
		assert.NotNil(t, tokenString)
		assert.NoError(t, err)

		fmt.Println(tokenString, err)
	})
	t.Run("low-level-profile-cpu", func(t *testing.T) {
		// CPU profiling by default
		defer profile.Start().Stop()

		for n := 0; n < 1000000; n++ {
			//simulate custom claims
			claims := jwt.LowClaims{
				jwt.LowClaim{Key: "foo", Value: "bar"},
				jwt.LowClaim{Key: "nbf", Value: time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix()},
			}
			tokenString, err := jwt.GenerateHS256Jwt(claims, hmacSampleSecret)
			if err == nil && tokenString != "" {

			}
		}
	})
	t.Run("low-level-profile-mem", func(t *testing.T) {
		// MEM profiling by default
		defer profile.Start(profile.MemProfile, profile.MemProfileRate(100)).Stop()

		for n := 0; n < 1000000; n++ {
			//simulate custom claims
			claims := jwt.LowClaims{
				jwt.LowClaim{Key: "foo", Value: "bar"},
				jwt.LowClaim{Key: "nbf", Value: time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix()},
			}
			tokenString, err := jwt.GenerateHS256Jwt(claims, hmacSampleSecret)
			if err == nil && tokenString != "" {

			}
		}
	})
}

// BenchmarkJwt/standard-implementation-12         	  300000	      4710 ns/op	   0.21 MB/s	    2809 B/op	      47 allocs/op
func BenchmarkJwt(b *testing.B) {
	b.Run("standard-implementation", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			// Create a new token object, specifying signing method and the claims
			// you would like it to contain.
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"foo": "bar",
				"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
			})

			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString(hmacSampleSecret)
			if err == nil && tokenString != "" {

			}
		}
	})
	b.Run("low-level-implementation", func(b *testing.B) {
		//simulate custom claims
		claims := jwt.LowClaims{
			jwt.LowClaim{Key: "foo", Value: "bar"},
			jwt.LowClaim{Key: "nbf", Value: time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix()},
		}
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			tokenString, err := jwt.GenerateHS256Jwt(claims, hmacSampleSecret)
			if err == nil && tokenString != "" {

			}
		}
	})
}
