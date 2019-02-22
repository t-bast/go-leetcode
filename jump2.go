package leetcode

// https://leetcode.com/problems/jump-game-ii/

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

	if nums[0] >= len(nums)-1 {
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
				// If we can't backtrack more, that means we tried everything.
				if len(moves) == 0 {
					break
				}

				// Otherwise we keep on backtracking.
				lastMove = moves[len(moves)-1]
				position = lastMove.start
				continue
			} else {
				moves = append(moves, move{start: position, val: lastMove.val + 1})
				position += lastMove.val + 1
				continue
			}
		}

		// If stuck, backtrack one move.
		if cur == 0 {
			// Backtrack one move.
			position = lastMove.start
			continue
		}

		// Otherwise we should keep going forward, starting with 1.
		moves = append(moves, move{start: position, val: 1})
		position++
	}

	return bestScore
}
