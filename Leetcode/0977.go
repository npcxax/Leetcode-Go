package leetcodego

func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	ret := make([]int, right+1)
	idx := right
	for idx >= 0 {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			ret[idx] = nums[left] * nums[left]
			left++
		} else {
			ret[idx] = nums[right] * nums[right]
			right--
		}
		idx--
	}
	return ret
}
