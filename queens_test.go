package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueens(t *testing.T) {
	assert.Equal(t, [][]string{
		[]string{
			".Q..",
			"...Q",
			"Q...",
			"..Q.",
		},
		[]string{
			"..Q.",
			"Q...",
			"...Q",
			".Q..",
		},
	}, solveNQueens(4))

	assert.Len(t, solveNQueens(8), 92)
}
