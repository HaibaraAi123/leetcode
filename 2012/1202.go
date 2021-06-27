package leetcode

//https://leetcode-cn.com/problems/roman-to-integer/
func romanToInt(s string) int {
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	result := 0
	for i, v := range []byte(s) {
		if i == len(s)-1 {
			result += m[v]
			break
		}
		if m[v] < m[s[i+1]] {
			result -= m[v]
		} else if m[v] > m[s[i+1]] {
			result += m[v]
		} else if m[v] == m[s[i+1]] {
			result += m[v]
		}
	}
	return result
}
