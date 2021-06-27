package leetcode

//https://leetcode-cn.com/problems/max-consecutive-ones/
func findMaxConsecutiveOnes(nums []int) int {
	result, lp := 0, 0
	maxfunc := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for i, v := range nums {
		if lp != i {
			if v == 1 && i != len(nums)-1 {
				continue
			} else if v == 1 && i == len(nums)-1 {
				result = maxfunc(result, i-lp+1)
			} else {
				result = maxfunc(result, i-lp)
				lp = i + 1
			}
		} else if v == 1 {
			if i == len(nums)-1 {
				result = maxfunc(result, i-lp+1)
			} else {
				continue
			}
		} else {
			lp = i + 1
		}
	}
	return result
}
