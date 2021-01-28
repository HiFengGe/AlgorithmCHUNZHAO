/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	// 1 首先看字符串长度是否一样，不一样则为false
	// 2 看每个字符出现的频率是否一样，可以用到hash表。申请一个26长度的int数组。首先遍历字符串s然后，将每个字符串换算成索引后存入数组，并同时进行计数
	// 3 遍历字符串t，然后依次对字符对对应的数组值减一，如果出现小于0的情况则说明不匹配

	// if len([]rune(s)) != len([]rune(t)) {
	// 	return false
	// }

	// rune 等同于int32,处理unicode或utf-8字符
	hash := map[rune]int{}
	//++处理
	for _, v := range s {
		hash[v]++
	}

	//--处理
	for _, v := range t {
		if _, ok := hash[v]; !ok {
			return false
		}

		//-- 及unicode或utf-8字符处理
		if hash[v] == 1 {
			delete(hash, v)
		} else {
			hash[v]--
		}
	}

	return len(hash) == 0
}

// @lc code=end

