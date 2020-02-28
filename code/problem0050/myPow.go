package problem0050

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	return helper(x, n)
}
func helper(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n&1 == 0 {
		temp := helper(x, n/2)
		return temp * temp
	} else {
		temp := helper(x, n/2)
		return temp * temp * x
	}
}
