package leetcode

// https://leetcode.com/problems/reaching-points/

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	// We "reverse" the problem. We start from the end and we know that we can
	// only come from one point by looking at which coordinate is bigger (since
	// we continuously added coordinates together).

	for {
		if sx == tx && sy == ty {
			return true
		}

		if sx > tx {
			return false
		}

		if sy > ty {
			return false
		}

		if tx > ty {
			tx = tx - ty
		} else {
			ty = ty - tx
		}
	}
}
