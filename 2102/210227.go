package leetcode

//https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters/
func longestSubstring(s string, k int) int {
	m := make([]int, 26)
	for _, v := range s {
		m[int(v-'a')]++
	}
	res := 0
	/*
		maxfunc := func(i, j int) int {
			if i > j {
				return i
			}
			return j
		}
		news := []byte(s)
	*/
	return res
}
