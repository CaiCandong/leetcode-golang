package problem0031

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i == -1 {
		reverse(nums)
		return
	}
	j := i
	for j < len(nums)-1 && nums[j+1] >= nums[i] {
		j++
	}
	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums[i+1:])
}
func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
