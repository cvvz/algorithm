/*
 * @lc app=leetcode.cn id=706 lang=golang
 *
 * [706] 设计哈希映射
 */

// @lc code=start
type LinkList struct {
	Key  int
	Val  int
	Next *LinkList
}

type MyHashMap struct {
	hmap []*LinkList
}

func hash(num int) int {
	return num % 100
}

func Constructor() MyHashMap {
	hashMap := MyHashMap{
		hmap: make([]*LinkList, 100, 100),
	}

	for i := 0; i < 100; i++ {
		hashMap.hmap[i] = &LinkList{}
	}

	return hashMap
}

func (this *MyHashMap) Put(key int, value int) {
	dummyHead := this.hmap[key%100]
	head, tail := dummyHead.Next, dummyHead

	for head != nil {
		if head.Key == key {
			head.Val = value
			return
		}
		tail = head
		head = head.Next
	}

	tail.Next = &LinkList{
		Key: key,
		Val: value,
	}
}

func (this *MyHashMap) Get(key int) int {
	dummyHead := this.hmap[key%100]
	head := dummyHead.Next

	for head != nil {
		if head.Key == key {
			return head.Val
		}
		head = head.Next
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	dummyHead := this.hmap[key%100]
	head, prev := dummyHead.Next, dummyHead

	for head != nil {
		if head.Key == key {
			prev.Next = head.Next
			return
		}
		prev = head
		head = head.Next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end

