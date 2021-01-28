/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	//暴力
	// n := len(nums)
	// k %= n // 优化 k > len
	// for k > 0 {
	// 	temp := nums[n-1]
	// 	for i := n - 1; i > 0; i-- {
	// 		nums[i] = nums[i-1]
	// 	}
	// 	nums[0] = temp
	// 	k--
	// }

	// 反转
	n := len(nums)
	k %= n                // 优化 k > len
	reverse(nums, 0, n-1) //第一个，最后一个
	reverse(nums, 0, k-1) //k前面的
	reverse(nums, k, n-1) //k后面的

}

// reverse 反转
func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// @lc code=end

