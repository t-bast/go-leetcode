package leetcode

// https://leetcode.com/contest/weekly-contest-71/problems/minimum-distance-between-bst-nodes/

// TreeNode is a binary search tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	flat := flatten(root)
	dist := flat[1] - flat[0]
	for i := 2; i < len(flat); i++ {
		cur := flat[i] - flat[i-1]
		if cur < dist {
			dist = cur
		}
	}

	return dist
}

func flatten(root *TreeNode) []int {
	var left, right []int
	if root.Left != nil {
		left = flatten(root.Left)
	}
	if root.Right != nil {
		right = flatten(root.Right)
	}

	return append(
		left,
		append(
			[]int{root.Val},
			right...,
		)...,
	)
}
