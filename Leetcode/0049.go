package leetcode

import "sort"

// my solution
type myString []byte

func (ms myString) Len() int {
	return len(ms)
}

func (ms myString) Less(i, j int) bool {
	return ms[i] < ms[j]
}

func (ms myString) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func groupAnagrams(strs []string) [][]string {
	var hashmap = make(map[string]int)
	var result = make([][]string, 0)

	for _, str := range strs {
		b := []byte(str)
		sort.Sort(myString(b))
		sortStr := string(b)
		if v, ok := hashmap[sortStr]; ok {
			result[v] = append(result[v], str)
		} else {
			hashmap[sortStr] = len(result)
			result = append(result, []string{str})
		}
	}

	return result
}

// official
func groupAnagrams(strs []string) [][]string {
	var hashmap = make(map[string]int)
	var result = make([][]string, 0)
	for _, str := range strs {
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		var tempStr = string(b)
		if v, ok := hashmap[tempStr]; ok {
			result[v] = append(result[v], str)
		} else {
			idx := len(result)
			result = append(result, []string{str})
			hashmap[tempStr] = idx
		}
	}
	return result
}

func groupAnagrams(strs []string) [][]string {
	var hashmap = make(map[string]int)
	var result = make([][]string, 0)

	for _, str := range strs {
		var b = make([]byte, 26)
		for _, ch := range str {
			b[ch-'a']++
		}
		tempStr := string(b)
		if v, ok := hashmap[tempStr]; ok {
			result[v] = append(result[v], str)
		} else {
			hashmap[tempStr] = len(result)
			result = append(result, []string{str})
		}
	}

	return result
}
