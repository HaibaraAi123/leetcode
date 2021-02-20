package leetcode

//https://leetcode-cn.com/problems/count-primes/submissions/
func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
func countPrimes(n int) int {
	//厄拉多塞筛法
	if n < 3 {
		return 0
	}
	count := 0
	m := make([]int, n)
	m[0], m[1] = 0, 0
	for i := 2; i < n; i++ {
		m[i] = 1
	}
	for i := 2; i < n; i++ {
		if m[i] == 1 {
			for cnt := 2; cnt*i < n; cnt++ {
				m[i*cnt] = 0
			}
		}
	}
	for _, v := range m {
		if v == 1 {
			count++
		}
	}
	return count

	/*
		for i := 2; i < n; i++ {
			if isPrime((i)) {
				count++
			}
		}*/

}
