package leetcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

func Test210407(t *testing.T) {
	type test struct {
		nums   []int
		target int
		want   bool
	}

	testGroup := []test{
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
	}

	for _, v := range testGroup {
		got := search(v.nums, v.target)
		if v.want != got {
			t.Errorf("error at test:%v %v,want:%v, got:%v", v.nums, v.target, v.want, got)
		}
	}
}
func Benchmark210407(b *testing.B) {
}
func Test040701(t *testing.T) {
	/*
		为每个数字定义一个只有”Y“,"N"的编码方式，每一位表示某个数字是i的倍数，
		i从左往右，从1开始，比如 "YNYYNNY" 表示这个数是1，3，4，7的倍数，但不是2，5，6的倍数。
		现给定正整数 L, 表示某数字的编码后长度，请求出合法的编码的个数总和。
		提示：不合法表示符合编码的数字不存在，比如 "YNNY" 是不合法的，因为不存在是4的倍数但是不是2的倍数的数字
	*/

}
func Test040702(t *testing.T) {
	var s, a, b string
	s = "10;11"
	a, b = strings.Split(s, ";")[0], strings.Split(s, ";")[1]
	resa, resb := 0, 0
	lena, lenb := len(a), len(b)
	for i, v := range []byte(a) {
		resa += int(v-'0') << (lena - i - 1)
	}
	for i, v := range []byte(b) {
		resb += int(v-'0') << (lenb - i - 1)
	}
	res := resa * resb
	println(strconv.FormatInt(int64(res), 2))
	for i := int(math.Log2(float64(res))); i >= 0; i-- {
		tmp := 1 << i
		if res >= (1 << i) {
			fmt.Printf("1")
			res -= tmp
		} else {
			fmt.Printf("0")
		}

	}
}
func Test040703(t *testing.T) {
	/*
		有限小数或者无限循环小数都可以转化为分数。
		比如：0.9=9/10 0.333(3)=1/3（括号中的数字表示是循环节）
		当然一个小数可以用好几种分数形式来表示。如：0.333(3)=1/3=3/9 给定一个有限小数或无限循环小数，
		你能否以分母最小的分数形式来返回这个小数呢？如果输入为循环小数，循环节用括号标记出来。
	*/
	var s string = "0.3(333)"
	var a, b float64
	a, _ = strconv.ParseFloat(strings.Split(s, "(")[0], 10)

	b, _ = strconv.ParseFloat(strings.Split(strings.Split(s, "(")[1], ")")[0], 10)
	println(a, b)
}
