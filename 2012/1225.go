package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	leng, lens := len(g), len(s)
	sort.Ints(g)
	sort.Ints(s)
	sp, gp := 0, 0
	for sp < lens && gp < leng {
		if s[sp] >= g[gp] {
			gp++
		}
		sp++
	}
	return gp
}
