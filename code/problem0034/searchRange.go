package problem0034

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	return searchRange3(nums, target)
}
func findLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var flag bool
	for left <= right {
		mid := (left + right) / 2
		fmt.Println(left, right, mid)
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			flag = true
			right = mid - 1
		}
	}
	if flag {
		return right + 1
	} else {
		return -1
	}
}
func findRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var flag bool
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			flag = true
			left = mid + 1
		}
	}
	if flag {
		return left - 1
	} else {
		return -1
	}
}

//解法3:二分查找:找到最右数字
func searchRange3(nums []int, target int) []int {
	return []int{findLeft(nums, target), findRight(nums, target)}
}

//解法2:二分查找+线性搜索
func searchRange2(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	index := -1
	for left <= right {
		mid := (left + right) / 2
		fmt.Println("left:", left, "right:", right, "mid", mid)
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			index = mid
			break
		}
	}
	if index == -1 {
		return []int{-1, -1}
	} else {
		i := index
		for i > 0 && nums[i-1] == nums[i] {
			i--
		}
		j := index
		for j < len(nums)-1 && nums[j+1] == nums[j] {
			j++
		}
		return []int{i, j}
	}
}

//解法1:线性搜索
func searchRange1(nums []int, target int) []int {
	res1, res2 := -1, -1
	for i, num := range nums {
		if num == target {
			res1 = i
			break
		}
	}
	if res1 == -1 {
		return []int{res1, res2}
	}
	res2 = res1
	for res2+1 < len(nums) && nums[res2+1] == nums[res1] {
		res2++
	}
	return []int{res1, res2}
}
