package leetcode

// https://leetcode.com/problems/longest-substring-without-repeating-characters

func lengthOfLongestSubstring(s string) int {
	res := 0

	for i := range s {
		seen := map[rune]struct{}{}
		l := 0

		for _, c := range s[i:] {
			_, ok := seen[c]
			if ok {
				if l > res {
					res = l
				}

				break
			}

			seen[c] = struct{}{}
			l++
		}

		if l > res {
			res = l
		}
	}

	return res
}
