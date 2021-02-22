/*
 * @lc app=leetcode.cn id=410 lang=golang
 *
 * [410] 分割数组的最大值
 */

// @lc code=start
func splitArray(nums []int, m int) int {
	// 二分查找
	sum, max := 0, 0
	for _, item := range nums {
		sum += item
		max = Max(max, item)
	}
	l, r := max, sum
	for l < r {
		mid := int(uint(l+r) >> 1)
		cnt, tmp := 1, 0
		for _, item := range nums {
			tmp += item
			if tmp > mid {
				cnt++
				tmp = item
			}
		}
		if cnt > m {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func Max(a int, bs ...int) int {
	for _, b := range bs {
		if a < b {
			a = b
		}
	}
	return a
}

// @lc code=end

