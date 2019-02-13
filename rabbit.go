package leetcode

// https://leetcode.com/contest/weekly-contest-71/problems/rabbits-in-forest/

func numRabbits(answers []int) int {
	if len(answers) == 0 {
		return 0
	}

	if len(answers) == 1 {
		return answers[0] + 1
	}

	total := 0
	buckets := make(map[int]int)

	for _, answer := range answers {
		// if answer == 0 {
		// 	total++
		// 	continue
		// }

		count, ok := buckets[answer]
		if !ok {
			buckets[answer] = 1
			continue
		}

		if count < answer+1 {
			// Bucket not full, keep adding to it.
			buckets[answer]++
		} else {
			// Bucket is full, we need to reset it.
			total += count
			buckets[answer] = 1
		}
	}

	for answer := range buckets {
		total += answer + 1
	}

	return total
}
