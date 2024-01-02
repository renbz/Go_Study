package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeNodes(head.Next)
	if head.Next != nil && head.Next.Val > head.Val {
		return head.Next
	}
	return head
}
