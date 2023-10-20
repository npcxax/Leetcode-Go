package leetcoode

func longestCommonPrefix(strs []string) string {
	i := 0
	for {
		var ch byte
		for idx, str := range strs {
			if i >= len(str) {
				return strs[0][:i]
			}
			if idx == 0 {
				ch = str[i]
			} else {
				if ch != str[i] {
					return strs[0][:i]
				}
			}
		}
		i++
	}
}
