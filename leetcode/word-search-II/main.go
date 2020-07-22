package main

type Trie struct {
	children map[byte]*Trie
	word     string
}

func findWords(board [][]byte, words []string) []string {
	trie := &Trie{children: map[byte]*Trie{}}
	for _, word := range words {
		trie.Insert(word)
	}

	ans := []string{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, board, trie, &ans)
		}
	}

	return ans
}

func dfs(i, j int, board [][]byte, node *Trie, ans *[]string) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}

	c := board[i][j]
	if c == '#' || node.children[c-'a'] == nil {
		return
	}

	node = node.children[c-'a']
	if len(node.word) != 0 {
		*ans = append(*ans, node.word)
		node.word = ""
	}

	board[i][j] = '#'
	dfs(i+1, j, board, node, ans)
	dfs(i-1, j, board, node, ans)
	dfs(i, j+1, board, node, ans)
	dfs(i, j-1, board, node, ans)
	board[i][j] = c

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
