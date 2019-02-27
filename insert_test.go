package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertInterval(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		intervals := insert([]Interval{NewInterval(1, 3), NewInterval(6, 9)}, NewInterval(2, 5))
		assert.ElementsMatch(t, []Interval{NewInterval(1, 5), NewInterval(6, 9)}, intervals)
	})

	t.Run("example 2", func(t *testing.T) {
		intervals := insert([]Interval{
			NewInterval(1, 2),
			NewInterval(3, 5),
			NewInterval(6, 7),
			NewInterval(8, 10),
			NewInterval(12, 16),
		}, NewInterval(4, 8))
		assert.ElementsMatch(t, []Interval{
			NewInterval(1, 2),
			NewInterval(3, 10),
			NewInterval(12, 16),
		}, intervals)
	})

	t.Run("all before", func(t *testing.T) {
		intervals := insert([]Interval{
			NewInterval(1, 2),
			NewInterval(3, 5),
			NewInterval(6, 7),
		}, NewInterval(8, 9))
		assert.ElementsMatch(t, []Interval{
			NewInterval(1, 2),
			NewInterval(3, 5),
			NewInterval(6, 7),
			NewInterval(8, 9),
		}, intervals)
	})

	t.Run("all after", func(t *testing.T) {
		intervals := insert([]Interval{
			NewInterval(3, 5),
			NewInterval(6, 7),
		}, NewInterval(1, 2))
		assert.ElementsMatch(t, []Interval{
			NewInterval(1, 2),
			NewInterval(3, 5),
			NewInterval(6, 7),
		}, intervals)
	})

	t.Run("merge beginning", func(t *testing.T) {
		intervals := insert([]Interval{
			NewInterval(1, 3),
			NewInterval(4, 7),
			NewInterval(9, 10),
		}, NewInterval(2, 5))
		assert.ElementsMatch(t, []Interval{
			NewInterval(1, 7),
			NewInterval(9, 10),
		}, intervals)
	})

	t.Run("merge ending", func(t *testing.T) {
		intervals := insert([]Interval{
			NewInterval(1, 2),
			NewInterval(4, 5),
			NewInterval(6, 7),
		}, NewInterval(4, 8))
		assert.ElementsMatch(t, []Interval{
			NewInterval(1, 2),
			NewInterval(4, 8),
		}, intervals)
	})
}
