package problem08_11

func waysToChange(n int) int {
	return waysToChange1(n)
}

func waysToChange1(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	coins := []int{1, 5, 10, 25}
	for _, coin := range coins {
		for i, _ := range dp {
			if i+coin > n {
				break
			}
			dp[i+coin] = (dp[i+coin] + dp[i]) % 1000000007
		}
	}
	return dp[n]
}
