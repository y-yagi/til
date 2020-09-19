package main

import "container/heap"

type MedianFinder struct {
	less   *MaximumHeap
	more   *MinimumHeap
	middle int
	length int
}

func Constructor() MedianFinder {
	less, more := &MaximumHeap{}, &MinimumHeap{}
	heap.Init(less)
	heap.Init(more)

	return MedianFinder{
		less: less,
		more: more,
	}
}

// AddNum satisfy the following rules:
// 1) 0 ≤ Len(mf.less)-Len(mf.more) ≤ 1
// 2) when mf.length is odd, mf.middle always be the middle one
// 3) when mf.length is even, mf.middle always be the bigger of two middle values, another middle value is (*mf.less)[0]
func (mf *MedianFinder) AddNum(num int) {
	switch mf.length {
	case 0:
		mf.middle = num
	case 1:
		less := num
		if num > mf.middle {
			less = mf.middle
			mf.middle = num
		}
		heap.Push(mf.less, less)
	case 2:
		if num >= mf.middle {
			heap.Push(mf.more, num)
		} else if num >= (*mf.less)[0] {
			heap.Push(mf.more, mf.middle)
			mf.middle = num
		} else {
			lessValue := heap.Pop(mf.less).(int)
			heap.Push(mf.more, mf.middle)
			mf.middle = lessValue
			heap.Push(mf.less, num)
		}
	default:
		magic, little, big := 0, 0, 0
		if num < mf.middle {
			maxOfLess := (*mf.less)[0]
			if maxOfLess < num {
				little = num
				big = mf.middle
			} else {
				_ = heap.Pop(mf.less)
				magic = num
				little = maxOfLess
				big = mf.middle
				heap.Push(mf.less, magic)
			}
		} else {
			minOfMore := (*mf.more)[0]
			if minOfMore < num {
				_ = heap.Pop(mf.more)
				little = mf.middle
				big = minOfMore
				magic = num
				heap.Push(mf.more, magic)
			} else {
				little = mf.middle
				big = num
			}

		}

		if mf.length%2 == 0 {
			mf.middle = little
			heap.Push(mf.more, big)
		} else {
			mf.middle = big
			heap.Push(mf.less, little)
		}
	}

	mf.length++
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.length == 0 {
		return float64(0)
	}

	if mf.length%2 != 0 {
		return float64(mf.middle)
	} else {
		return float64(mf.middle+(*mf.less)[0]) / 2
	}
}

// code for MinimumHeap & MaximumHeap, ignorable

type MinimumHeap []int
type MaximumHeap []int

func (mh MinimumHeap) Less(i, j int) bool {
	return mh[i] < mh[j]
}

func (mh MinimumHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh MinimumHeap) Len() int {
	return len(mh)
}

func (mh *MinimumHeap) Push(v interface{}) {
	*mh = append(*mh, v.(int))
}

func (mh *MinimumHeap) Pop() interface{} {
	length := len(*mh)
	v := (*mh)[length-1]
	*mh = (*mh)[:length-1]
	return v
}

func (mh MaximumHeap) Less(i, j int) bool {
	return mh[i] > mh[j]
}

func (mh MaximumHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh MaximumHeap) Len() int {
	return len(mh)
}

func (mh *MaximumHeap) Push(v interface{}) {
	*mh = append(*mh, v.(int))
}

func (mh *MaximumHeap) Pop() interface{} {
	length := len(*mh)
	v := (*mh)[length-1]
	*mh = (*mh)[:length-1]
	return v
}
