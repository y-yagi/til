package main

type pair struct {
	key   string
	value int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	l := len(beginWord)
	dict := map[string][]string{}

	for _, word := range wordList {
		for i := 0; i < l; i++ {
			newWord := word[0:i] + "*" + word[i+1:l]
			transformations, found := dict[newWord]
			if !found {
				transformations = make([]string, 0)
			}
			transformations = append(transformations, word)
			dict[newWord] = transformations
		}
	}

	q := []pair{}
	q = append(q, pair{key: beginWord, value: 1})

	visited := map[string]bool{}
	visited[beginWord] = true

	for len(q) != 0 {
		node := q[0]
		q = q[1:]

		word := node.key
		level := node.value

		for i := 0; i < l; i++ {
			newWord := word[0:i] + "*" + word[i+1:l]
			transformations, _ := dict[newWord]
			for _, adjacentWord := range transformations {
				if adjacentWord == endWord {
					return level + 1
				}

				if _, found := visited[adjacentWord]; !found {
					visited[adjacentWord] = true
					q = append(q, pair{key: adjacentWord, value: level + 1})
				}
			}
		}
	}

	return 0
}
