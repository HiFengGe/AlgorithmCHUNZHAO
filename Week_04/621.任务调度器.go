/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	tsk := make([]int, 26)
	for _, t := range tasks {
		tsk[t-'A']++
	}
	max, count := len(tasks)/(n+1)+1, 0
	for _, v := range tsk {
		if v > max {
			max = v
			count = 1
		} else if v == max {
			count++
		}
	}
	if count == 0 {
		return len(tasks)
	}
	cpu := (max-1)*(n+1) + count
	if cpu < len(tasks) {
		return len(tasks)
	}
	return cpu
}

// @lc code=end

