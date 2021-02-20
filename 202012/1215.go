package leetcode

import (
	"math"
)

func int2slice(n int) []int {
	result := []int{}
	for tmp := 0; ; {
		tmp = n % 10
		n = n / 10
		if n < 10 {
			result = append(result, tmp)
			result = append(result, n)
			return result
		}
		result = append(result, tmp)
		continue
	}
}
func monotoneIncreasingDigits(N int) int {
	tmp := int2slice(N)
	index := len(tmp) - 1
	for ; index > 0; index-- {
		if tmp[index] > tmp[index-1] {
			break
		}
	}
	if index == 0 {
		return N
	}
	for i := index; i < len(tmp); {
		if tmp[i] != 0 {
			tmp[i]--
			for j := i - 1; j >= 0; j-- {
				tmp[j] = 9
			}
			if i == len(tmp)-1 {
				break
			} else if tmp[i] >= tmp[i+1] {
				break
			}
		} else {
			i++
			continue
		}
	}
	result := 0
	for i, v := range tmp {
		result += v * int(math.Pow(10.0, float64(i)))
	}
	return result
}

/*
func monotoneIncreasingDigits(n int) int {
    s := []byte(strconv.Itoa(n))
    i := 1
    for i < len(s) && s[i] >= s[i-1] {
        i++
    }
    if i < len(s) {
        for i > 0 && s[i] < s[i-1] {
            s[i-1]--
            i--
        }
        for i++; i < len(s); i++ {
            s[i] = '9'
        }
    }
    ans, _ := strconv.Atoi(string(s))
    return ans
}

*/
