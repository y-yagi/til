package main

import "strings"

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	base := []string{"", "Thousand", "Million", "Billion", "Trillion"}
	resultInSlice := []string{}
	baseOrder := 0
	for num != 0 {
		part := transOnce(num % 1000)
		if len(part) != 0 && baseOrder != 0 {
			part = append(part, base[baseOrder])
		}
		resultInSlice = append(part, resultInSlice...)
		num /= 1000
		baseOrder++
	}
	return strings.Join(resultInSlice, " ")
}

func transOnce(val int) []string {
	part := trans(val % 100)
	result := []string{}
	switch val / 100 {
	case 1:
		result = append(result, "One Hundred")
	case 2:
		result = append(result, "Two Hundred")
	case 3:
		result = append(result, "Three Hundred")
	case 4:
		result = append(result, "Four Hundred")
	case 5:
		result = append(result, "Five Hundred")
	case 6:
		result = append(result, "Six Hundred")
	case 7:
		result = append(result, "Seven Hundred")
	case 8:
		result = append(result, "Eight Hundred")
	case 9:
		result = append(result, "Nine Hundred")
	}
	return append(result, part...)
}

func trans(val int) []string {
	switch {
	case val == 0:
		return []string{}
	case val == 10:
		return []string{"Ten"}
	case val == 11:
		return []string{"Eleven"}
	case val == 12:
		return []string{"Twelve"}
	case val == 13:
		return []string{"Thirteen"}
	case val == 14:
		return []string{"Fourteen"}
	case val == 15:
		return []string{"Fifteen"}
	case val == 16:
		return []string{"Sixteen"}
	case val == 17:
		return []string{"Seventeen"}
	case val == 18:
		return []string{"Eighteen"}
	case val == 19:
		return []string{"Nineteen"}
	}
	transString := []string{}

	switch val / 10 {
	case 2:
		transString = append(transString, "Twenty")
	case 3:
		transString = append(transString, "Thirty")
	case 4:
		transString = append(transString, "Forty")
	case 5:
		transString = append(transString, "Fifty")
	case 6:
		transString = append(transString, "Sixty")
	case 7:
		transString = append(transString, "Seventy")
	case 8:
		transString = append(transString, "Eighty")
	case 9:
		transString = append(transString, "Ninety")
	}

	switch val % 10 {
	case 1:
		transString = append(transString, "One")
	case 2:
		transString = append(transString, "Two")
	case 3:
		transString = append(transString, "Three")
	case 4:
		transString = append(transString, "Four")
	case 5:
		transString = append(transString, "Five")
	case 6:
		transString = append(transString, "Six")
	case 7:
		transString = append(transString, "Seven")
	case 8:
		transString = append(transString, "Eight")
	case 9:
		transString = append(transString, "Nine")
	}

	return transString
}
