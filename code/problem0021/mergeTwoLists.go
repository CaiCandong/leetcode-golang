package problem0021

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	res := &ListNode{}
	p := res
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p = p.Next
			p1 = p1.Next
		} else {
			p.Next = p2
			p = p.Next
			p2 = p2.Next
		}
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return res
}
