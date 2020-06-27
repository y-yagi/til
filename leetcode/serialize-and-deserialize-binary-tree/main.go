package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	values []string
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	str := ""
	this.rserialize(root, &str)
	return str
}

func (this *Codec) rserialize(node *TreeNode, str *string) {
	if node == nil {
		*str += "null,"
	} else {
		*str += fmt.Sprintf("%d,", node.Val)
		this.rserialize(node.Left, str)
		this.rserialize(node.Right, str)
	}
}

func (this *Codec) deserialize(data string) *TreeNode {
	this.values = strings.Split(data, ",")
	return this.rdeserialize()
}

func (this *Codec) rdeserialize() *TreeNode {
	value := this.values[0]
	this.values = this.values[1:]
	if value == "null" {
		return nil
	}

	i, _ := strconv.Atoi(value)
	root := &TreeNode{Val: i}
	root.Left = this.rdeserialize()
	root.Right = this.rdeserialize()
	return root
}
