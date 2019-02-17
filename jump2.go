package leetcode

type move struct {
	start int
	val   int
}

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	if len(nums) == 2 {
		return 1
	}

	// DFS with back-tracking (instead of recursion) so that we don't blow up
	// memory usage / stack size.

	bestScore := 0
	position := 0
	var moves []move

	for {
		cur := nums[position]

		// If we can reach the last cell, we have a solution.
		if position+cur >= len(nums)-1 {
			score := 1 + len(moves)
			if bestScore == 0 || score < bestScore {
				bestScore = score
			}

			// Backtrack one move.
			lastMove := moves[len(moves)-1]
			position = lastMove.start
			continue
		}

		// If this is the first move, let's try moving forward.
		if len(moves) == 0 {
			moves = append(moves, move{start: position, val: 1})
			position++
			continue
		}

		lastMove := moves[len(moves)-1]

		// Backtrack if needed.
		if position <= lastMove.start {
			moves = moves[:len(moves)-1]

			// If we exhausted the current position's possibilities, continue
			// back-tracking.
			if lastMove.val == cur {
				// TODO: we probably break the for loop here.
			} else {
				position += lastMove.val + 1
				moves = append(moves, move{start: position, val: lastMove.val + 1})
				continue
			}
		}

		// TODO: select next move and apply it (just +1).
	}

	return bestScore
}
