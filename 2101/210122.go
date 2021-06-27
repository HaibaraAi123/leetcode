package leetcode

import "fmt"

//https://leetcode-cn.com/problems/add-to-array-form-of-integer/
func addToArrayForm(A []int, K int) []int {
	add := 0
	for i := len(A) - 1; K >= 0 && i >= 0; i-- {
		tmp := (A[i] + add + K%10) / 10
		A[i] = (A[i] + add + K%10) % 10
		K = K / 10
		add = tmp
	}
	func1 := func(b int) []int {
		result := []int{}
		if b == 0 {
			result = append(result, 0)
			return result
		}
		for b > 0 {
			result = append(result, b%10)
			b = b / 10
		}
		if len(result) == 1 {
			return result
		}
		for lp, rp := 0, len(result)-1; lp <= rp; {
			result[lp], result[rp] = result[rp], result[lp]
			lp++
			rp--
		}
		return result
	}
	if K+add != 0 {
		return append(func1(K+add), A...)
	}
	fmt.Println(A)
	return A

}

/*
func addToArrayForm(A []int, K int) []int {
	fmt.Println(math.MaxInt64, math.MaxInt32)
	func1 := func(a []int) int {
		result := 0
		for i := 0; i < len(a); i++ {
			if i != 0 {
				result *= 10
			}
			result += a[i]
		}
		fmt.Println(result)
		return result
	}
	func2 := func(b int) []int {
		result := []int{}
		if b == 0 {
			result = append(result, 0)
			return result
		}
		for b > 0 {
			result = append(result, b%10)
			b = b / 10
		}
		if len(result) == 1 {
			return result
		}
		for lp, rp := 0, len(result)-1; lp <= rp; {
			result[lp], result[rp] = result[rp], result[lp]
			lp++
			rp--
		}
		return result
	}

	return func2(func1(A) + K)
}
*/
