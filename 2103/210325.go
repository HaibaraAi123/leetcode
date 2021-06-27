package leetcode

const a = 1.0
const b = 2.0
const h = 3.0

var High = [][]float64{{0, b, 2 * b, 3 * b, 4 * b, 5 * b, 6 * b}, {h, h, h, h, h, h, h}}

var Low = [][]float64{{a, a + b, a + 2*b, a + 3*b, a + 4*b, a + 4*b, a + 5*b}, {0, 0, 0, 0, 0, 0, 0}}

func Point(x, y float64, a []float64) (float64, float64) {
	resx, resy := 0.0, 0.0
	return resx, resy
}
func UpdatePoint() {
	var n int = 5
	for i := 0; i <= n; i++ {
		var tmp = []float64{High[0][i], High[1][i], Low[0][i+1], Low[1][i+1]}
		for vi := range High[0] {
			if vi <= i {
				continue
			}
			High[0][vi], High[1][vi] = Point(High[0][vi], High[1][vi], tmp)
		}
		for vi := range Low[0] {
			if vi <= i+1 {
				continue
			}
			Low[0][vi], Low[1][vi] = Point(Low[0][vi], Low[1][vi], tmp)
		}
	}
}
