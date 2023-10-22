package leetcode

import "sort"

// timeout
// func threeSum(nums []int) [][]int {
// 	var n = len(nums)
// 	quickSort(nums)
// 	var result = make([][]int, 0)
// 	for i := 0; i < n-2; i++ {
// 		if i == 0 || nums[i] != nums[i-1] {
// 			for j := i + 1; j < n-1; j++ {
// 				if j == i+1 || nums[j] != nums[j-1] {
// 					for k := j + 1; k < n; k++ {
// 						if k == j+1 || nums[k] != nums[k-1] {
// 							if nums[i]+nums[j]+nums[k] == 0 {
// 								result = append(result, []int{nums[i], nums[j], nums[k]})
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return result
// }

// my solution
func threeSum(nums []int) [][]int {
	var n = len(nums)
	sort.Ints(nums)
	var result = make([][]int, 0)

	for i := 0; i < n-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			var left, right = i + 1, n - 1
			for left < right {
				if left != i+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				if right != n-1 && nums[right] == nums[right+1] {
					right--
					continue
				}

				if nums[left]+nums[right] == -nums[i] {
					result = append(result, []int{nums[i], nums[left], nums[right]})
					left++
					right--
				} else if nums[left]+nums[right] > -nums[i] {
					right--
				} else {
					left++
				}
			}
		}
	}

	return result
}

// official
func threeSum(nums []int) [][]int {
	var n = len(nums)
	sort.Ints(nums)
	var result = make([][]int, 0)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		var k = n - 1
		for j := i + 1; j < n-1; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[i]+nums[j]+nums[k] > 0 {
				k--
			}
			if j == k {
				break
			}
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
			}
		}
	}

	return result
}
