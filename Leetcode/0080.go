package leetcode

func removeDuplicates(nums []int) int {
	var n = len(nums)
	if n <= 2 {
		return n
	}
	var slow, fast = 2, 2

	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
