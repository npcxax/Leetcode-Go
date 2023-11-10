package leetcode

// my solution
func combine(n int, k int) [][]int {
	var (
		result    = make([][]int, 0)
		backtrack func(cur int, arr []int)
		seen      = make([]bool, n+1)
	)
	backtrack = func(cur int, arr []int) {
		arr = append(arr, cur)
		seen[cur] = true
		if len(arr) == k {
			comb := make([]int, k)
			copy(comb, arr)
			result = append(result, comb)
			seen[cur] = false
			return
		}
		for i := cur + 1; i <= n; i++ {
			if !seen[i] {
				backtrack(i, arr)
			}
		}
		seen[cur] = false
	}

	for i := 1; i <= n; i++ {
		backtrack(i, []int{})
	}
	return result
}

// official
func combine(n int, k int) [][]int {
	var (
		result    = make([][]int, 0)
		backtrack func(cur int, arr []int)
	)
	backtrack = func(cur int, arr []int) {
		if len(arr)+(n-cur+1) < k {
			return
		}
		if len(arr) == k {
			comb := make([]int, k)
			copy(comb, arr)
			result = append(result, comb)
			return
		}
		arr = append(arr, cur)
		backtrack(cur+1, arr)
		arr = arr[:len(arr)-1]
		backtrack(cur+1, arr)
	}

	backtrack(1, []int{})
	return result
}
