/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) int {
	// 动态规划：
	// 1.寻找边界
	// 2.找出递推关系式
	dp := make([][]int, len(s))
	for k := range dp {
		dp[k] = make([]int, len(s)) //初始化DP数组
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = 1 //递推边界
	}
	for j := 1; j < len(s); j++ {
		dp[0][j] = dp[0][j-1] + Count(j, s) //递推等式
	}
	return dp[0][len(s)-1]
}

func Count(j int, s string) int {
	var count = 0
	for k := 0; k <= j; k++ {
		count += Judge(s, k, j)
	}
	return count
}
func Judge(s string, i, j int) int {
	s = s[i : j+1]
	length := len(s)
	for a, b := 0, length-1; a < b; a, b = a+1, b-1 {
		if s[a] != s[b] {
			return 0
		}
	}
	return 1
}

// @lc code=end

