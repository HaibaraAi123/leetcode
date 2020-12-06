package leetcode

//https://leetcode-cn.com/problems/pascals-triangle/
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	result := make([][]int, numRows)
	for i := range result {
		if i == 0 {
			result[i] = append(result[i], 1)
			continue
		}
		result[i] = append(result[i], 1)
		for j := 1; j < i; j++ {
			result[i] = append(result[i], result[i-1][j-1]+result[i-1][j])
		}
		result[i] = append(result[i], 1)
	}
	return result
}
