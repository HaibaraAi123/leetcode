package leetcode

import "strings"

//https://leetcode-cn.com/problems/remove-duplicate-letters/
func removeDuplicateLetters(s string) string {
	if len(s) < 2 {
		return s
	}
	m := make([]int, 26)

	for _, v := range []byte(s) {
		m[v-'a']++
	}
	b := strings.Builder{}
	news := []byte(s)
	for lp := 0; lp < len(s); {
		if m[news[lp]-'a'] == 0 {
			lp++
			continue
		} else if m[news[lp]-'a'] == 1 {
			b.WriteByte(news[lp])
			m[news[lp]-'a'] = 0
			lp++
		} else if m[news[lp]-'a'] > 1 {
			min, minindex, rp, end := news[lp], lp, lp+1, true
			m[news[lp]-'a']--
			for ; rp < len(s); rp++ {
				if m[news[rp]-'a'] == 0 {
					continue
				}
				if m[news[rp]-'a'] == 1 {
					if news[rp] < min {
						min, minindex = news[rp], rp
						m[news[rp]-'a']--
						end = false
						break
					} else {
						end = false
						break
					}
				} else if m[news[rp]-'a'] > 1 && news[rp] < min {
					min, minindex = news[rp], rp
				}
				m[news[rp]-'a']--
				end = false
			}
			if end == true {
				return b.String()
			}
			b.WriteByte(min)
			if minindex == lp {
				for i := lp; i < rp; i++ {
					if m[news[i]-'a'] == 0 {
						continue
					}
					m[news[i]-'a']++
				}
			}
			m[news[minindex]-'a'] = 0
			lp = minindex + 1
		}
	}
	return b.String()
}
