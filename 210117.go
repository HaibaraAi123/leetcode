package leetcode

//https://leetcode-cn.com/problems/check-if-it-is-a-straight-line/
func checkStraightLine(coordinates [][]int) bool {
	if coordinates[0][0] == coordinates[1][0] {
		for _, v := range coordinates {
			if v[0] != coordinates[0][0] {
				return false
			}
		}
		return true
	} else if coordinates[0][1] == coordinates[1][1] {
		for _, v := range coordinates {
			if v[1] != coordinates[0][1] {
				return false
			}
		}
		return true
	}
	nums := make([][]float32, len(coordinates))
	for i := 0; i < len(coordinates); i++ {
		nums[i] = []float32{float32(coordinates[i][0]), float32(coordinates[i][1])}
	}
	var k float32 = (nums[1][1] - nums[0][1]) / (nums[1][0] - nums[0][0])
	var b float32 = nums[0][1] - k*nums[0][0]
	for _, v := range nums {
		if v[1] != k*v[0]+b {
			return false
		}
	}
	return true
}
