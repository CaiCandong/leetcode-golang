/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	position:=len(nums1)-1
	p,q:=m-1,n-1
	for p>=0&&q>=0{
		if nums1[p]>nums2[q]{
			nums1[position]=nums1[p]
			p--
		}else{
			nums1[position]=nums2[q]
			q--
		}
			position--
	}
	for q>=0{
		nums1[position]=nums2[q]
		q--
		position--
	}
}
// @lc code=end

