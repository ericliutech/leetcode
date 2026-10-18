package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	return reverseNext(head, head.Next)
}

func reverseNext(curr, next *ListNode) *ListNode {
	if next == nil {
		return curr
	}

	reversed := reverseNext(next, next.Next)
	next.Next = curr
	curr.Next = nil

	return reversed
}
