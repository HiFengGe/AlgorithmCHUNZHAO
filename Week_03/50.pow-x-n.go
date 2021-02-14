/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {

	//终止条件
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / (x * myPow(x, -(n+1)))
	}
	result := 1.0
	for n > 1 {
		if n%2 == 1 {
			result *= x
		}
		x = x * x
		n /= 2
	}

	//当前状态处理
	result *= x
	return result

}

// @lc code=end

