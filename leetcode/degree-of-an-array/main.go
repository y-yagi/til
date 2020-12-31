package main

func findShortestSubArray(nums []int) int {
	type inner struct {
		freq  int
		left  int
		right int
	}
	m := map[int]*inner{}
	for i, v := range nums {
		if m[v] == nil {
			m[v] = &inner{left: i}
		}
		m[v].freq++
		m[v].right = i
	}
	maxFreq, smallestLength := 0, len(nums)
	for _, v := range m {
		if v.freq > maxFreq {
			maxFreq = v.freq
			smallestLength = v.right - v.left + 1
		} else if v.freq == maxFreq {
			if v.right-v.left+1 < smallestLength {
				smallestLength = v.right - v.left + 1
			}
		}
	}
	return smallestLength
}
