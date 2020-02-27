package problem0034

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	res := findLeft([]int{5, 5, 5, 6, 7, 7}, 7)
	fmt.Println(res)
}
