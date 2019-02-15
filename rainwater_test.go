package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRainWater(t *testing.T) {
	t.Run("empty input", func(t *testing.T) {
		assert.Equal(t, 0, trap([]int{}))
	})

	t.Run("tiny input", func(t *testing.T) {
		assert.Equal(t, 0, trap([]int{3}))
	})

	t.Run("small step", func(t *testing.T) {
		assert.Equal(t, 0, trap([]int{1, 2}))
	})

	t.Run("full leak on both sides", func(t *testing.T) {
		heights := []int{0, 1, 1, 2, 3, 3, 2, 2, 1}
		assert.Equal(t, 0, trap(heights))
	})

	t.Run("small leak on both sides", func(t *testing.T) {
		heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		assert.Equal(t, 6, trap(heights))
	})

	t.Run("left-only leak", func(t *testing.T) {
		heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2}
		assert.Equal(t, 6, trap(heights))
	})

	t.Run("right-only leak", func(t *testing.T) {
		heights := []int{1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		assert.Equal(t, 6, trap(heights))
	})
}
