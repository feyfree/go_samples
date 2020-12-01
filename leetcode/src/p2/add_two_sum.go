package p2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	result := &ListNode{}
	tempHead := result
	var flag = 0
	for l1 != nil && l2 != nil {
		tempSum := l1.Val + l2.Val + flag
		tempResult := tempSum % 10
		flag = tempSum / 10
		data := &ListNode{}
		data.Val = tempResult
		tempHead.Next = data
		l1 = l1.Next
		l2 = l2.Next
		tempHead = tempHead.Next
	}
	for l1 != nil {
		tempSum := l1.Val + flag
		tempResult := tempSum % 10
		flag = tempSum / 10
		data := &ListNode{}
		data.Val = tempResult
		tempHead.Next = data
		l1 = l1.Next
		tempHead = tempHead.Next
	}
	for l2 != nil {
		tempSum := l2.Val + flag
		tempResult := tempSum % 10
		flag = tempSum / 10
		data := &ListNode{}
		data.Val = tempResult
		tempHead.Next = data
		l2 = l2.Next
		tempHead = tempHead.Next
	}
	if flag == 1 {
		data := &ListNode{}
		data.Val = 1
		tempHead.Next = data
	}
	return result.Next
}
