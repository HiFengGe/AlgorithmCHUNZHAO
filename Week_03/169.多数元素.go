/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	length := len(nums)
	arr := make([]int, length)
	i := 0
	for j := 0; j < length; j++ {
		if i == 0 || arr[i-1] == nums[j] {
			arr[i] = nums[j]
			i++
		} else {
			i--
		}
	}
	return arr[i-1]
}

func majorityElement0(nums []int) int {
	length := len(nums)
	sort.Ints(nums)
	return nums[length/2]
}

// @lc code=end

