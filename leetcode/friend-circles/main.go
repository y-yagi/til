package main

func findCircleNum(M [][]int) int {
	friends := make([]int, len(M))
	for i := range friends {
		friends[i] = i
	}
	for i := 0; i < len(M); i++ {
		for j := i + 1; j < len(M); j++ {
			if M[i][j] == 1 {
				union(i, j, friends)
			}
		}
	}
	m := map[int]interface{}{}
	for _, v := range friends {
		m[find(v, friends)] = nil
	}
	return len(m)
}

func union(i, j int, friends []int) {
	fi, fj := find(i, friends), find(j, friends)
	if fi > fj {
		friends[fi] = fj
	} else if fi < fj {
		friends[fj] = fi
	}
}

func find(i int, friends []int) int {
	for friends[i] != i {
		friends[i], i = friends[friends[i]], friends[i]
	}
	return i
}
