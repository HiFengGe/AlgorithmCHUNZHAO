/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
// 最终结果
var result [][]int
func permute(nums []int) [][]int {
	var pathNums []int
	var used = make([]bool, len(nums))
    result = [][]int{} //清除初始化下
	backtrack(nums, pathNums, used)
	return result
}

// 回溯核心
// nums: 原始列表
// pathNums: 路径上的数字
// used: 是否访问过
func backtrack(nums, pathNums []int, used[]bool) {
	// 	回溯，顾名思义，要往回走，为啥往回走，因为没路可走（走到底了）
	//  怎么走？即选择
	//  怎么回来？即撤销选择
	//  什么时候结束？即退出条件
	// 结束条件：走完了，也就是路径上的数字总数等于原始列表总数

	
	if len(nums) == len(pathNums) {
		tmp := make([]int, len(nums))
		// 切片底层公用数据，所以要copy
		copy(tmp, pathNums)
		result = append(result, tmp)
		return
	}

	// 开始遍历原始数组的每个数字
	for i:=0; i<len(nums); i++ {
		// 检查是否访问过
		if !used[i] {
			// 没有访问过就选择它，然后标记成已访问过的
			used[i] = true
			// 做选择：将这个数字加入到路径的尾部，这里用数组模拟链表
			pathNums = append(pathNums, nums[i])
			backtrack(nums,pathNums,used)
			// 撤销刚才的选择，也就是恢复操作
			pathNums = pathNums[:len(pathNums) -1]
			// 标记成未使用
			used[i] = false
		}
	}
}
// @lc code=end

