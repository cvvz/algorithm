// https://leetcode-cn.com/problems/4ueAj6/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

func insert(aNode *Node, x int) *Node {
	node := &Node{
		Val: x,
	}

	if aNode == nil {
		node.Next = node
		return node
	}

	p := aNode
	for processNext(p, x) && p.Next != aNode {
		p = p.Next
	}

	//insert
	node.Next = p.Next
	p.Next = node

	return aNode
}

func processNext(p *Node, x int) bool {
	if p.Val > x {
		if isTail(p) && p.Next.Val > x {
			return false
		}
		return true
	}

	if p.Next.Val >= x {
		return false
	} else if isTail(p) {
		return false
	}

	return true

}

func isTail(p *Node) bool {
	return p.Val > p.Next.Val
}
