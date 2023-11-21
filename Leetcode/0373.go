package leetcode

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var m, n = len(nums1), len(nums2)
	h := &hp{
		arr:   nil,
		nums1: nums1,
		nums2: nums2,
	}

	for i := 0; i < m && i < k; i++ {
		heap.Push(h, pair{i, 0})
	}

	var result = make([][]int, 0)

	for h.Len() > 0 && k > 0 {
		p := heap.Pop(h).(pair)
		result = append(result, []int{nums1[p.x], nums2[p.y]})
		if p.y+1 < n {
			heap.Push(h, pair{p.x, p.y + 1})
		}
		k--
	}

	return result
}

type pair struct {
	x, y int
}

type hp struct {
	arr          []pair
	nums1, nums2 []int
}

func (h *hp) Len() int {
	return len(h.arr)
}

func (h *hp) Less(i, j int) bool {
	return h.nums1[h.arr[i].x]+h.nums2[h.arr[i].y] < h.nums1[h.arr[j].x]+h.nums2[h.arr[j].y]
}

func (h *hp) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *hp) Push(v interface{}) {
	h.arr = append(h.arr, v.(pair))
}

func (h *hp) Pop() interface{} {
	v := h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	return v
}
