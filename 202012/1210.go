package leetcode

//https://leetcode-cn.com/problems/lemonade-change/
func lemonadeChange(bills []int) bool {
	result := [4]int{}
	for _, v := range bills {
		result[v/5-1]++
		switch v {
		case 5:
			continue
		case 10:
			{
				if result[0]--; result[0] < 0 {
					return false
				}
			}
		case 20:
			{
				if result[1] > 0 {
					result[1]--
					result[0]--
				} else {
					result[0] -= 3
				}
				if result[0] < 0 {
					return false
				}
			}
		}
	}
	return true
}
