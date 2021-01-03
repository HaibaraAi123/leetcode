package leetcode

//https://leetcode-cn.com/problems/can-place-flowers/
func canPlaceFlowers(flowerbed []int, n int) bool {
	for p := 0; ; {
		if flowerbed[p] == 1 {
			p += 2
		} else if p == len(flowerbed)-1 {
			if flowerbed[p] != 1 {
				n--
				break
			}
		} else if flowerbed[p+1] != 1 {
			n--
			p += 2
		} else {
			p += 3
		}
		if p > len(flowerbed)-1 {
			break
		}
	}
	return n <= 0
}
