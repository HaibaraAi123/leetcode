package leetcode

import "strings"

//https://leetcode-cn.com/problems/word-pattern/
func wordPattern(pattern string, s string) bool {
	m := make(map[int]string)
	endm := make(map[string]bool)
	news := strings.Fields(s)
	if len(news) != len(pattern) {
		return false
	}
	for i, v := range []byte(pattern) {
		if v1, ok := m[int(v-'a')]; ok {
			if news[i] != v1 {
				return false
			}
		} else {
			if _, ok := endm[news[i]]; ok {
				return false
			}
			m[int(v-'a')] = news[i]
			endm[news[i]] = true
		}
	}
	return true
}
