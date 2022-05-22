package leetcode

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur,nxt:=head,head
	for cur!=nil{
		nxt=cur.Next
		cur.Next=pre
		pre=cur
		cur=nxt
	}
	return pre
}