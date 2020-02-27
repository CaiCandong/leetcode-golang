package problem0035

func searchInsert(nums []int, target int) int {
	//return solutionOne(nums,target)
	return solutionTwo(nums, target)
}

// O(n) 解法
func solutionOne(nums []int, target int) int {
	i := 0
	for i < len(nums) && nums[i] < target {
		i++
	}
	return i
}

// O(log(n)) 二分算法
func solutionTwo(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
