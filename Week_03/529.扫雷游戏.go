/*
 * @lc app=leetcode.cn id=529 lang=golang
 *
 * [529] 扫雷游戏
 */

// @lc code=start
func updateBoard(board [][]byte, click []int) [][]byte {
	if len(board) == 0 || len(board[0]) == 0 {
		return board
	}

	rows, cols := len(board), len(board[0])
	// 检查坐标函数
	var check func(int, int) bool
	check = func(i int, j int) bool {
		return 0 <= i && i < rows && 0 <= j && j < cols
	}

	if !check(click[0], click[1]) {
		return board
	}

	// 如果是雷，则标记为X，退出
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	// 处理为E的情况
	var dfs func(int, int)
	dfs = func(i int, j int) {
		if !(check(i, j) && board[i][j] == 'E') {
			return
		}
		d := [][]int{{0, 1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {1, 0}, {1, -1}, {1, 1}}
		mineCount := 0
		for k := 0; k < len(d); k++ {
			if check(i+d[k][0], j+d[k][1]) && board[i+d[k][0]][j+d[k][1]] == 'M' {
				mineCount++
			}
		}

		if mineCount > 0 {
			board[i][j] = byte(mineCount + '0')
		} else {
			board[i][j] = 'B'
			for k := 0; k < len(d); k++ {
				dfs(i+d[k][0], j+d[k][1])
			}
		}
		return
	}
	dfs(click[0], click[1])
	return board
}

// @lc code=end

