package p23

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	n := len(lists)
	if n == 1 {
		return lists[0]
	}
	mid := n / 2
	pre := lists[:mid]
	after := lists[mid:]
	preData := mergeKLists(pre)
	afterData := mergeKLists(after)
	combination := []*ListNode{preData, afterData}
	return mergeKLists(combination)
}

func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
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
	}
	for l2 != nil {
		temp.Next = l2
	}
	return result.Next
}
