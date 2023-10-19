package leetcoode

// dp
func trap(height []int) int {
	var n = len(height)
	var left, right = make([]int, n), make([]int, n)
	var total = 0

	left[0] = height[0]
	right[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i-1])
		right[n-i-1] = max(right[n-i], height[n-i])
	}

	for i := 1; i < n-1; i++ {
		if min(left[i], right[i])-height[i] > 0 {
			total += (min(left[i], right[i]) - height[i])
		}
	}

	return total
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

// Monotone Stack
func trap(height []int) int {
	var stack = make([]int, 0)
	var total = 0

	for i, h := range height {
		for len(stack) > 0 && h > stack[len(stack)-1] {
			var top = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			var left = stack[len(stack)-1]
			total += (i - left - 1) * (min(height[left], height[i]) - height[top])
		}
		stack = append(stack, i)
	}

	return total
}
