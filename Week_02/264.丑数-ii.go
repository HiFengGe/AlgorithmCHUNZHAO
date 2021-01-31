/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	// dp 动态规划 (三指针)先排再存
	//  小顶堆的方法是先存再排，dp 的方法则是先排再存
	// 我们设 3 个指针 i,j,k
	// 代表的是第几个数的2倍、第几个数 3 倍、第几个数 5 倍
	// 动态方程最小值 dp[i]=min(dp[i]*2,dp[j]*3,dp[k]*5) =>  dp[i]=min(min(dp[i]*2,dp[j]*3),dp[k]*5)
	// 小顶堆是一个元素出来然后存 3 个元素
	// 动态规划则是标识 3 个元素，通过比较他们的 2 倍、3 倍、5 倍的大小，来一个一个存

	if n == 1 {
		return 1
	}

	result := []int{1}
	res := 0
	i, j, k := 0, 0, 0
	for n > 1 {
		x := result[i] * 2
		y := result[j] * 3
		z := result[k] * 5

		res = min(min(x, y), z)
		if res == x {
			i++
		}

		if res == y {
			j++
		}

		if res == z {
			k++
		}
		result = append(result, res)
		n--
	}

	// 最后一位
	return result[len(result)-1]

}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
// @lc code=end

