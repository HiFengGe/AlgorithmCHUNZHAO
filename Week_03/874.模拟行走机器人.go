/*
 * @lc app=leetcode.cn id=874 lang=golang
 *
 * [874] 模拟行走机器人
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	obstaclesMap := make(map[[2]int]bool, len(obstacles))
	for k := range obstacles {
		obstaclesMap[[2]int{
			obstacles[k][0],
			obstacles[k][1],
		}] = true
	}

	directions :=
		[4][2]int{
			{0, 1},  // 北
			{1, 0},  // 东
			{0, -1}, // 南
			{-1, 0}, // 西
		}
	var curDirection, x, y int
	var result int
	for _, command := range commands {
		switch command {
		case -2:
			curDirection = (curDirection + 3) % 4
		case -1:
			curDirection = (curDirection + 1) % 4
		default:
			for i := 0; i < command; i++ {
				x += directions[curDirection][0]
				y += directions[curDirection][1]
				if _, ok := obstaclesMap[[2]int{x, y}]; ok {
					x -= directions[curDirection][0]
					y -= directions[curDirection][1]
				}
				result = max(result, x*x+y*y)
			}
		}
	}
	return result
}

func max(result int, i int) int {
	if result < i {
		return i
	}
	return result
}

// @lc code=end

