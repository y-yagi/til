package main

type node struct {
	child [2]*node
}

func insert(root *node, n int) {
	cur := root
	for i := 0; i <= 30; i++ {
		s := uint(30 - i)
		c := (n >> s) & 1
		child := cur.child[c]
		if child == nil {
			child = &node{}
			cur.child[c] = child
		}
		cur = child
	}
}

func buildTrie(nums []int) *node {
	root := &node{}
	for _, n := range nums {
		insert(root, n)
	}
	return root
}

func findMaximumXOR(nums []int) int {
	root := buildTrie(nums)
	max := 0
	for _, n := range nums {
		val := 0
		cur := root
		for i := 0; i <= 30; i++ {
			s := uint(30 - i)
			c := (n >> s) & 1
			var child *node
			opposite := true
			if c == 0 {
				child = cur.child[1]
				if child == nil {
					child = cur.child[0]
					opposite = false
				}
			} else {
				child = cur.child[0]
				if child == nil {
					child = cur.child[1]
					opposite = false
				}
			}

			if child == nil {
				break
			}

			if opposite {
				val = val | (1 << s)
			}
			cur = child
		}
		if max < val {
			max = val
		}
	}
	return max
}
