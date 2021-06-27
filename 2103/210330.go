package leetcode

import "fmt"

//https://leetcode-cn.com/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	transfunc := func(cnt int) (i, j int) {
		i = cnt / n
		j = cnt % n

		fmt.Println(cnt, i, j)
		return
	}
	if target > matrix[m-1][n-1] || target < matrix[0][0] {
		return false
	} else if target == matrix[m-1][n-1] || target == matrix[0][0] {
		return true
	}
	for lp, rp := 0, m*n-1; rp-lp != 1; {
		if rp-lp == 1 {
			return false
		}
		mid := (lp + rp) / 2
		i, j := transfunc(mid)
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			rp = mid
		} else {
			lp = mid
		}
	}
	return false
}
