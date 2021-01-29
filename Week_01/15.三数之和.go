/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	// a + b + c = 0 => a + b = -c
	// 1.外层循环：指针 i 遍历数组。
	// 1.0.先对数组排序，便于跳过重复元素，如果当前元素和前一个元素相同，跳过。
	// 2. 内层循环：用双指针，去寻找满足三数和等于0的项
	sort.Ints(nums)
	var results [][]int //返回结果
	for i := 0; i < len(nums)-2;i++{
		// 如果当前元素和前一个元素相同，跳过
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}

		// 左右双指针
		target, left, right := -nums[i],i+1,len(nums)-1
		for left < right {
			//两数之各
			sum := nums[left] + nums[right]
			if sum == target{
				results = append(results,[]int{nums[i],nums[left],nums[right]})
				// 夹逼处理 -->（加处理） ... <--(减处理)
				left++
				right--

				//判断 --> 相邻值是否一样  跳过处理
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				//判断 <-- 相邻值是否一样  跳过处理
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}else if sum > target{//保留往left right<--处理
				right--
			}else if sum < target{//保留往right left--> 处理
				left++
			}

		}
	}
	return results

}
// @lc code=end

