// Copyright gotools
// SPDX-License-Identifier: GNU GPL v3

// Package jwt is a Go implementation of JSON Web Tokens: http://self-issued.info/docs/draft-jones-json-web-token.html
//
// See README.md for more info.
package jwt

// Initial benchmarking
// BenchmarkJwt/standard-implementation-12         	  300000	      4829 ns/op	   0.21 MB/s	    2809 B/op	      47 allocs/op

// after creating a custom method called GenerateHS256Jwt the benchmark results are
// BenchmarkJwt/low-level-implementation-12        	  300000	      4634 ns/op	   0.22 MB/s	    2425 B/op	      43 allocs/op
// 4 Allocations less

// after removing the marshalling to json of token header (since we are using always hs256 algorithm), the speedup is quite considerable:
// BenchmarkJwt/low-level-implementation-12        	  500000	      3386 ns/op	   0.30 MB/s	    1976 B/op	      32 allocs/op

// after hardocoding token header as base64 content
// and converting dynamic parts array to fixed length array
// BenchmarkJwt/low-level-implementation-12        	  500000	      3270 ns/op	   0.31 MB/s	    1848 B/op	      29 allocs/op

// after replacing custom json.marshalling implementation with low level claims implementation
// BenchmarkJwt/low-level-implementation-12        	 1000000	      2222 ns/op	   0.45 MB/s	    1136 B/op	      21 allocs/op

// after using bytes instead of strings
// BenchmarkJwt/low-level-implementation-12        	 1000000	      2170 ns/op	   0.46 MB/s	    1088 B/op	      19 allocs/op

// after using bytes.buffer for struct serialization insisde Json() method
// BenchmarkJwt/low-level-implementation-12          1000000	      2034 ns/op	   0.49 MB/s	     976 B/op	      15 allocs/op

// after using byte arrays for data conversion and trimming
// BenchmarkJwt/low-level-implementation-12        	 1000000	      2104 ns/op	   0.48 MB/s	    1056 B/op	      13 allocs/op
