<p align="center">
  <img alt="gotools logo" src="https://user-images.githubusercontent.com/6706342/58815376-bf217c00-8627-11e9-87a4-00e4370502b8.png" width="200px"></img>
  <h3 align="center"><b>gotools</b></h3>
  <p align="center">High Performance Computing Go</p>
</p>

<p align="center">
    <a href="https://coveralls.io/r/zerjioang/gotools">
      <img alt="Covergage" src="https://img.shields.io/coveralls/zerjioang/gotools.svg?style=flat-square">
    </a>
    <a href="https://travis-ci.org/zerjioang/gotools">
      <img alt="Build Status" src="https://travis-ci.org/zerjioang/gotools.svg?branch=master">
    </a>
    <a href="https://goreportcard.com/report/github.com/zerjioang/gotools">
       <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/zerjioang/gotools">
    </a>
    <a href="https://github.com/zerjioang/gotools/blob/master/LICENSE">
        <img alt="Software License" src="http://img.shields.io/:license-gpl3-brightgreen.svg?style=flat-square">
    </a>
    <a href="https://godoc.org/github.com/zerjioang/gotools">
       <img alt="Go Docs" src="https://godoc.org/github.com/zerjioang/gotools?status.svg">
    </a>
</p>

## Install

Package `gotools` is a set of multipurpose libraries modified for **High Performance Computing** in **Go**

## Install

```bash
go get github.com/zerjioang/gotools
```

## Features

Following a list of implemented features are described

| Feature               | Description                                                                                   |
|-----------------------|-----------------------------------------------------------------------------------------------|
| `IsIpv4`              | Ipv4 string address validator                                                                 |
| `Ip2intLow`           | Converts given IP address string to its numeric representation                                |
| `IpToIntAssemblyAmd64` | Converts given IP address string to its numeric representation using `amd64` Go assembly code |

## Other Libraries

* A deadly simple thread-safe, zero-alloc event bus for Golang (https://github.com/zerjioang/go-bus)
* A pure go High Performance Finite State Machine with GraphViz support (https://github.com/zerjioang/go-fsm)

## Development 

### Inline anaylsis

```bash
go test ./... -gcflags="-m=2" 2>&1 | grep "too complex"
```

### Bound check analysis

```bash
go test ./... -gcflags="-d=ssa/check_bce/debug=1" 
```

You can do an early bounds check deep into the slice to avoid multiple bounds checks further on.
If that won't work:

* propagate constants
* unroll loops
* reuse previously allocated vars

This can make the code super ugly! a dozen lines becomes a few hundred. but... 80% throughput boost!

## Contributing

Everybody is welcome to contribute to **gotools**. And any other comments will be very appreciate.

## License

All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

 * Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
 * Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 * Uses GPL license described below

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
