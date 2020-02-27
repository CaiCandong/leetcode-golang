package problem0009

type ListNode struct {
	Val  int
	Next *ListNode
}

// previous 指向第n+1个元素
// now 指向第1个元素
//1,2,3,4,5 2
// preivous->3 now->1
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	previous, now := head, head
	for n != 0 {
		previous = previous.Next
		n--
	}
	if previous == nil {
		return head.Next
	}
	for previous.Next != nil {
		previous = previous.Next
		now = now.Next
	}
	now.Next = now.Next.Next
	return head
}
