/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
func countPrimes(n int) int {
	var count int
	flag:=make([]bool,n+1)
	for i:=2;i<n;i++{
		if !flag[i]{
		for j:=2;j*i<=n;j++{
			flag[j*i]=true//标记为不是质数
		}
		count++
		}
	}
	return count
}
// @lc code=end

