package leetcode

import (
	"log"
	"reflect"
	"runtime"
	"testing"
)

func Test210213(t *testing.T) {
	type test struct {
		nums, want []int
	}

	testGroup := []test{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
		{[]int{1, 1}, []int{2}},
	}
	TraceMemStatusFunc := func() uint64 {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		return m.Alloc
	}
	TestAlloc := TraceMemStatusFunc()
	for i, v := range testGroup {
		got := findDisappearedNumbers(v.nums)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at test:%d, want:%v, got %v", i, v.want, got)
		}
	}
	log.Println(TraceMemStatusFunc()-TestAlloc, "Byte")
}
func Benchmark210213(b *testing.B) {
}
