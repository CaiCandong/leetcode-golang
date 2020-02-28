package main

import "fmt"

func trailingZeroes(n int) int {
	count5 := 0
	for n >= 5 {
		count5 += n / 5
		n /= 5
	}
	return count5
}
func main() {
	res := trailingZeroes(5)
	fmt.Println(res)
}
