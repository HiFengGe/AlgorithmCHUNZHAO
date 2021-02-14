/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {

	// for row := 0; row < len(matrix); row++ {
	// 	if matrix[row][0] >= target && target <= matrix[row][len(matrix[row])-1] {
	// 		for col := 0; col < len(matrix[row]); col++ {
	// 			if target == matrix[row][col] {
	// 				return true
	// 			}
	// 		}
	// 	}
	// }
	// return false

	row := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] >= target
	})

	if row < len(matrix) && matrix[row][0] == target {
		return true
	}

	row--
	if row < 0 {
		return false
	}

	//search for column
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}

// @lc code=end

