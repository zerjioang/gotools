package example_test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	/*
	Server Software:
	Server Hostname:        localhost
	Server Port:            8080

	Document Path:          /
	Document Length:        20 bytes

	Concurrency Level:      500
	Time taken for tests:   0.129 seconds
	Complete requests:      1000
	Failed requests:        0
	Total transferred:      136000 bytes
	HTML transferred:       20000 bytes
	Requests per second:    7770.67 [#/sec] (mean)
	Time per request:       64.344 [ms] (mean)
	Time per request:       0.129 [ms] (mean, across all concurrent requests)
	Transfer rate:          1032.04 [Kbytes/sec] received

	Connection Times (ms)
	              min  mean[+/-sd] median   max
	Connect:       12   22   6.3     22      38
	Processing:     9   30  10.9     29      57
	Waiting:        6   21  10.2     20      41
	Total:         37   51   8.1     54      70

	Percentage of the requests served within a certain time (ms)
	  50%     54
	  66%     56
	  75%     58
	  80%     59
	  90%     61
	  95%     64
	  98%     67
	  99%     68
	 100%     70 (longest request)
	*/
	t.Run("standard-server", func(t *testing.T) {
		http.HandleFunc("/", HelloServer)
		_ = http.ListenAndServe(":8080", nil)
	})
}


func HelloServer(w http.ResponseWriter, r *http.Request) {
	_,_=fmt.Fprint(w, "<h1>hello world</h1>")
}
