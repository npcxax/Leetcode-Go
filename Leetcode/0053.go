package leetcode

// dp
func maxSubArray(nums []int) int {
	var (
		result = nums[0]
		cur    = 0
	)
	for _, num := range nums {
		cur += num
		if cur > result {
			result = cur
		}
		if cur < 0 {
			cur = 0
		}

	}
	return result
}

// Divide and Conquer
// lSum,rSum,iSum,mSum
type Status struct {
	lSum, rSum, iSum, mSum int
}

func pushup(l, r Status) Status {
	s := Status{}
	s.iSum = l.iSum + r.iSum
	s.lSum = max(l.lSum, l.iSum+r.lSum)
	s.rSum = max(r.rSum, r.iSum+l.rSum)
	s.mSum = max(max(l.mSum, r.mSum), l.rSum+r.lSum)
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

func maxSubArray(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
