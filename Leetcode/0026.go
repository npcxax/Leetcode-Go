package leetcoode

// if nums[fast]!=nums[fast-1], then let nums[slow] = nums[fast]
func removeDuplicates(nums []int) int {
	var slow, fast = 1, 1

	for fast < len(nums) {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
