package leetcode

import (
	"fmt"
	"strings"
)

//https://leetcode-cn.com/problems/summary-ranges/
func summaryRanges(nums []int) []string {
	result := make([]string, 0, len(nums))

	for lp, rp := 0, 0; rp < len(nums); {
		if rp != len(nums)-1 && nums[rp+1] == nums[rp]+1 {
			rp++
		} else {
			b := strings.Builder{}
			if rp == lp {
				b.WriteString(fmt.Sprint(nums[lp]))
				lp, rp = rp+1, rp+1
			} else {
				b.WriteString(fmt.Sprint(nums[lp], "->", nums[rp]))
				lp, rp = rp+1, rp+1
			}
			//strings.Join([]string{string(nums[lp]), "->", string(nums[rp])}, "")
			result = append(result, b.String())
		}
	}

	return result
}
