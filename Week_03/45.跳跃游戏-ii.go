/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	k := 0
	buf := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i]+i > k {
			k = nums[i] + i
			for j := i + 1; j <= k && j < len(nums); j++ {
				if buf[j] == 0 {
					buf[j] = buf[i] + 1
				} else if buf[j] > buf[i]+1 {
					buf[j] = buf[i] + 1
				}
			}
		}
	}
	return buf[len(buf)-1]

}

// @lc code=end

