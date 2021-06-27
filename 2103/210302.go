package leetcode

//https://leetcode-cn.com/problems/range-sum-query-2d-immutable/

//NumMatrix sum(n)
type NumMatrix struct {
	Num [][]int
}

//Constructor02 构造NumMatrix
func Constructor02(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{[][]int{}}
	}
	row, col := len(matrix), len(matrix[0])
	res := make([][]int, row)
	for i := 0; i < row; i++ {
		res[i] = make([]int, col)
		for j := 0; j < col; j++ {
			if j != 0 {
				res[i][j] += res[i][j-1]
			}
			res[i][j] += matrix[i][j]
		}
	}
	return NumMatrix{res}
}

//SumRegion 计算区域和
func (n *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	res := 0
	for i := row1; i <= row2; i++ {
		if col1 == 0 {
			res += n.Num[i][col2]
			continue
		}
		res = res + n.Num[i][col2] - n.Num[i][col1-1]
	}
	return res
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
