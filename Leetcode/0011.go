package leetcode

func maxArea(height []int) int {
	var maxarea = 0
	var curarea = 0
	var left, right = 0, len(height) - 1
	for left < right {
		curarea = min(height[left], height[right]) * (right - left)
		if curarea > maxarea {
			maxarea = curarea
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxarea
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
