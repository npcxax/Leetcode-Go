package leetcode

import (
	"container/heap"
	"sort"
)

type MedianFinder struct {
	left  hp
	right hp
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() == 0 || -this.left.IntSlice[0] > num {
		heap.Push(&this.left, -num)
		if this.left.Len()-1 > this.right.Len() {
			heap.Push(&this.right, -heap.Pop(&this.left).(int))
		}
	} else {
		heap.Push(&this.right, num)
		if this.right.Len() > this.left.Len() {
			heap.Push(&this.left, -heap.Pop(&this.right).(int))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() != this.right.Len() {
		return -float64(this.left.IntSlice[0])
	}
	return float64(this.right.IntSlice[0]-this.left.IntSlice[0]) / 2
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
