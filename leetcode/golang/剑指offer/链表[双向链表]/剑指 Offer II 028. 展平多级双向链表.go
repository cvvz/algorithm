// https://leetcode-cn.com/problems/Qv1Da2/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	current := root

	for current != nil {
		if current.Child == nil {
			current = current.Next
			continue
		}
		insert(current, current.Child)

		current = root
	}

	return root
}

func insert(l1 *Node, l2 *Node) {
	l1next := l1.Next
	l1.Next = l2
	l2.Prev = l1
	l1.Child = nil

	for l2.Next != nil {
		l2 = l2.Next
	}

	l2.Next = l1next
	if l1next != nil {
		l1next.Prev = l2
	}

}
