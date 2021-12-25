package test

type BTree struct {
	Val         int
	Left, Right *BTree
}

func TestMakeBTree(arr []int) *BTree {
	root := &BTree{
		Val: arr[0],
	}

	if len(arr) == 1 {
		return root
	}

	var q queue
	q.push(root)

	i := 1

	for {
		node := q.pop()
		node.Left = &BTree{
			Val: arr[i],
		}
		i++
		if i > len(arr)-1 {
			break
		}
		q.push(node.Left)

		node.Right = &BTree{
			Val: arr[i],
		}
		i++
		if i > len(arr)-1 {
			break
		}
		q.push(node.Right)
	}

	return root
}

type queue []*BTree

func (q *queue) push(node *BTree) {
	*q = append(*q, node)
}

func (q *queue) pop() *BTree {
	top := (*q)[0]
	*q = (*q)[1:]
	return top
}

func (q *queue) empty() bool {
	return len(*q) == 0
}
