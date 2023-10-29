package leetcode

import (
	"strings"
)

// my solution
func simplifyPath(path string) string {
	vs := strings.Split(path, "/")
	var stk = make([]string, 0)
	var result []byte
	result = append(result, '/')
	for _, item := range vs {
		if item == "" || item == "." {
			continue
		}
		if item == ".." {
			if len(stk) > 0 {
				stk = stk[:len(stk)-1]
			}
		} else {
			stk = append(stk, item)
		}
	}

	for _, item := range stk {
		result = append(result, []byte(item)...)
		result = append(result, '/')
	}
	if len(stk) > 0 {
		result = result[:len(result)-1]
	}

	return string(result)
}

// official
func simplifyPath(path string) string {
	var stk = make([]string, 0)
	for _, item := range strings.Split(path, "/") {
		if item == ".." {
			if len(stk) > 0 {
				stk = stk[:len(stk)-1]
			}
		} else if item != "." && item != "" {
			stk = append(stk, item)
		}
	}
	return "/" + strings.Join(stk, "/")
}
