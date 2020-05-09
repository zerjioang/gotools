package cpu

import (
	"fmt"
	"testing"
	"time"
)

func TestCpu(t *testing.T) {
	t.Run("instant-sample", func(t *testing.T) {
		idle0, total0 := getCPUSample()
		fmt.Printf("CPU usage [busy: %d, total: %d]\n", idle0, total0)
	})
	t.Run("1-second-sample", func(t *testing.T) {
		idle0, total0 := getCPUSample()
		time.Sleep(1 * time.Second)
		idle1, total1 := getCPUSample()

		idleTicks := float64(idle1 - idle0)
		totalTicks := float64(total1 - total0)
		cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

		fmt.Printf("CPU usage is %f%% [busy: %f, total: %f]\n", cpuUsage, totalTicks-idleTicks, totalTicks)
	})
}
