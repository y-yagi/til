package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	answer := [][]int{[]int{1}}

	for i := 1; i < numRows; i++ {
		tmp := []int{}
		before := answer[i-1]
		for j := 0; j < i+1; j++ {
			v := 0
			if j == 0 || j >= len(before) {
				v = 1
			} else {
				v = before[j-1] + before[j]
			}
			tmp = append(tmp, v)
		}
		answer = append(answer, tmp)
	}
	return answer
}
