/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	// 动态规划 dp[i] = j
	// 表示不包括 第 i 个字符 的 前 j 个字符是有效子串
	max := 0
	dp := make([]int, len(s))
	for i, l := 1, len(s); i < l; i++ {
		if s[i] == ')' {
			if k := i - dp[i] - 1; k >= 0 && s[k] == '(' {
				m := dp[i] + 2 + dp[k]
				if i+1 < l {
					dp[i+1] = m
				}
				if m > max {
					max = m
				}
			}
		}
	}
	return max
}

// @lc code=end

