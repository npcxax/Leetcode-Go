package leetcoode

func productExceptSelf(nums []int) []int {
	var n = len(nums)
	var R = 1
	var answer = make([]int, n)
	answer[0] = 1
	R = 1

	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		answer[i] = answer[i] * R
		R *= nums[i]
	}

	return answer
}
