package leetcode

//https://leetcode-cn.com/problems/grumpy-bookstore-owner/
func maxSatisfied(customers []int, grumpy []int, X int) int {
	res := 0
	maxfunc := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	for lp, rp, cnt := 0, X-1, 0; rp < len(customers); {
		if lp == 0 {
			for i := 0; i <= rp; i++ {
				if grumpy[i] == 1 {
					res += customers[i]
				}
			}
			lp, rp = lp+1, rp+1
			cnt = res
			continue
		}
		cnt += grumpy[rp] * customers[rp]
		cnt -= grumpy[lp-1] * customers[lp-1]
		lp, rp = lp+1, rp+1
		res = maxfunc(res, cnt)
	}
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			res += customers[i]
		}
	}
	return res
}
