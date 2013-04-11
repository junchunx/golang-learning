package newmath

//This is a terrible implementation
//Real code should import "math" and use math.Sqrt
func Sqrt(x float64) float64 {
	y := 0.0
	for i := 0; i < 1000; i++ {
		y -= (y*y - x) / (2 * x)
	}
	return y
}
