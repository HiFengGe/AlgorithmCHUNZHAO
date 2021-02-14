/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	subMaps := [][]int{[]int{}}
	for _, num := range nums {
		for _, subset := range subMaps {
			newSubset := append([]int{}, subset...)
			newSubset = append(newSubset, num)
			subMaps = append(subMaps, newSubset)
		}
	}
	return subMaps
}

// @lc code=end

