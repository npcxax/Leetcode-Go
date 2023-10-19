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
func merge(nums []int, tempArr []int, left int, right int, rightEnd int) {
	var leftEnd = right
	var c = left
	var minNum = 0
	for left <= leftEnd || right <= rightEnd {
		if left > leftEnd {
			minNum = nums[right]
			right++
		} else if right > rightEnd {
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
		merge(nums, tempArr, i, i+length, n-1)
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

// 6
func quickSort(nums []int) {

}

// 7
func heapSort(nums []int) {

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
