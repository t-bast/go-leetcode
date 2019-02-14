package leetcode

// https://leetcode.com/problems/median-of-two-sorted-arrays/

import "fmt"

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)

	// Initially, split in half.
	i1 := n1 / 2
	i2 := n2 / 2

	// Compensate potential imbalance if both arrays have odd length.
	if n1%2 == 1 && n2%2 == 1 {
		if n1 < n2 {
			i2++
		} else {
			i1++
		}
	}

	// Then iteratively adapt the split.
	for {
		// Verify that the split is correct.
		sizeDiff := len(nums1[:i1]) + len(nums2[:i2]) - len(nums1[i1:]) - len(nums2[i2:])
		if sizeDiff < -1 || 1 < sizeDiff {
			panic(fmt.Sprintf("incorrect split: n1=%d, n2=%d, i1=%d, i2=%d", n1, n2, i1, i2))
		}

		// If we have correctly split the arrays, we can obtain the median.
		cond12 := i1-1 < 0 || i2 >= n2 || nums1[i1-1] <= nums2[i2]
		cond21 := i2-1 < 0 || i1 >= n1 || nums2[i2-1] <= nums1[i1]
		if cond12 && cond21 {
			var left, right int

			if i1 >= 1 && i2 >= 1 {
				left = max(nums1[i1-1], nums2[i2-1])
			} else if i1 >= 1 {
				left = nums1[i1-1]
			} else if i2 >= 1 {
				left = nums2[i2-1]
			}

			if i1 < n1 && i2 < n2 {
				right = min(nums1[i1], nums2[i2])
			} else if i1 < n1 {
				right = nums1[i1]
			} else if i2 < n2 {
				right = nums2[i2]
			}

			if sizeDiff == 0 {
				return float64(left+right) / 2
			} else if sizeDiff < 0 {
				return float64(right)
			} else {
				return float64(left)
			}
		}

		// Otherwise we adjust the split.
		// We could (should?) do better than an offset of 1 if we want a log complexity.
		// We could do something like smallest set length / 2 or something like it, but we
		// must do a lot of annoying bounds checks.
		if !cond12 {
			i1--
			i2++
		} else {
			i1++
			i2--
		}
	}
}
