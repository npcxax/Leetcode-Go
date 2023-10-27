package leetcode

func twoSum(nums []int, target int) []int {
	var hashmap = make(map[int]int)
	for idx, num := range nums {
		if v, ok := hashmap[target-num]; ok {
			return []int{idx, v}
		}
		hashmap[num] = idx
	}
	return []int{0, 0}
}
