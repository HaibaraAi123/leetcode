package leetcode

import "math"

//https://leetcode-cn.com/problems/maximum-points-you-can-obtain-from-cards/
func maxScore(cardPoints []int, k int) int {
	sum, minCount := 0, math.MaxInt32
	for _, v := range cardPoints {
		sum += v
	}
	for lp, rp, count := 0, len(cardPoints)-k-1, 0; rp < len(cardPoints); {
		if lp == 0 {
			for i := 0; i <= rp; i++ {
				count += cardPoints[i]
			}
		} else {
			count = count + cardPoints[rp] - cardPoints[lp-1]
		}
		if count < minCount {
			minCount = count
		}
		lp++
		rp++
	}
	return sum - minCount
}
