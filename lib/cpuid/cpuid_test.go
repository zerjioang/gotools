package cpuid

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCPuId(t *testing.T) {
	t.Run("get-features", func(t *testing.T) {
		f := GetCpuFeatures()
		assert.NotNil(t, f)
		t.Log(f)
	})
	t.Run("get-features-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				f := GetCpuFeatures()
				assert.NotNil(t, f)
				g.Done()
			}()
		}
		g.Wait()
	})
}
