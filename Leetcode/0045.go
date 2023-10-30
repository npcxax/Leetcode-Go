package leetcode

// Reverse Lookup Departure Location
// func jump(nums []int) int {
// 	var n = len(nums)
// 	var right = n - 1

// 	for count := 0; ; count++ {
// 		if right == 0 {
// 			return count
// 		}
// 		for i := 0; i < right; i++ {
// 			if i+nums[i] >= right {
// 				right = i
// 				break
// 			}
// 		}
// 	}
// }

// Maximum position reachable by forward lookup
func jump(nums []int) int {
	var n = len(nums)
	var maxPosition = 0
	var end = 0
	var steps = 0

	for i := 0; i < n-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}

	return steps
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
