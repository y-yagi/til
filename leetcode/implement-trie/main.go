package main

type Trie struct {
	isWord      bool
	childrenMap map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{childrenMap: map[byte]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if _, found := cur.childrenMap[c]; !found {
			cur.childrenMap[c] = &Trie{childrenMap: map[byte]*Trie{}}
		}
		cur = cur.childrenMap[c]
	}
	cur.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if _, found := cur.childrenMap[c]; !found {
			return false
		}
		cur = cur.childrenMap[c]
	}
	return cur.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		c := prefix[i] - 'a'
		if _, found := cur.childrenMap[c]; !found {
			return false
		}
		cur = cur.childrenMap[c]
	}
	return true
}
