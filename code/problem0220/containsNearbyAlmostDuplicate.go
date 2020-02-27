package problem0220

import "sort"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	return containsNearbyAlmostDuplicate1(nums, k, t)
}

//解法3:桶排序
//思路:桶排序,抽屉原则,设置一个数值范围为t的抽屉
func containsNearbyAlmostDuplicate3(nums []int, k, t int) bool { //k 索引差 t 数值差
	size := len(nums)
	if k == 0 || t < 0 || size <= 1 {
		return false
	}
	t++
	nMap := make(map[int]int, size)
	var ni, m int
	for i := 0; i < size; i++ {
		ni = nums[i]
		m = ni / t
		if ni < 0 {
			m--
		}
		if _, ok := nMap[m]; ok {
			// 说明存在 nums[j]== n, j < i，使得
			// m*t <= ni, n < m*(t+1) 成立
			// 即，|ni-n|<t 成立
			return true
		} else if n, ok := nMap[m-1]; ok && abs(ni, n) < t {
			// 说明存在 nums[j]== n, j < i，使得
			// n < m*t <= ni 成立
			// 这也还是有可能使得 |ni-n|<t 成立的
			// 需要检查 abs(ni, n)
			return true
		} else if n, ok := nMap[m+1]; ok && abs(ni, n) < t {
			// 说明存在 nums[j]== n, j < i，使得
			// ni < m*(t+1) <= n 成立
			// 这也还是有可能使得 |ni-n|<t 成立的
			// 需要检查 abs(ni, n)
			return true
		}
		nMap[m] = ni
		if i >= k {
			delete(nMap, nums[i-k]/t)
		}
	}

	return false
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

//解法2:线性搜索,按元素的值作为依据进行排序,线性搜索索引值
//思路: 基数排序,优先级排序
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	type Pair struct {
		Value int
		Index int
	}
	in := make([]Pair, 0)
	for i, num := range nums {
		in = append(in, Pair{num, i})
	}
	sort.Slice(in, func(i, j int) bool {
		if in[i].Value < in[j].Value {
			return true
		}
		return false
	})
	for i := 0; i < len(in)-1; i++ {
		for j := 1; in[i+j].Value <= in[i].Value+t; j++ {
			if in[i+1].Index-in[i].Index <= k && in[i].Index-in[i+1].Index <= k {
				return true
			}
		}
	}
	return false
}

//解法1:暴力
func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	for i, num1 := range nums {
		for j, num2 := range nums {
			//if i+j>=len(nums){
			//	break
			//}
			if i != j && num1-num2 <= t && num2-num1 <= t && i-j <= k && j-i <= k {
				return true
			}
		}
	}
	return false
}
