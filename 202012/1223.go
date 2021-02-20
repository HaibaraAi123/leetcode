package leetcode

//https://leetcode-cn.com/problems/first-unique-character-in-a-string/
func firstUniqChar(s string) int {
	m := make([]int, 26)
	for _, v := range []byte(s) {
		m[v-'a']++
	}
	for i, v := range []byte(s) {
		if m[v-'a'] == 1 {
			return i
		}
	}
	return -1
}
