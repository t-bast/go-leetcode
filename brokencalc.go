package leetcode

func brokenCalc(X int, Y int) int {
	ops := 0

	for {
		if X >= Y {
			return ops + X - Y
		}

		// If Y is even, we try to reach Y/2 then we double X.
		if Y%2 == 0 {
			Y = Y / 2
		} else {
			// If Y is odd, we try to reach Y+1 then we decrement.
			Y++
		}

		ops++
	}
}
