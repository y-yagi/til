package main

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	dict := map[string]bool{}

	for i := 0; i < len(emails); i++ {
		localAndDomain := strings.Split(emails[i], "@")
		if strings.Contains(localAndDomain[0], ".") {
			localAndDomain[0] = strings.ReplaceAll(localAndDomain[0], ".", "")
		}
		if index := strings.Index(localAndDomain[0], "+"); index != -1 {
			localAndDomain[0] = localAndDomain[0][0:index]
		}

		dict[localAndDomain[0]+"@"+localAndDomain[1]] = true
	}
	return len(dict)
}
