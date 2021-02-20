package leetcode

import "testing"

func Test1205(t *testing.T) {
	type test struct {
		tasks []byte
		n     int
		want  int
	}
	testGroup := []test{
		{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 2, want: 8},
		{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B'}, n: 0, want: 6},
		{tasks: []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, n: 2, want: 16},
		{tasks: []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'}, n: 2, want: 12},
	}
	for _, v := range testGroup {
		got := leastInterval(v.tasks, v.n)
		if got != v.want {
			t.Errorf("error at %v n=%v want %v got %v", v.tasks, v.n, v.want, got)
		}
	}
}
func Benchmark1205(b *testing.B) {

}
