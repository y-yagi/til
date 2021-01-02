package main

import "strconv"

func calPoints(ops []string) int {
	record := []int{}

	for _, op := range ops {
		if op == "C" {
			record = record[:len(record)-1]
			continue
		}

		if op == "+" {
			l := len(record) - 1
			record = append(record, record[l]+record[l-1])
			continue
		}

		if op == "D" {
			l := len(record) - 1
			record = append(record, record[l]*2)
			continue
		}

		v, _ := strconv.Atoi(op)
		record = append(record, v)
	}

	ans := 0
	for _, v := range record {
		ans += v
	}
	return ans
}
