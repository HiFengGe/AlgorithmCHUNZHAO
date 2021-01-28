/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	maxLen := len(nums)
	if maxLen < 2 {
		return maxLen
	}
	var count int = 1
	for k := range nums {
		next := k + 1
		if next != maxLen && nums[k] != nums[next] {
			nums[count] = nums[next]
			count++
		}
	}
	return count
}

// @lc code=end

