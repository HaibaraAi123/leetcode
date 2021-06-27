package leetcode

//https://leetcode-cn.com/problems/clumsy-factorial/
func clumsy(N int) int {
	res := 0
	for i := N; i > 0; {
		if i == N {
			switch i {
			case 1:
				res += 1
				return res
			case 2:
				res += 2 * 1
				return res
			case 3:
				res += 3 * 2 / 1
				return res
			default:
				res += i*(i-1)/(i-2) + i - 3
				i -= 4
			}
		} else {
			switch i {
			case 1:
				res -= 1
				return res
			case 2:
				res -= 2 * 1
				return res
			case 3:
				res -= 3 * 2 / 1
				return res
			default:
				res -= i * (i - 1) / (i - 2)
				res += i - 3
				i -= 4
			}
		}
	}
	return res
}
