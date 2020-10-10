package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	subdomains := map[string]int{}

	for _, cpdomain := range cpdomains {
		splitted := strings.Split(cpdomain, " ")
		count, _ := strconv.Atoi(splitted[0])
		subdomains[splitted[1]] += count

		for {
			i := strings.Index(splitted[1], ".")
			if i == -1 {
				break
			}
			l := len(splitted[1])
			splitted[1] = splitted[1][i+1 : l]
			subdomains[splitted[1]] += count
		}
	}

	ans := []string{}
	for k, v := range subdomains {
		ans = append(ans, fmt.Sprintf("%d %s", v, k))
	}
	return ans
}
