package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode // starting prev is nil

	for curr != nil {
		next := curr.Next // keep a pointer to the next node
		curr.Next = prev  // reverse the link to prev
		prev = curr       // current node becomes the new prev
		curr = next       // next becomes current
	}

	return prev // prev is the head since curr is nil
}
