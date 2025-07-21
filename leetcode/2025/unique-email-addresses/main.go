package main

import "strings"

func numUniqueEmails(emails []string) int {
	unique := make(map[string]struct{})

	for _, email := range emails {
		parts := strings.Split(email, "@")
		local := strings.ReplaceAll(parts[0], ".", "")
		if idx := strings.Index(local, "+"); idx != -1 {
			local = local[:idx]
		}
		unique[local+"@"+parts[1]] = struct{}{}
	}

	return len(unique)
}
