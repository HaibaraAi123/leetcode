package leetcode

//https://leetcode-cn.com/problems/max-consecutive-ones-iii/
func longestOnes(A []int, K int) int {
	res := 0
	maxfunc := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for lp, rp, cnt0 := 0, 0, 0; rp < len(A); {
		if rp == len(A)-1 {
			if (A[rp] == 0 && cnt0 < K) || A[rp] == 1 {
				res = maxfunc(res, rp-lp+1)
			}
		}
		if A[rp] == 0 && cnt0 < K {
			rp, cnt0 = rp+1, cnt0+1
			continue
		} else if A[rp] == 1 {
			rp++
			continue
		} else if A[rp] == 0 && cnt0 == K {
			res = maxfunc(res, rp-lp)
			for ; lp <= rp; lp++ {
				if A[lp] == 0 {
					lp++
					break
				}
			}
			rp++
			continue
		}
	}
	return res
}
func longestOnesTest(A []int, K int) int {
	lp, cnt0 := 0, 0
	for _, v := range A {
		if v == 0 {
			cnt0++
		}
		if cnt0 > K {
			if A[lp] == 0 {
				cnt0--
			}
			lp++
		}
	}
	return len(A) - lp
}
