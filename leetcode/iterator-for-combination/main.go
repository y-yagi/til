package main

type CombinationIterator struct {
	characters        string
	combinationLength int
	combinations      []string
	index             int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	iterator := CombinationIterator{combinationLength: combinationLength, characters: characters}
	iterator.build(0, []byte{})
	return iterator
}

func (this *CombinationIterator) Next() string {
	if !this.HasNext() {
		return ""
	}

	i := this.index
	this.index++
	return this.combinations[i]
}

func (this *CombinationIterator) HasNext() bool {
	return len(this.combinations) > this.index
}

func (this *CombinationIterator) build(index int, current []byte) {
	if len(current) == this.combinationLength {
		cr := make([]byte, len(current))
		copy(cr, current)
		this.combinations = append(this.combinations, string(cr))
		return
	}

	for index < len(this.characters) {
		current = append(current, this.characters[index])
		index++
		this.build(index, current)
		current = current[:len(current)-1]
	}

	return
}
