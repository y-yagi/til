package main

import "strings"

func reformatDate(date string) string {
	months := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}

	dates := strings.Split(date, " ")
	ans := dates[2] + "-"

	month, _ := months[dates[1]]
	ans += month + "-"

	days := []byte{}
	for i := 0; i < len(dates[0]); i++ {
		if dates[0][i] > '9' || dates[0][i] < '0' {
			break
		}
		days = append(days, dates[0][i])
	}

	if len(days) == 1 {
		t := days[0]
		days[0] = '0'
		days = append(days, t)
	}

	ans += string(days)
	return ans
}
