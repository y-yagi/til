package main

import (
	"strings"
)

func findDuplicate(paths []string) [][]string {
	dict := map[string][]string{}
	for i := 0; i < len(paths); i++ {
		pathAndFiles := strings.Split(paths[i], " ")
		path := pathAndFiles[0]
		for i := 1; i < len(pathAndFiles); i++ {
			fileWithContent := strings.Split(pathAndFiles[i], "(")
			fileWithoutContent := strings.TrimRight(fileWithContent[0], "(")
			dict[fileWithContent[1]] = append(dict[fileWithContent[1]], path+"/"+fileWithoutContent)
		}
	}

	ans := [][]string{}
	for _, v := range dict {
		if len(v) > 1 {
			ans = append(ans, v)
		}
	}
	return ans
}
