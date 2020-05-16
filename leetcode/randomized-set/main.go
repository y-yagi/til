package main

import (
	"math/rand"
)

type RandomizedSet struct {
	dict map[int]int
	list []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	dict := make(map[int]int, 0)
	list := make([]int, 0)
	return RandomizedSet{dict: dict, list: list}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, exist := this.dict[val]; exist {
		return false
	}

	this.list = append(this.list, val)
	this.dict[val] = len(this.list) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, exist := this.dict[val]; !exist {
		return false
	}

	last := this.list[len(this.list)-1]
	idx := this.dict[val]
	this.list[idx] = last
	this.dict[last] = idx

	this.list = this.list[0 : len(this.list)-1]
	delete(this.dict, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	n := rand.Intn(len(this.list))
	return this.list[n]
}
