/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
type LRUCache struct {
	kv             map[int]*Node
	head, tail     *Node
	size, capacity int
}

type Node struct {
	key, value int
	next, prev *Node
}

func (this *LRUCache) moveToHead(node *Node) {
	this.remove(node)
	this.addToHead(node)
}

func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) addToHead(node *Node) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// 🌟特别要注意这里，removeTail之后，this.tail.prev就变了，所以这里返回tailNode
func (this *LRUCache) removeTail() *Node {
	node := this.tail.prev
	this.remove(node)
	return node
}

func Constructor(capacity int) LRUCache {
	head := &Node{
		key:   -1,
		value: -1,
	}
	tail := &Node{
		key:   -1,
		value: -1,
		prev:  head,
	}
	head.next = tail

	return LRUCache{
		kv: make(map[int]*Node, capacity),
		//sentinel
		head: head,
		//sentinel
		tail:     tail,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.kv[key]
	if !exist {
		return -1
	}
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, exist := this.kv[key]
	if exist {
		node.value = value
		this.moveToHead(node)
	} else {
		if this.size == this.capacity {
			tail := this.removeTail()
			this.size--
			delete(this.kv, tail.key)
		}

		newNode := &Node{
			value: value,
			key:   key,
		}
		this.kv[key] = newNode
		this.addToHead(newNode)
		this.size++
	}

}

// 🌟哨兵节点的应用🌟

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

