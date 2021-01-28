/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + 1
			//小于9的前面的没有必要再处理了，直接返回
			return digits
		}
		//9处理 进位
		digits[i] = 0
	}

	//首位为0  => 9 ,99
	//类型值和返回类型一样，长度比传参长度+1，初始化默认都是0 ，然后把首位设置为1就可以
	var newDigits = make([]int, n+1)
	newDigits[0] = 1
	return newDigits
}

// @lc code=end

