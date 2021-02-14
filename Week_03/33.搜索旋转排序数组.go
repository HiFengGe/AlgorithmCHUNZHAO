/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	arrMap := map[int]int{}
	for k, v := range nums {
		arrMap[v] = k
	}

	if result, ok := arrMap[target]; ok {
		return result
	}

	return -1

}

// @lc code=end

