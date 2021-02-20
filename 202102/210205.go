package leetcode

//https://leetcode-cn.com/problems/get-equal-substrings-within-budget/
//s := "abc" t:="def" reflect.Typeof(s[0]).Name() -> uint8(rune)
//s[0] - t[0] != -(t[0] - s[0]) 溢出
func equalSubstring(s string, t string, maxCost int) int {
	maxL := 0
	diff := make([]int, len(s))
	maxfunc := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	absfunc := func(i int) int {
		if i < 0 {
			return -1 * i
		}
		return i
	}
	for i, v := range []byte(s) {
		diff[i] = absfunc(int(t[i]) - int(v))
	}
	for lp, rp, cost := 0, 0, 0; rp < len(s); {
		if cost+diff[rp] > maxCost {
			if rp != lp {
				maxL = maxfunc(maxL, rp-lp)
				cost -= diff[lp]
				lp++
			} else {
				lp++
				rp++
				cost = 0
			}
		} else {
			if rp == len(s)-1 && rp != lp {
				maxL = maxfunc(maxL, rp-lp+1)
				break
			}
			cost += diff[rp]
			rp++
		}
	}
	return maxL
}
