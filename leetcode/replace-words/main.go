package main

import (
	"fmt"
	"strings"
)

type Trie struct {
	children map[byte]*Trie
	word     string
}

func replaceWords(dict []string, sentence string) string {
	trie := Trie{children: map[byte]*Trie{}}
	for _, v := range dict {
		trie.Insert(v)
	}

	ans := []string{}
	for _, word := range strings.Split(sentence, " ") {
		result := trie.Search(word)
		if len(result) == 0 {
			ans = append(ans, word)
		} else {
			ans = append(ans, result)
		}
	}

	return strings.Join(ans, " ")
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if _, found := cur.children[c]; !found {
			cur.children[c] = &Trie{children: map[byte]*Trie{}}
		}
		cur = cur.children[c]
	}
	cur.word = word
}

func (this *Trie) Search(word string) string {
	cur := this
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if _, found := cur.children[c]; !found || len(cur.word) != 0 {
			return cur.word
		}
		cur = cur.children[c]
	}

	return cur.word
}

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}
