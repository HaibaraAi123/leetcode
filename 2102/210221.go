package leetcode

import "fmt"

//https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
func longestSubarray(nums []int, limit int) int {
	res := 0
	maxfunc := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	absfunc := func(i int) int {
		if i > 0 {
			return i
		}
		return -i
	}
	for lp, rp := 0, 0; rp < len(nums); rp++ {
		flag, i := false, rp-1
		for ; i >= lp; i-- {
			if absfunc(nums[i]-nums[rp]) > limit {
				fmt.Println(res, lp, rp)
				res = maxfunc(res, rp-lp)
				flag = true
				break
			}
		}
		if rp == len(nums)-1 && flag == false {
			fmt.Println(res, lp, rp, "rp=num-1")
			res = maxfunc(res, rp-lp+1)
			break
		}
		if flag == true {
			lp = i + 1
			continue
		}

	}
	return res
}
