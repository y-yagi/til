package main

func replaceElements(arr []int) []int {
	greatest := arr[len(arr)-1]

	for i := len(arr) - 1; i >= 0; i-- {
		before := greatest
		greatest = max(arr[i], greatest)
		arr[i] = before
	}
	arr[len(arr)-1] = -1
	return arr
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
