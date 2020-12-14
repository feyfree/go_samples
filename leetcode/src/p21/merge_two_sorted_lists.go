package p21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
			temp = temp.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
			temp = temp.Next
		}
	}
	for l1 != nil {
		temp.Next = l1
		break
	}
	for l2 != nil {
		temp.Next = l2
		break
	}
	return result.Next
}
