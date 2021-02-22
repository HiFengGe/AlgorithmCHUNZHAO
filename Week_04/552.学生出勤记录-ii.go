/*
 * @lc app=leetcode.cn id=552 lang=golang
 *
 * [552] 学生出勤记录 II
 */

// @lc code=start
func checkRecord(n int) int {
	mod := int(math.Pow(10, 9) + 7)
	if n == 1 {
		return 3
	}
	if n == 2 {
		return 8
	}

	dp := [][]int{}
	dp[0][0] = 1
	dp[0][1] = 1
	dp[1][0] = 2
	dp[1][1] = 2

	for i := 2; i < n; i++ {
		dp[i][0] = (dp[i-1][0] + dp[i-1][1]) % mod
		dp[i][1] = (dp[i-1][0] + dp[i-2][0]) % mod
	}

	ans := ((dp[n-2][0]+dp[n-2][1])*2%mod + dp[n-1][0] + dp[n-1][1]) % mod
	for i := 1; i < n-1; i++ {
		ans += (dp[i-1][0] + dp[i-1][1]) * (dp[n-i-2][0] + dp[n-i-2][1])
		ans %= mod
	}
	return ans
}

// @lc code=end

