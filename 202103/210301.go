package leetcode

//https://leetcode-cn.com/problems/range-sum-query-immutable/

//NumArray [][]int
type NumArray struct {
	array []int
}

//Constructor 构造NumArray
func Constructor(nums []int) NumArray {
	array := make([]int, len(nums))
	for i := range array {
		if i != 0 {
			array[i] += array[i-1]
		}
		array[i] += nums[i]
	}
	return NumArray{array}
}

//SumRange sum nums[i]--nums[j]
func (n *NumArray) SumRange(i int, j int) int {

	if i == 0 {
		return n.array[j]
	}
	return n.array[j] - n.array[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
