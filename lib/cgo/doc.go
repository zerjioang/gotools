package cgo

// Initial performance
//
// BenchmarkCgoVersion/cgo-version-4   9929176	       121 ns/op	    8.29 MB/s	      16 B/op	       1 allocs/op
//
// Since gcc version wont change (under normal conditions) we cache the results for better performance
//
// BenchmarkCgoVersion/cgo-version-4   1000000000	 0.378 ns/op	 2643.16 MB/s	       0 B/op	       0 allocs/op
//
// buf if yor any reason changes, method FlushCache() will reload its new values
