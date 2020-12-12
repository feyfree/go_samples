package p19

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	head = dummy
	for n > 0 {
		head = head.Next
		n--
	}
	endNode := dummy
	for head.Next != nil {
		head = head.Next
		endNode = endNode.Next
	}
	endNode.Next = endNode.Next.Next
	return dummy.Next
}
