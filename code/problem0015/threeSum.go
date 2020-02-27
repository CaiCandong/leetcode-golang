package problem0015

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	res := make([][]int, 0)

	for left := 0; left <= len(nums)-3; left++ {
		target := -nums[left]
		center, right := left+1, len(nums)-1
		for center < right && left < center {
			switch {
			case nums[center]+nums[right] < target:
				center++
			case nums[center]+nums[right] > target:
				right--
			case nums[center]+nums[right] == target:
				res = append(res, []int{nums[left], nums[center], nums[right]})
				for center < len(nums)-1 && nums[center] == nums[center+1] {
					center++
				}
				center++
			}
		}
		for left < len(nums)-1 && nums[left] == nums[left+1] {
			left++
		}
	}
	return res
}
