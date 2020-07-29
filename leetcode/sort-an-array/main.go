package main

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	pivot := len(nums) / 2
	leftList := sortArray(nums[0:pivot])
	rightList := sortArray(nums[pivot:len(nums)])

	return merge(leftList, rightList)
}

func merge(left []int, right []int) []int {
	ret := make([]int, len(left)+len(right))
	leftCursor, rightCursor, retCursor := 0, 0, 0

	for leftCursor < len(left) && rightCursor < len(right) {
		if left[leftCursor] < right[rightCursor] {
			ret[retCursor] = left[leftCursor]
			leftCursor++
		} else {
			ret[retCursor] = right[rightCursor]
			rightCursor++
		}
		retCursor++
	}

	for leftCursor < len(left) {
		ret[retCursor] = left[leftCursor]
		leftCursor++
		retCursor++
	}
	for rightCursor < len(right) {
		ret[retCursor] = right[rightCursor]
		rightCursor++
		retCursor++
	}

	return ret
}
