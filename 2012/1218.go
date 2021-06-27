package leetcode

//https://leetcode-cn.com/problems/find-the-difference/submissions/
func findTheDifference(s string, t string) byte {
	m := make([]int, 26)
	for _, v := range []byte(s) {
		m[v-'a']++
	}
	for _, v := range []byte(t) {
		if m[v-'a'] == 0 {
			return v
		}
		m[v-'a']--
	}
	return ' '
}
