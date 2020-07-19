package main

type TrieNode struct {
	children map[byte]*TrieNode
	sum      int
}

type MapSum struct {
	Map  map[string]int
	root *TrieNode
}

func Constructor() MapSum {
	return MapSum{
		Map:  make(map[string]int),
		root: &TrieNode{children: make(map[byte]*TrieNode)},
	}
}

func (this *MapSum) Insert(key string, val int) {
	diff := val - this.Map[key]
	this.Map[key] = val
	cur := this.root

	for i := range key {
		if _, found := cur.children[key[i]]; !found {
			trie := TrieNode{children: make(map[byte]*TrieNode)}
			cur.children[key[i]] = &trie
		}
		cur = cur.children[key[i]]
		cur.sum += diff
	}
}

func (this *MapSum) Sum(prefix string) int {
	cur := this.root
	for i := range prefix {
		if _, found := cur.children[prefix[i]]; !found {
			return 0
		}
		cur = cur.children[prefix[i]]
	}
	return cur.sum
}
