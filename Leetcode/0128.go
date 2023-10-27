package leetcode

func longestConsecutive(nums []int) int {
	var hashmap = make(map[int]bool)
	var result = 0

	for _, num := range nums {
		hashmap[num] = true
	}

	for num := range hashmap {
		if !hashmap[num-1] {
			var temp = 1
			for hashmap[num+1] {
				temp++
				num++
			}
			if temp > result {
				result = temp
			}
		}
	}

	return result
}
