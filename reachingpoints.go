package leetcode

// https://leetcode.com/contest/weekly-contest-71/problems/reaching-points/

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	// We reached the point. Yay!
	if sx == tx && sy == ty {
		return true
	}

	// We missed it.
	if sx > tx {
		return false
	}
	if sy > ty {
		return false
	}

	/* Solution 1 */

	// This simple recursion goes out of memory very fast.
	// return reachingPoints(sx, sy+sx, tx, ty) || reachingPoints(sx+sy, sy, tx, ty)

	/* Solution 2 */

	// Instead we kind of linearize the problem. We represent a path by a number,
	// where each bit represents the kind of move taken.
	// We try all solutions, but this shouldn't use memory this time.

	// The tricky part is to evaluate the longest possible path to bound the
	// number of attempts.
	// This is currently not working.
	maxCount := (ty-sy)/sx + (tx-sx)/sy

	for i := uint64(0); i < 2<<uint(maxCount); i++ {
		curx := sx
		cury := sy
		k := i
		for {
			if k%2 == 0 {
				cury += curx
			} else {
				curx += cury
			}

			// We reached the point. Yay!
			if curx == tx && cury == ty {
				return true
			}

			// We missed it.
			if curx > tx {
				break
			}
			if cury > ty {
				break
			}

			// Next move.
			k = k >> 1
		}
	}

	return false
}
