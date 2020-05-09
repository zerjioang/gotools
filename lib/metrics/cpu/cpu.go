package cpu

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/zerjioang/gotools/lib/logger"
)

const (
	source = "/proc/stat"
)

// The very first "cpu" line aggregates the numbers in all of the other "cpuN" lines.
// These numbers identify the amount of time the CPU has spent performing different kinds of work.
// Time units are in USER_HZ or Jiffies (typically hundredths of a second).
// The meanings of the columns are as follows, from left to right:
//
//    user: normal processes executing in user mode
//    nice: niced processes executing in user mode
//    system: processes executing in kernel mode
//    idle: twiddling thumbs
//    iowait: waiting for I/O to complete
//    irq: servicing interrupts
//    softirq: servicing softirqs
//
// * The "intr" line gives counts of interrupts serviced since boot time, for each of the possible system interrupts. The first column is the total of all interrupts serviced; each subsequent column is the total for that particular interrupt.
// * The "ctxt" line gives the total number of context switches across all CPUs.
// * The "btime" line gives the time at which the system booted, in seconds since the Unix epoch.
// * The "processes" line gives the number of processes and threads created, which includes (but is not limited to) those created by calls to the fork() and clone() system calls.
// * The "procs_running" line gives the number of processes currently running on CPUs.
// * The "procs_blocked" line gives the number of processes currently blocked, waiting for I/O to complete.
func getCPUSample() (idle, total uint64) {
	contents, err := ioutil.ReadFile(source)
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					logger.Error("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}
