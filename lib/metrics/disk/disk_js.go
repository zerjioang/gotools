// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

// +build !wasm !js

package disk

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type DiskStatus struct {
	all  uint64
	used uint64
	free uint64
}

// constructor like function
func DiskUsage() DiskStatus {
	disk := DiskStatus{}
	return disk
}

func DiskUsagePtr() *DiskStatus {
	d := DiskUsage()
	return &d
}

// disk usage of path/disk
func (disk *DiskStatus) Start(path string) {
}

// internal ticker based monitor
func (disk *DiskStatus) monitor(path string) {
}

func (disk *DiskStatus) read(path string) error {
	disk.all = 0
	disk.free = 0
	disk.used = 0
	return nil
}

// get all value
func (disk *DiskStatus) All() uint64 {
	raw := disk.all / GB
	return raw
}

// get used value
func (disk *DiskStatus) Used() uint64 {
	raw := disk.used / GB
	return raw
}

// get free value
func (disk *DiskStatus) Free() uint64 {
	raw := disk.free / GB
	return raw
}

func (disk DiskStatus) IsMonitoring() bool {
	return true
}
