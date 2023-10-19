package leetcoode

// 1 stable
func bubbleSort(nums []int) {
	var n = len(nums)
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
}

// 2 not stable
func selectionSort(nums []int) {
	var n = len(nums)
	var t = 0
	for i := 0; i < n-1; i++ {
		t = i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[t] {
				t = j
			}
		}
		nums[i], nums[t] = nums[t], nums[i]
	}
}

// 3 stable
func insertionSort(nums []int) {
	var n = len(nums)
	var j = 0
	var num = 0
	for i := 1; i < n; i++ {
		j = i - 1
		num = nums[i]
		for ; j >= 0 && num < nums[j]; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = num
	}
}

// 4 not stable
func shellSort(nums []int) {
	var n = len(nums)
	var gap = 1
	for gap < n/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < n; i++ {
			var j = i - gap
			var num = nums[i]
			for j >= 0 && num < nums[j] {
				nums[j+gap] = nums[j]
				j -= gap
			}
			nums[j+gap] = num
		}
		gap /= 3
	}
}

// 5
// iterate
func merge(nums []int, tempArr []int, left int, right int, rightEnd int) {
	var leftEnd = right
	var c = left
	var minNum = 0
	for left < leftEnd || right < rightEnd {
		if left >= leftEnd {
			minNum = nums[right]
			right++
		} else if right >= rightEnd {
			minNum = nums[left]
			left++
		} else if nums[left] < nums[right] {
			minNum = nums[left]
			left++
		} else {
			minNum = nums[right]
			right++
		}
		tempArr[c] = minNum
		c++
	}
}

func mergeA(nums []int, tempArr []int, length int) {
	var n = len(nums)
	var i = 0
	for ; i < n-2*length; i += 2 * length {
		merge(nums, tempArr, i, i+length, i+2*length)
	}
	if i+length < n {
		merge(nums, tempArr, i, i+length, n)
	} else {
		for j := i; j < n; j++ {
			tempArr[j] = nums[j]
		}
	}
}

func mergeSort(nums []int) {
	var n = len(nums)
	var length = 1
	var tempArr = make([]int, n)
	for length < n {
		mergeA(nums, tempArr, length)
		length *= 2
		mergeA(tempArr, nums, length)
		length *= 2
	}
}

// recrusion
func merge(left []int, right []int) []int {
	var result = make([]int, 0)
	var minNum = 0
	var leftEnd, rightEnd = len(left), len(right)
	var l, r = 0, 0
	for l < leftEnd || r < rightEnd {
		if l >= leftEnd {
			minNum = right[r]
			r++
		} else if r >= rightEnd {
			minNum = left[l]
			l++
		} else if left[l] < right[r] {
			minNum = left[l]
			l++
		} else {
			minNum = right[r]
			r++
		}
		result = append(result, minNum)
	}
	return result
}

func mergeB(nums []int) []int {
	var n = len(nums)
	if n < 2 {
		return nums
	}
	left := nums[:n/2]
	right := nums[n/2:]
	return merge(mergeB(left), mergeB(right))
}

// 6
func getPivot(nums []int, left int, right int) int {
	var mid = (left + right) >> 1
	if nums[left] > nums[mid] {
		nums[left], nums[mid] = nums[mid], nums[left]
	}
	if nums[left] > nums[right] {
		nums[left], nums[right] = nums[right], nums[left]
	}
	if nums[mid] > nums[right] {
		nums[mid], nums[right] = nums[right], nums[mid]
	}
	nums[right-1], nums[mid] = nums[mid], nums[right-1]
	return nums[right-1]
}

func qSort(nums []int, left int, right int) {
	if left < right {
		var pivot = getPivot(nums, left, right)
		var l, r = left, right - 2
		for l <= r {
			if nums[l] >= pivot {
				if nums[r] <= pivot {
					nums[l], nums[r] = nums[r], nums[l]
					l++
					r--
				} else {
					r--
				}
			} else {
				l++
			}
		}
		nums[l], nums[right-1] = nums[right-1], nums[l]
		qSort(nums, left, l-1)
		qSort(nums, l+1, right)
	}
}

func quickSort(nums []int) {
	qSort(nums, 0, len(nums)-1)
}

// 7
func adjustHeap(nums []int, i int, n int) {
	var num = nums[i]
	var parent = i
	for child := 2*parent + 1; 2*parent+1 < n; parent = child {
		child = 2*parent + 1
		if child+1 < n && nums[child+1] > nums[child] {
			child += 1
		}
		if nums[child] <= num {
			break
		}
		nums[parent] = nums[child]
	}
	nums[parent] = num
}

func heapSort(nums []int) {
	var n = len(nums)
	// build heap
	for i := n / 2; i >= 0; i-- {
		adjustHeap(nums, i, n)
	}

	for i := n - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		adjustHeap(nums, 0, i)
	}
}

// 8
func countSort(nums []int) {

}

// 9
func buckerSort(nums []int) {

}

// 10
func radixSort(nums []int) {

}
