package leetcode

type LRUCache struct {
	m map[int]*LinkedNode
	cap int
	head *LinkedNode
	tail *LinkedNode
}

type LinkedNode struct{
	key int
	value int
	prev *LinkedNode
	next *LinkedNode
}


func Constructor(capacity int) LRUCache {
	m:=make(map[int]*LinkedNode)
	head:=&LinkedNode{0,0,nil,nil}
	tail:=&LinkedNode{0,0,nil,nil}
	head.next=tail
	tail.prev=head
	return LRUCache{m,capacity,head,tail}
}

func (this *LRUCache)AddNode(node *LinkedNode){
	head:=this.head
	node.next=head.next
	node.next.prev=node
	node.prev=head
	head.next=node
}

func (this *LRUCache)RemoveNode(node *LinkedNode){
	node.prev.next=node.next
	node.next.prev=node.prev
}

func (this *LRUCache)MoveToHead(node *LinkedNode){
	this.RemoveNode(node)
	this.AddNode(node)
}

func (this *LRUCache) Get(key int) int {
	cache:=this.m
	if v,ok:=cache[key];ok{
		this.MoveToHead(v)
		return v.value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	cache:=this.m
	if v,ok:=cache[key];ok{
		v.value=value
		this.MoveToHead(v)
	}else{
		node:=&LinkedNode{key,value,nil,nil}
		if len(cache)==this.cap{
			tail:=this.tail
			delete(cache,tail.prev.key)
			this.RemoveNode(tail.prev)
		}
		cache[key]=node
		this.AddNode(node)
	}
}