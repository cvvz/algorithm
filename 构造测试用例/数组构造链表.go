package main

import "fmt"

func main() {
	head := construct([]int{1, 2, 3, 4, 5, 6})
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
}

type LinkList struct {
	Val  int
	Next *LinkList
}

func construct(arr []int) *LinkList {
	dummyHead := &LinkList{}
	tail := dummyHead

	n := len(arr)
	for i := 0; i < n; i++ {
		tail.Next = &LinkList{
			Val: arr[i],
		}
		tail = tail.Next
	}

	return dummyHead.Next
}
