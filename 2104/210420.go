package leetcode

import "strings"

//https://leetcode-cn.com/problems/implement-strstr/
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	} else if haystack == "" {
		return -1
	}
	return strings.Index(haystack, needle)
}
