package leetcode

import (
	"sort"
	"strings"
)

//https://leetcode-cn.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	} else if len(strs) == 0 {
		return ""
	}
	sort.Sort(sort.StringSlice(strs))

	for i := len(strs[0]); i > 0; i-- {
		flag := true
		for j := 1; j < len(strs); j++ {
			if !strings.HasPrefix(strs[j], strs[0][0:i]) {
				flag = false
				break
			}
		}
		if flag == true {
			return strs[0][0:i]
		}
	}
	return ""
}
