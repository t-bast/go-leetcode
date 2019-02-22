package leetcode

// https://leetcode.com/problems/jump-game-ii/

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

	jumps := 0
	end := 0
	furthest := 0

	// Since we are allowed to jump further than the end, our goal is to go as
	// far as possible at each step.
	for i := 0; i < len(nums)-1; i++ {
		// For each position we evaluate if it can get us further than our last
		// known jump.
		// If it does we can consider than we're using it (because if we're
		// able to go further than it, we're also able to go to it).
		if i+nums[i] > furthest {
			furthest = i + nums[i]
		}

		// If we reach the end of our current slice, we will have to make a
		// jump and it extends the range we can reach in `jumps` jumps.
		if i == end {
			jumps++
			end = furthest
		}
	}

	return jumps
}
