package leetcode

// base on heap sort
func findKthLargest(nums []int, k int) int {
	var (
		m = len(nums)
		s func(parent int, n int)
	)

	s = func(parent, n int) {
		pre := nums[parent]
		for child := 2*parent + 1; 2*parent+1 < n; parent = child {
			child = 2*parent + 1
			if child+1 < n && nums[child] < nums[child+1] {
				child += 1
			}
			if pre >= nums[child] {
				break
			}
			nums[parent] = nums[child]
		}
		nums[parent] = pre
	}

	for i := m / 2; i >= 0; i-- {
		s(i, m)
	}

	for i := 1; i <= k; i++ {
		nums[0], nums[m-i] = nums[m-i], nums[0]
		s(0, m-i)
	}

	return nums[m-k]
}

// base on quick sort with selection
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickSortWithSelection(nums, 0, n-1, n-k)
}

func quickSortWithSelection(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	i, j := l-1, r+1
	partition := nums[l]
	for i < j {
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickSortWithSelection(nums, l, j, k)
	} else {
		return quickSortWithSelection(nums, j+1, r, k)
	}
}
