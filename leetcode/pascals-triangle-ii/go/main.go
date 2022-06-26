package main

func getRow(rowIndex int) []int {
	return generate(1, rowIndex+1, []int{1})
}

func generate(index, rowIndex int, before []int) []int {
	if index == rowIndex {
		return before
	}

	row := make([]int, index+1)
	for i := 0; i <= index; i++ {
		if i == 0 {
			row[i] = before[0]
		} else if i == index {
			row[i] = before[index-1]
		} else {
			row[i] = before[i-1] + before[i]
		}
	}

	index++
	return generate(index, rowIndex, row)
}
