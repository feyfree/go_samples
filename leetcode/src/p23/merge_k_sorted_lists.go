package p23

type ListNode struct {
	Val  int
	Next *ListNode
}

// [[1,4,5],[1,3,4],[2,6]]
//func main()  {
//	node1 := &ListNode{Val:1}
//	node2 := &ListNode{Val:4}
//	node3 := &ListNode{Val:5}
//	node1.Next = node2
//	node2.Next = node3
//
//	node4 := &ListNode{Val:1}
//	node5 := &ListNode{Val:3}
//	node6 := &ListNode{Val:4}
//	node4.Next = node5
//	node5.Next = node6
//
//	node7 := &ListNode{Val:2}
//	node8 := &ListNode{Val:6}
//	node7.Next = node8
//
//	nodes:= []*ListNode{node1, node4, node7}
//	mergeKLists(nodes)
//}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	n := len(lists)
	if n == 1 {
		return lists[0]
	} else if n == 2 {
		return mergeTwoList(lists[0], lists[1])
	} else {
		mid := n / 2
		pre := lists[:mid]
		after := lists[mid:]
		preData := mergeKLists(pre)
		afterData := mergeKLists(after)
		combination := []*ListNode{preData, afterData}
		return mergeKLists(combination)
	}
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
	if l1 != nil {
		temp.Next = l1
	}
	if l2 != nil {
		temp.Next = l2
	}
	return result.Next
}
