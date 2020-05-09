// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package model

import (
	"strconv"

	"github.com/zerjioang/gotools/common"
	"github.com/zerjioang/gotools/lib/cpuid"
)

type ServerStatusResponse struct {
	Disk   Disk   `json:"disk"`
	Memory Memory `json:"memory"`
	Gc     Gc     `json:"gc"`
}

func (r *ServerStatusResponse) Reset() {
	r.Disk.Reset()
	r.Memory.Reset()
	r.Gc.Reset()
}

type ServerInfoResponse struct {
	Architecture string  `json:"arch"`
	Os           string  `json:"os"`
	Cpus         Cpus    `json:"cpus"`
	Runtime      Runtime `json:"runtime"`
	Version      Version `json:"version"`
}

type Cpus struct {
	Cores    int               `json:"cores"`
	Features cpuid.CpuFeatures `json:"features"`
}
type Runtime struct {
	Compiler string `json:"compiler"`
}
type Version struct {
	AppName string `json:"app_name"`
	HTTP      string `json:"http"`
	Go        string `json:"go"`
}
type Disk struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

func (disk *Disk) Reset() {
	disk.All = 0
	disk.Used = 0
	disk.Free = 0
}

type Memory struct {
	Frees     uint64 `json:"frees"`
	Heapalloc uint64 `json:"heapalloc"`
	Mallocs   uint64 `json:"mallocs"`
	Alloc     uint64 `json:"alloc"`
	Total     uint64 `json:"total"`
	Sys       uint64 `json:"sys"`
	App       struct {
		Alloc      uint64 `json:"allocMb"`
		TotalAlloc uint64 `json:"totalMb"`
		Sys        uint64 `json:"sysMb"`
	} `json:"app"`
}

func (memory *Memory) Reset() {
	memory.Frees = 0
	memory.Heapalloc = 0
	memory.Alloc = 0
	memory.Total = 0
	memory.Sys = 0
	memory.Mallocs = 0
	// reset app stats
	memory.App.Alloc = 0
	memory.App.TotalAlloc = 0
	memory.App.Sys = 0
}

type Gc struct {
	Numgc       uint32 `json:"numgc"`
	NumForcedGC uint32 `json:"numForcedGC"`
}

func (gc *Gc) Reset() {
	gc.Numgc = 0
	gc.NumForcedGC = 0
}

func (r ServerInfoResponse) Bytes(serializer common.Serializer) []byte {
	raw, _ := serializer(r)
	return raw
}

func itoa(v int) string {
	return strconv.Itoa(v)
}

func itoau32(v uint32) string {
	return strconv.Itoa(int(v))
}

func itoau64(v uint64) string {
	return strconv.Itoa(int(v))
}

func tofloat64(v float64) string {
	//return fmt.Sprintf("%.6f", v)
	return strconv.FormatFloat(v, 'f', 6, 64)
}
