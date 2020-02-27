package problem0189

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums[:k+1])
	reverse(nums[k+1:])
	reverse(nums)
}
func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
