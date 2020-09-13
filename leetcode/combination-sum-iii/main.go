package main

func combinationSum3(k int, n int) [][]int {
	combinations := [][]int{}
	r := []int{}
	m := map[int]bool{}
	candidates := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	find(k, n, candidates, &combinations, r, &m, 0)
	return combinations
}

func find(numbers, target int, candidates []int, combinations *[][]int, r []int, m *map[int]bool, i int) {
	if target < 0 || len(r) > numbers {
		return
	}

	if target == 0 {
		if len(r) != numbers {
			return
		}

		cr := make([]int, len(r))
		copy(cr, r)
		*combinations = append(*combinations, cr)
		return
	}

	for i < len(candidates) {
		if _, found := (*m)[candidates[i]]; found {
			i++
			continue
		}

		r = append(r, candidates[i])
		(*m)[candidates[i]] = true
		find(numbers, target-candidates[i], candidates, combinations, r, m, i)
		delete(*m, candidates[i])
		i++
		r = r[:len(r)-1]
	}

	return
}
