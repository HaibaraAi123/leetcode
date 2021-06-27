package leetcode

import (
	"fmt"
	"reflect"
)

//https://leetcode-cn.com/problems/permutation-in-string/
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	m := make([]int, 26)
	for _, v := range s1 {
		m[v-'a']++
	}
	fmt.Println(m)
	for lp, rp := 0, len(s1)-1; rp < len(s2); {
		tmp := make([]int, 26)
		flag := false
		for i := lp; i <= rp; i++ {
			if m[s2[i]-'a'] == 0 {
				lp = i + 1
				rp = i + len(s1)
				flag = true
				break
			}
			tmp[s2[i]-'a']++
		}
		if flag {
			continue
		} else if reflect.DeepEqual(tmp, m) {
			fmt.Println(tmp)
			return true
		}
		lp++
		rp++
	}
	return false
}
