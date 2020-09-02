package main

import (
	"fmt"
	"sort"
)

func largestTimeFromDigits(A []int) string {
	sort.Ints(A)
	for i := 3; i >= 0; i-- {
		for j := 3; j >= 0; j-- {
			if i == j {
				continue
			}
			for k := 3; k >= 0; k-- {
				if k == i || k == j {
					continue
				}
				h, m := 10*A[i]+A[j], 10*A[k]+A[6-i-j-k]
				if h < 24 && m < 60 {
					return fmt.Sprintf("%02d:%02d", h, m)
				}
			}
		}
	}
	return ""
}
