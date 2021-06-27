package leetcode

//https://leetcode-cn.com/problems/flipping-an-image/
func flipAndInvertImage(A [][]int) [][]int {
	for i, v := range A {
		for lp, rp := 0, len(v)-1; lp <= rp; {
			if A[i][lp] == 1 {
				A[i][lp] = 0
			} else {
				A[i][lp] = 1
			}
			if A[i][rp] == 1 && lp != rp {
				A[i][rp] = 0
			} else if A[i][rp] == 0 && lp != rp {
				A[i][rp] = 1
			}
			A[i][lp], A[i][rp] = A[i][rp], A[i][lp]
			lp, rp = lp+1, rp-1
		}
	}
	return A
}
