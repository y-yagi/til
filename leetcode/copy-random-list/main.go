package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cloned := map[*Node]*Node{}

	return copyList(head, cloned)

}

func copyList(head *Node, cloned map[*Node]*Node) *Node {
	if head == nil {
		return nil
	}

	if clone, ok := cloned[head]; ok {
		return clone
	}

	clone := &Node{
		Val: head.Val,
	}
	cloned[head] = clone

	clone.Next = copyList(head.Next, cloned)
	clone.Random = copyList(head.Random, cloned)

	return clone
}
