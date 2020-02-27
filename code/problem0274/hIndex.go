package problem0274

import (
	"sort"
)

func HIndex(citations []int) int {
	return hIndex(citations)
}

func hIndex(citations []int) int {
	sort.Ints(citations)
	res := 0
	for i := 1; i < len(citations); i++ {
		if citations[len(citations)-1-i] >= i {
			res = i
		}
	}
	return res
}
