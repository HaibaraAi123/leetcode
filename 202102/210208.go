package leetcode

import "fmt"

//https://leetcode-cn.com/problems/longest-turbulent-subarray/
func maxTurbulenceSize(arr []int) int {
	maxLen := 1
	maxfunc := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	if len(arr) <= 1 {
		return len(arr)
	}
	for lp, rp, flag := 0, 1, true; lp < len(arr) && rp < len(arr); {
		if arr[rp-1] == arr[rp] {
			maxLen = maxfunc(maxLen, rp-lp)
			fmt.Println(lp, rp, "1")
			lp++
			rp = lp + 1
			continue
		}
		if rp == lp+1 {
			flag = arr[lp] > arr[rp]
			rp++
			if rp == len(arr) {
				return maxfunc(maxLen, rp-lp)
			}
			continue
		}
		if arr[rp-1] > arr[rp] != flag {
			fmt.Println(lp, rp, "3", flag)
			flag = arr[rp-1] > arr[rp]
			rp++
			if rp == len(arr) {
				return maxfunc(maxLen, rp-lp)
			}
			continue
		} else {
			maxLen = maxfunc(maxLen, rp-lp)
			fmt.Println(lp, rp, "2")
			lp++
			rp = lp + 1
			continue
		}
	}
	return maxLen
}
