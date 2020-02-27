package main

func findDisappearedNumbers(nums []int) []int {
	dict := make(map[int]bool, len(nums))
	missing := []int{}

	for _, n := range nums {
		dict[n] = true
	}

	for i := 1; i < len(nums)+1; i++ {
		if _, found := dict[i]; !found {
			missing = append(missing, i)
		}
	}

	return missing
}
