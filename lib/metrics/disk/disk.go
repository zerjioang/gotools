// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

// +build !wasm !js

package disk

import (
	"sync"
	"syscall"
	"time"

	"github.com/zerjioang/gotools/lib/logger"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

var (
	//shared coordinator for all disk status reader
	ticker = time.NewTicker(5 * time.Second)
)

type DiskStatus struct {
	all  uint64
	used uint64
	free uint64
	lock *sync.Mutex
}

// constructor like function
func DiskUsage() DiskStatus {
	disk := DiskStatus{}
	disk.lock = new(sync.Mutex)
	return disk
}

func DiskUsagePtr() *DiskStatus {
	d := DiskUsage()
	return &d
}

// disk usage of path/disk
func (disk *DiskStatus) Start(path string) {
	// initialize read values
	err := disk.read(path)
	if err != nil {
		logger.Error("failed to read disk statistics", err)
	}
	// start monitor
	go disk.monitor(path)
}

// internal ticker based monitor
func (disk *DiskStatus) monitor(path string) {
	for range ticker.C {
		// logger.Debug("reading node disk space statistics")
		rErr := disk.read(path)
		if rErr != nil {
			logger.Error("disk status read error", rErr)
		}
	}
}

func (disk *DiskStatus) read(path string) error {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return err
	}
	disk.lock.Lock()
	disk.all = fs.Blocks * uint64(fs.Bsize)
	disk.free = fs.Bfree * uint64(fs.Bsize)
	disk.used = disk.all - disk.free
	disk.lock.Unlock()
	return nil
}

// get all value
func (disk *DiskStatus) All() uint64 {
	disk.lock.Lock()
	raw := disk.all / GB
	disk.lock.Unlock()
	return raw
}

// get used value
func (disk *DiskStatus) Used() uint64 {
	disk.lock.Lock()
	raw := disk.used / GB
	disk.lock.Unlock()
	return raw
}

// get free value
func (disk *DiskStatus) Free() uint64 {
	disk.lock.Lock()
	raw := disk.free / GB
	disk.lock.Unlock()
	return raw
}

func (disk DiskStatus) IsMonitoring() bool {
	return true
}
