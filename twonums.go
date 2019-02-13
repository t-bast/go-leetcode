package leetcode

// https://leetcode.com/problems/add-two-numbers/

// ListNode linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	currentRes := res
	current1 := l1
	current2 := l2
	carry := 0

	for {
		currentRes.Val = carry
		if current1 != nil {
			currentRes.Val += current1.Val
		}
		if current2 != nil {
			currentRes.Val += current2.Val
		}

		carry = 0
		if currentRes.Val > 9 {
			carry = 1
			currentRes.Val = currentRes.Val % 10
		}

		if current1 != nil {
			current1 = current1.Next
		}
		if current2 != nil {
			current2 = current2.Next
		}

		if current1 == nil && current2 == nil {
			if carry == 1 {
				currentRes.Next = &ListNode{Val: 1}
			}

			break
		}

		currentRes.Next = &ListNode{}
		currentRes = currentRes.Next
	}

	return res
}
