package leetcode

//https://leetcode-cn.com/problems/monotonic-array/
func isMonotonic(A []int) bool {
	if len(A) <= 1 {
		return true
	}
	flag := false
	for lp, tmp := 1, 0; lp < len(A); lp++ {
		if A[lp-1] == A[lp] {
			continue
		}
		if tmp == 0 {
			flag = A[lp-1] > A[lp]
			tmp = 1
			continue
		} else if flag != (A[lp-1] > A[lp]) {
			return false
		}
	}
	return true
}
