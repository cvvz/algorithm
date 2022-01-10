/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
type LRUCache struct {
	cache        map[int]*Node
	size         int
	sentinelHead *Node
	sentinelTail *Node
}

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := new(Node), new(Node)
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		size:         capacity,
		cache:        make(map[int]*Node),
		sentinelHead: head,
		sentinelTail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, exist := this.cache[key]; exist {
		// remove
		remove(node)
		// insertHead
		insertHead(node, this.sentinelHead)

		return node.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exist := this.cache[key]; exist {
		node.Val = value
		remove(node)
		// insertHead
		insertHead(node, this.sentinelHead)
		return
	}

	node := &Node{
		Key: key,
		Val: value,
	}

	this.cache[key] = node
	// insertHead
	insertHead(node, this.sentinelHead)

	if len(this.cache) > this.size {
		// removeTail
		tail := this.sentinelTail.Prev
		remove(tail)
		delete(this.cache, tail.Key)
	}
}

func remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func insertHead(node *Node, head *Node) {
	head.Next.Prev = node
	node.Next = head.Next
	node.Prev = head
	head.Next = node
}

// 用map维护key
// value是链表中的一个节点，当get和put时，节点放到链表的头部
// 删除时删除链表的尾
// 双向链表

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

