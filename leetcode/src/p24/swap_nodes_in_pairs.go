package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for cur != nil && cur.Next != nil {
		nodeA := cur.Next
		nodeB := cur.Next.Next
		if nodeB != nil {
			nodeA.Next = nodeB.Next.Next
			cur.Next = nodeB
			nodeB.Next = nodeA
		} else {
		}
		cur = cur.Next.Next
	}
	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1}
	head1 := &ListNode{Val: 2}
	head2 := &ListNode{Val: 3}
	head3 := &ListNode{Val: 4}
	head2.Next = head3
	head1.Next = head2
	head.Next = head1
	fmt.Println(swapPairs(head))
}
