package leetcode

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(stones)))
	for lp, rp := 0, 1; ; {
		if stones[lp] != stones[rp] {
			v := stones[lp] - stones[rp]
			for i := rp + 1; i <= len(stones); i++ {
				if i == len(stones) {
					stones = append(stones, v)
					break
				}
				if v > stones[i] {
					v, stones[i] = stones[i], v
				}
			}
			fmt.Println(stones, "add")
		}
		lens := len(stones)
		if rp == lens-1 {
			return 0
		} else if rp+1 == lens-1 {
			return stones[rp+1]
		}
		lp, rp = lp+2, rp+2
	}
}
