package main

import "fmt"

type ValidWordAbbr struct {
	abbrDict map[string]bool
	dict     map[string]bool
}

func Constructor(dictionary []string) ValidWordAbbr {
	this := ValidWordAbbr{}
	abbrDict := map[string]bool{}
	dict := map[string]bool{}

	for _, word := range dictionary {
		dict[word] = true
	}

	for word, _ := range dict {
		abbr := this.omit(word)
		_, found := abbrDict[abbr]
		abbrDict[abbr] = !found
	}

	this.dict = dict
	this.abbrDict = abbrDict
	return this
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	v, found := this.abbrDict[this.omit(word)]
	if !found {
		return true
	}

	_, found = this.dict[word]
	if v && found {
		return true
	}

	return false
}

func (this *ValidWordAbbr) omit(word string) string {
	l := len(word)
	if l < 3 {
		return word
	}

	return fmt.Sprintf("%b%d%b", word[0], l-2, word[l-1])
}
