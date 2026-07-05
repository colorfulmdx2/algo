package linked_list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var hash = make(map[*Node]*Node)

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	if node, ok := hash[head]; ok {
		return node
	}

	copyNode := &Node{Val: head.Val}
	hash[copyNode] = copyNode
	copyNode.Next = copyRandomList(head.Next)
	copyNode.Random = hash[head.Random]

	return copyNode
}
