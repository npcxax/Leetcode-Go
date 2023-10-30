package leetcode

func majorityElement(nums []int) int {
	var count = 1
	var majorityNum = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == majorityNum {
			count++
		} else {
			count--
			if count < 0 {
				majorityNum = nums[i]
				count = 1
			}
		}
	}
	return majorityNum
}
