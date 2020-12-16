package p25

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	count := 0
	for true {
		loop := pre
		for loop != nil {
			loop = loop.Next
			count++
		}
		if count <= k {
			break
		}
		// 置换的核心思想是 节点前移
		cur := pre.Next
		nxt := cur.Next
		for i := 1; i < k && nxt != nil; i++ {
			cur.Next = nxt.Next
			nxt.Next = pre.Next
			pre.Next = nxt
			nxt = cur.Next
		}
		pre = cur
		count = 0
	}
	return dummy.Next
}
