package main

type Node struct {
	Val      int
	Children []*Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() *Codec {
	return new(Codec)
}

func (this *Codec) encode(root *Node) *TreeNode {
	if root == nil {
		return nil
	}

	newRoot := &TreeNode{Val: root.Val}
	if len(root.Children) > 0 {
		newRoot.Left = this.encode(root.Children[0])
	}

	sibling := newRoot.Left
	for i := 1; i < len(root.Children); i++ {
		sibling.Right = this.encode(root.Children[i])
		sibling = sibling.Right
	}

	return newRoot
}

func (this *Codec) decode(root *TreeNode) *Node {
	if root == nil {
		return nil
	}

	newRoot := &Node{Val: root.Val, Children: make([]*Node, 0)}
	sibling := root.Left

	for sibling != nil {
		newRoot.Children = append(newRoot.Children, this.decode(sibling))
		sibling = sibling.Right
	}
	return newRoot
}
