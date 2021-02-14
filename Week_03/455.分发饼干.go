/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	//排序
	sort.Ints(g)
	sort.Ints(s) //尺寸

	i := 0 //孩子
	j := 0 //尺寸
	n := 0 //最后数
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
			n++
		}
		j++
	}
	return n
}

// @lc code=end

