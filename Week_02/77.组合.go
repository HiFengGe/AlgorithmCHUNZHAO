/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	// 不能有重复的排列，因此用过的数字在下一层就不能再用了。
	// 强行传入返回结果指针，外加一个切片来存储当前遍历的路径结果集。
    result := [][]int{}
	if n <= 0 || k <= 0 || n < k {
		return result
	}
	res := make([]int, k)
	dfs(n, k, 1, res, &result)
	return result
}

func dfs(n, k, start int, res []int, result *[][]int) {
    l := len(res)
	for i := start; i <= n; i++ {
		if k == 1 {
			res[l-1] = i
			dst := make([]int, l)
			copy(dst, res)
			*result = append(*result, dst)
		} else {
			res[l-k] = i
			dfs(n, k-1, i+1, res, result)
		}
	}
}
// @lc code=end

