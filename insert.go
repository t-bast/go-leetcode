package leetcode

// https://leetcode.com/problems/insert-interval/

// Interval between two ints.
type Interval struct {
	Start int
	End   int
}

// NewInterval with the given bounds.
func NewInterval(start, end int) Interval {
	return Interval{
		Start: start,
		End:   end,
	}
}

/**
 * Definition for an interval.
 */
func insert(intervals []Interval, newInterval Interval) []Interval {
	var results []Interval

	i := 0

	// Strictly before - don't modify.
	for {
		if i == len(intervals) {
			results = append(results, newInterval)
			return results
		}

		cur := intervals[i]
		if cur.End < newInterval.Start {
			results = append(results, cur)
		} else {
			break
		}

		i++
	}

	// Define new start.
	if intervals[i].Start < newInterval.Start {
		newInterval.Start = intervals[i].Start
	}

	// Merge.
	for {
		if i == len(intervals) {
			results = append(results, newInterval)
			return results
		}

		cur := intervals[i]

		if newInterval.End < cur.Start {
			results = append(results, newInterval)
			break
		}

		if cur.End > newInterval.End {
			newInterval.End = cur.End
		}

		i++
	}

	// Strictly after - don't modify
	for {
		if i == len(intervals) {
			return results
		}

		cur := intervals[i]
		if newInterval.End < cur.Start {
			results = append(results, cur)
		} else {
			break
		}

		i++
	}

	return results
}
