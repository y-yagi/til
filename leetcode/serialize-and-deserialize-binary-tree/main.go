package main

import (
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

const Null = "n"

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	this.values = []string{}
	this.rserialize(root)
	return strings.Join(this.values, ",")
}

func (this *Codec) rserialize(node *TreeNode) {
	if node == nil {
		this.values = append(this.values, Null)
	} else {
		this.values = append(this.values, strconv.Itoa(node.Val))
		this.rserialize(node.Left)
		this.rserialize(node.Right)
	}
}

func (this *Codec) deserialize(data string) *TreeNode {
	this.values = strings.Split(data, ",")
	return this.rdeserialize()
}

func (this *Codec) rdeserialize() *TreeNode {
	value := this.values[0]
	this.values = this.values[1:]
	if value == Null {
		return nil
	}

	i, _ := strconv.Atoi(value)
	root := &TreeNode{Val: i}
	root.Left = this.rdeserialize()
	root.Right = this.rdeserialize()
	return root
}
