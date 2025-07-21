package main

func subsets(nums []int) [][]int {
	ans := [][]int{{}}
	for _, num := range nums {
		n := len(ans)
		for i := 0; i < n; i++ {
			newSubset := append([]int{}, ans[i]...)
			newSubset = append(newSubset, num)
			ans = append(ans, newSubset)
		}
	}

	return ans
}
