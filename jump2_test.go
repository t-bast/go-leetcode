package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump2(t *testing.T) {
	assert.Equal(t, 0, jump([]int{2}))
	assert.Equal(t, 1, jump([]int{1, 2}))
	assert.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, 3, jump([]int{3, 2, 2, 2, 2, 2, 2, 1}))
	assert.Equal(t, 2, jump([]int{2, 3, 0, 1, 4}))
	assert.Equal(t, 2, jump([]int{3, 1, 0, 1, 4}))
	assert.Equal(t, 1, jump([]int{3, 2, 1}))
	assert.Equal(t, 1, jump([]int{2, 2, 1}))

	stackBlower := make([]int, 1024)
	stackBlower[0] = 25
	for i := 1; i < len(stackBlower); i++ {
		stackBlower[i] = 1024 - i
	}

	assert.Equal(t, 2, jump(stackBlower))

	longTest := make([]int, 25000)
	for i := 0; i < len(longTest); i++ {
		longTest[i] = 10
	}

	assert.Equal(t, 2500, jump(longTest))
}
