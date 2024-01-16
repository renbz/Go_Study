package main

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ans := &ListNode{0, head}
	curPre := ans
	// 下一个不为null
	for curPre.Next != nil && curPre.Next.Next != nil {
		// 下一个和下下个不相等
		if curPre.Next.Val == curPre.Next.Next.Val {
			x := curPre.Next.Val
			for curPre.Next != nil && curPre.Next.Val == x {
				curPre.Next = curPre.Next.Next
			}
		} else {
			curPre = curPre.Next
		}
	}
	return ans.Next
}
