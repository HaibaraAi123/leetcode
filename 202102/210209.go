package leetcode

//https://leetcode-cn.com/problems/subarrays-with-k-different-integers/
//O(n*n)
func subarraysWithKDistinct(A []int, K int) int {
	result := 0
	for winLen := K; winLen <= len(A); winLen++ {
		m := make(map[int]int, winLen)
		sumK := 0
		for lp, rp := 0, winLen-1; rp < len(A); {
			if lp != 0 {
				if v1 := m[A[lp-1]]; v1 == 1 {
					sumK--
				}
				m[A[lp-1]]--
				if v, ok := m[A[rp]]; !ok || v == 0 {
					sumK++
				}
				m[A[rp]]++
			} else {
				for i := lp; i <= rp; i++ {
					if _, ok := m[A[i]]; !ok {
						sumK++
					}
					m[A[i]]++
				}
			}
			if sumK == K {
				result++
				//fmt.Println(A[lp : rp+1])
			}
			lp++
			rp++
		}
	}
	return result
}
