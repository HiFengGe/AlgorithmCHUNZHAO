/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	result := make(map[int]int, 0)
	res := make([]int, 0)
	for _, i := range nums {
		if _, ok := result[i]; ok {
			result[i]++
		} else {
			result[i] = 1
			res = append(res, i)
		}
	}
	// fmt.Println("result:", result)
	sort.Slice(res, func(i, j int) bool {
		return result[res[i]] > result[res[j]]
	})
	return res[:k]
}
// @lc code=end

