package main

type WordDictionary struct {
	trie *TrieNode
}

type TrieNode struct {
	m    map[byte]*TrieNode
	word bool
}

func Constructor() WordDictionary {
	trie := &TrieNode{m: map[byte]*TrieNode{}}
	return WordDictionary{trie: trie}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.trie
	for i := 0; i < len(word); i++ {
		w := word[i]
		if v, ok := node.m[w]; ok {
			node = v
		} else {
			node.m[w] = &TrieNode{m: make(map[byte]*TrieNode)}
			node = node.m[w]
		}
	}
	node.word = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.searchInNode(word, this.trie)
}

func (this *WordDictionary) searchInNode(word string, node *TrieNode) bool {
	if len(word) == 0 {
		return false
	}

	if word[0] == '.' {
		for _, v := range node.m {
			if v.m['0'] != nil && len(word) == 1 {
				return true
			}
			if this.searchInNode(word[1:], v) {
				return true
			}
		}
	} else {
		if v, ok := node.m[word[0]]; ok {
			if v.m['0'] != nil && len(word) == 1 {
				return true
			} else {
				return this.searchInNode(word[1:], v)
			}
		}
	}
	return node.word
}
