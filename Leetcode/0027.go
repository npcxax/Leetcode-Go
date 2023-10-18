package leetcode

// func removeElement(nums []int, val int) int {
// 	left, right := 0, len(nums)-1
// 	for left <= right {
// 		if nums[left] == val {
// 			for right > left && nums[right] == val {
// 				right--
// 			}
// 			if right <= left {
// 				break
// 			}
// 			nums[left], nums[right] = nums[right], nums[left]
// 		}
// 		left++
// 	}
// 	return left
// }

func removeElement(nums []int, val int) int {
	var left, right = 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
