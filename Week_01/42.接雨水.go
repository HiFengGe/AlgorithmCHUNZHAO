/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	// 双指针：实际是动态规划方法的深入分析后的答案
	var result, leftMax, rightMax int
	left, right := 0, n-1

	// 左右两指针向中间扫
	// 原则：低洼雨水的面积总是以低的一边为边界
	for left < right {
		// 以低的一边作为边界
		if height[left] < height[right] {
			// 当前点比左边的点低，即出现低洼，可以接雨水
			if height[left] < leftMax {
				result += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
			left++
		} else {
			if height[right] < rightMax {
				result += rightMax - height[right]
			} else {
				rightMax = height[right]
			}
			right--
		}
	}
	return result
}

// @lc code=end

