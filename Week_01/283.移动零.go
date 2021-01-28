/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	//非0直接打入开始
	//0直接设置0
	//边界条件是总个数据-非0个数据

	j := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	//0值处理
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}

// @lc code=end

