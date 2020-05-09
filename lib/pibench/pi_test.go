package pibench

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonteCarlo(t *testing.T) {
	t.Run("calculate-score", func(t *testing.T) {
		calculateScore()
	})
	t.Run("calculate-score-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 20
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				calculateScore()
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("get-score", func(t *testing.T) {
		calculateScore()
		v := GetScore()
		assert.NotNil(t, v)
		assert.True(t, v > 0)
	})
	t.Run("get-score-goroutines", func(t *testing.T) {
		// calculate score once
		calculateScore()
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				v := GetScore()
				assert.NotNil(t, v)
				assert.True(t, v > 0)
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("get-pibench-time", func(t *testing.T) {
		calculateScore()
		v := GetBenchTime().Nanoseconds()
		assert.NotNil(t, v)
		assert.True(t, v > 0)
	})
	t.Run("get-benchtime-goroutines", func(t *testing.T) {
		// calculate score once
		calculateScore()
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				v := GetBenchTime()
				assert.NotNil(t, v)
				assert.True(t, v > 0)
				g.Done()
			}()
		}
		g.Wait()
	})
}
