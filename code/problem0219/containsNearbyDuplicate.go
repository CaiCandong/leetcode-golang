package problem0219

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, 0)
	for i, num := range nums {
		j, ok := m[num]
		if ok && (i-j <= k || j-i <= k) {
			return true
		}
		m[num] = i
	}
	return false
}
