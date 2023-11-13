package leetcode

// my solution, sum[i:j] vs sum[0:i]+sum[j:n]
type Status struct {
	lSum, rSum, iSum, mSum int
}

func pushup(lS, rS Status) Status {
	s := Status{}
	s.iSum = lS.iSum + rS.iSum
	s.lSum = max(lS.lSum, lS.iSum+rS.lSum)
	s.rSum = max(rS.rSum, rS.iSum+lS.rSum)
	s.mSum = max(max(lS.mSum, rS.mSum), lS.rSum+rS.lSum)
	return s
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	mid := l + (r-l)>>1
	lS := get(nums, l, mid)
	rS := get(nums, mid+1, r)
	return pushup(lS, rS)
}

func maxSubarraySumCircular(nums []int) int {
	var (
		n       = len(nums)
		result  = nums[0]
		leftMax = make([]int, n)
		pre     = 0
	)

	result = get(nums, 0, n-1).mSum

	for i := 1; i < n; i++ {
		pre += nums[i-1]
		leftMax[i] = max(leftMax[i-1], pre)
	}
	pre = 0
	for j := n - 1; j > 0; j-- {
		pre += nums[j]
		result = max(result, leftMax[j]+pre)
	}

	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// official
func maxSubarraySumCircular(nums []int) int {
	var (
		n        = len(nums)
		result   = nums[0]
		leftMax  = make([]int, n)
		pre      = nums[0]
		leftSum  = nums[0]
		rightSum = 0
	)
	leftMax[0] = nums[0]
	for i := 1; i < n; i++ {
		pre = max(pre+nums[i], nums[i])
		result = max(result, pre)
		leftSum += nums[i]
		leftMax[i] = max(leftMax[i-1], leftSum)
	}

	for i := n - 1; i > 0; i-- {
		rightSum += nums[i]
		result = max(result, rightSum+leftMax[i-1])
	}

	return result
}

// reverse
func maxSubarraySumCircular(nums []int) int {
	var (
		n          = len(nums)
		maxRes     = nums[0]
		minRes     = nums[0]
		leftSum    = nums[0]
		leftMinSum = nums[0]
		total      = nums[0]
	)

	for i := 1; i < n; i++ {
		leftSum = max(leftSum+nums[i], nums[i])
		leftMinSum = min(leftMinSum+nums[i], nums[i])
		minRes = min(minRes, leftMinSum)
		maxRes = max(maxRes, leftSum)
		total += nums[i]
	}

	if maxRes < 0 {
		return maxRes
	}
	return max(maxRes, total-minRes)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// priority queue
func maxSubarraySumCircular(nums []int) int {
	type pair struct{ idx, val int }
	var (
		n   = len(nums)
		q   = make([]pair, 0)
		res = nums[0]
		pre = nums[0]
	)
	q = append(q, pair{0, nums[0]})
	for i := 1; i < 2*n; i++ {
		for len(q) > 0 && q[0].idx < i-n {
			q = q[1:]
		}
		pre += nums[i%n]
		res = max(res, pre-q[0].val)
		for len(q) > 0 && q[len(q)-1].val >= pre {
			q = q[:len(q)-1]
		}
		q = append(q, pair{i, pre})
	}
	return res
}
