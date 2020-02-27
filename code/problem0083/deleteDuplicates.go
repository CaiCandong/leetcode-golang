package problem0083

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针解法
func deleteDuplicates(head *ListNode) *ListNode {
	now, next := head, head
	for next != nil {
		for next != nil && next.Val == now.Val {
			next = next.Next
		}
		now.Next = next
		now = next
	}
	return head
}
