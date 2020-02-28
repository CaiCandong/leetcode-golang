package problem0172

func trailingZeroes(n int) int {
	count5 := 0
	for n >= 5 {
		count5 += n / 5
		n /= 5
	}
	return count5
}
