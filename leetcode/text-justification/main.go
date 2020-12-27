package main

import (
	"bytes"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	lines := []string{}
	current := []string{}
	currentLen := 0

	for i, w := range words {
		if currentLen+len(current)+len(w) > maxWidth {
			if len(current) == 1 {
				curLine := current[0] + strings.Repeat(" ", maxWidth-len(current[0]))
				lines = append(lines, curLine)
			} else {
				diff := maxWidth - currentLen
				spaces := diff / (len(current) - 1)
				more := diff % (len(current) - 1)
				curLine := bytes.Buffer{}
				for ci, cw := range current {
					curLine.WriteString(cw)
					if ci != len(current)-1 {
						moreBlaks := 0
						if more > 0 {
							moreBlaks = 1
							more--
						}
						curLine.WriteString(strings.Repeat(" ", spaces+moreBlaks))
					}
				}
				lines = append(lines, curLine.String())
			}
			current = []string{}
			currentLen = 0
		}
		currentLen += len(w)
		current = append(current, w)

		if i == len(words)-1 {
			lastLine := strings.Join(current, " ")
			lastLine = lastLine + strings.Repeat(" ", maxWidth-len(lastLine))
			lines = append(lines, lastLine)
		}

	}

	return lines
}
