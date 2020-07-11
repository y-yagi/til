package main

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	dict := map[string]int{}
	minimum := math.MaxInt32
	ans := []string{}

	for i := 0; i < len(list1); i++ {
		dict[list1[i]] = i
	}

	for i := 0; i < len(list2); i++ {
		if v, found := dict[list2[i]]; found {
			if minimum > v+i {
				minimum = v + i
				ans = []string{list2[i]}
			} else if minimum == v+i {
				ans = append(ans, list2[i])
			}
		}
	}

	return ans
}
