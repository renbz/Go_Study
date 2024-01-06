package main

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curPre := head
	for curPre.Next != nil {
		curPre.Next = &ListNode{getGcd(curPre.Val, curPre.Next.Val), curPre.Next}
		curPre = curPre.Next.Next
	}
	return head
}

func getGcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

type ListNode2807 struct {
	Val  int
	Next *ListNode
}
