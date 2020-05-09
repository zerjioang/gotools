// +build ignore

package wasm

/*

https://blog.gopheracademy.com/advent-2018/go-in-the-browser/

GOOS=js GOARCH=wasm go build -o main.wasm main.go

So now we have the wasm binary generated. But unlike in native systems, we need to run it inside the browser. For this, we need to throw in a few more things to accomplish this:

* A webserver which will serve our web app.
* An index.html file which contains some js glue code needed to load the wasm binary.
* And a js file which serves as the communication interface between the browser and our wasm binary.

$cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
$cp "$(go env GOROOT)/misc/wasm/wasm_exec.html" .
$# we rename the html file to index.html for convenience.
$mv wasm_exec.html index.html
$ls -l

*/
