package leetcode

import (
	"fmt"
	"testing"
)

func Test1202(t *testing.T) {
	v := romanToInt("MCMXCIV")
	t.Log(v)
	if romanToInt("III") != 3 {
		t.Errorf("测试失败")
	}
}
func Benchmark1202(b *testing.B) {
	fmt.Println(romanToInt(("MCMXCIV")))
}
