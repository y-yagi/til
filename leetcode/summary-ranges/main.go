package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	ans := []string{}
	tmp := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if (len(tmp) == 0) || (tmp[len(tmp)-1]+1 == nums[i]) {
			tmp = append(tmp, nums[i])
			continue
		}

		if len(tmp) == 1 {
			ans = append(ans, fmt.Sprintf("%d", tmp[0]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", tmp[0], tmp[len(tmp)-1]))
		}
		tmp = []int{nums[i]}
	}

	if len(tmp) != 0 {
		if len(tmp) == 1 {
			ans = append(ans, fmt.Sprintf("%d", tmp[0]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", tmp[0], tmp[len(tmp)-1]))
		}
	}

	return ans
}

func d(a ...interface{}) {
	format := ""
	for _, _ = range a {
		format += "%v, "
	}
	format += "\n"
	fmt.Printf(format, a...)
}
