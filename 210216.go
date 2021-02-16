package leetcode

//https://leetcode-cn.com/problems/array-partition-i/
/*
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	result := 0
	for i, v := range nums {
		if i%2 == 0 {
			result += v
		}
	}
	return result
}
*/
func arrayPairSum(nums []int) int {
	bucket, res := make([]int, 20001), 0
	for _, v := range nums {
		bucket[v+10000]++
	}
	for flag, i := true, 0; i < len(bucket)-1; {
		if bucket[i] > 0 {
			if flag {
				res += i - 10000
			}
			flag = !flag
			bucket[i]--
			continue
		} else {
			i++
		}
	}
	return res
}
