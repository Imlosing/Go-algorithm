package leetcode

// ListNode Definition for singly-linked list.
 type ListNode struct {
 	Val int
 	Next *ListNode
 }

func reverseKGroup(head *ListNode, k int) *ListNode {
	left,right:=head,head
	for i:=0;i<k;i++{
		if right==nil{
			return head
		}
		right=right.Next
	}
	newHead:=reverse(left,right)
	left.Next=reverseKGroup(right,k)
	return newHead
}

func reverse(left,right *ListNode)*ListNode{
	var pre *ListNode
	cur,nxt:=left,right
	for cur!=right{
		nxt=cur.Next
		cur.Next=pre
		pre=cur
		cur=nxt
	}
	return pre
}