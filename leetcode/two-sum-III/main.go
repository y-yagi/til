package main

type TwoSum struct {
	numbers map[int]int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum{numbers: map[int]int{}}
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	this.numbers[number]++
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	for k, v := range this.numbers {
		complement := value - k
		if complement != k {
			if _, found := this.numbers[complement]; found {
				return true
			}
		} else {
			if v > 1 {
				return true
			}
		}
	}
	return false
}
