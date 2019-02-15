package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrokenCalculator(t *testing.T) {
	t.Run("sample test cases", func(t *testing.T) {
		assert.Equal(t, 2, brokenCalc(2, 3))
		assert.Equal(t, 2, brokenCalc(5, 8))
		assert.Equal(t, 3, brokenCalc(3, 10))
		assert.Equal(t, 5, brokenCalc(2, 10))
		assert.Equal(t, 6, brokenCalc(2, 9))
		assert.Equal(t, 1023, brokenCalc(1024, 1))
		assert.Equal(t, 39, brokenCalc(1, 1000000000))
	})
}
