package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump2(t *testing.T) {
	assert.Equal(t, 0, jump([]int{2}))
	assert.Equal(t, 1, jump([]int{1, 2}))
	assert.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}))

	stackBlower := make([]int, 1024)
	stackBlower[0] = 25
	for i := 1; i < len(stackBlower); i++ {
		stackBlower[i] = 1024 - i
	}

	assert.Equal(t, 7, jump(stackBlower))
}
