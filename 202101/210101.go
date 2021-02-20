package leetcode

//https://leetcode-cn.com/problems/can-place-flowers/
/*
防御式编程思想：在 flowerbed 数组两端各增加一个 0，
这样处理的好处在于不用考虑边界条件，任意位置处只要连续出现三个 0 就可以栽上一棵花。
*/
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
