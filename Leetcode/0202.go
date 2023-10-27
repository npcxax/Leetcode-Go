package leetcode

func isHappy(n int) bool {
	var hashmap = make(map[int]bool)
	for n != 1 {
		if _, ok := hashmap[n]; ok {
			return false
		}
		hashmap[n] = true
		n = getNext(n)
	}
	return true
}

func getNext(n int) int {
	var total = 0
	for n > 0 {
		t := n % 10
		total += t * t
		n /= 10
	}
	return total
}
