package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range frequency {
		buckets[freq] = append(buckets[freq], num)
	}

	fmt.Println(buckets)
	var result []int
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		result = append(result, buckets[i]...)
	}

	if len(result) > k {
		result = result[:k]
	}

	return result
}
