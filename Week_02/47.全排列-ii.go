/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
// 最终结果
var result [][]int
func permuteUnique(nums []int) [][]int {
	var path []int
	// 先排序，这样可以判断nums[i-1]与nums[i]是否相等
    sort.Ints(nums)
	var visit = make([]bool, len(nums))
    result = [][]int{} //清除初始化下
	backtrack(nums, path, visit)
	return result
}

// 回溯核心
// nums: 原始列表
// path: 路径上的数字
// used: 是否访问过
func backtrack(nums, path []int, visit[]bool) {
	// 	回溯，顾名思义，要往回走，为啥往回走，因为没路可走（走到底了）
	if len(nums) == len(path) {
		result = append(result, append([]int(nil), path...))
		return
	}

	// 开始遍历原始数组的每个数字
	for i:=0; i<len(nums); i++ {
		// 排除一种情况，visit[i-1]尚未完成调用，正在执行中
		// [1,1,2] 如果path(1,1 )中现在是这样的，那么第二个1与前一个值相同如果不加!visit[i-1]
		// 那么这种情况就直接跳过了
		if visit[i] || i > 0 && !visit[i-1] && nums[i] == nums[i-1] {
			continue
		}
		// 没有访问过就选择它，然后标记成已访问过的
		visit[i] = true
		// 做选择：将这个数字加入到路径的尾部，这里用数组模拟链表
		path = append(path, nums[i])
		backtrack(nums,path,visit)
		// 撤销刚才的选择，也就是恢复操作
		path = path[:len(path) -1]
		// 标记成未使用
		visit[i] = false
	}
}
// @lc code=end

