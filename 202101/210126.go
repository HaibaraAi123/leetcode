package leetcode

//https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs/
func numEquivDominoPairs(dominoes [][]int) int {
	m := make(map[int]int, len(dominoes))
	var result int
	for _, v := range dominoes {
		lp, rp := v[0], v[1]
		if lp > rp {
			lp, rp = rp, lp
		}

		m[lp*10+rp]++
	}
	for _, v := range m {
		if v > 1 {
			result += v * (v - 1) / 2
		}
	}
	return result
}
