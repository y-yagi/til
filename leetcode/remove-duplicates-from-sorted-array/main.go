package main

func removeDuplicates(s []int) int {
	for i := 0; i < len(s); i++ {
		index := indexOf(s[i+1:], s[i])
		if index != -1 {
			index = index + i + 1
			s = append(s[:index], s[index+1:]...)
			i--
		}
	}

	return len(s)
}

func indexOf(s []int, e int) int {
	for i, v := range s {
		if e == v {
			return i
		}
	}
	return -1
}
