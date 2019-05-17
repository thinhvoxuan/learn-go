package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:  %s", e)
}

func abs(a float64, b float64) float64 {
	if a > b {
		return a - b
	}
	return b - a
}

//Sqrt
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(-2)
	}
	z := 1.0
	for abs(z*z, x) > 0.000000001 {
		z = z - (z*z-x)/(2*z)
	}
	return z, nil
}

// func main() {
// 	fmt.Println(Sqrt(2))
// 	fmt.Println(Sqrt(-2))
// }
