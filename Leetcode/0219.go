package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	var hashmap = make(map[int]struct{})
	for i, num := range nums {
		if i > k {
			delete(hashmap, nums[i-k-1])
		}
		if _, ok := hashmap[num]; ok {
			return true
		}
		hashmap[nums[i]] = struct{}{}
	}
	return false
}
