package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	return helper(x, n)
}
func helper(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n%2 == 0 {
		temp := helper(x, n/2)
		return temp * temp
	}
	if n%2 == 1 {
		temp := helper(x, n/2)
		return temp * temp * x
	}
	return 0
}
func main() {
	fmt.Print(myPow(2.0, -2))
}
