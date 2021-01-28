/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {

	keyMaps := map[[26]int][]string{}

	for _, s := range strs {
		k := [26]int{}
		//非空处理
		for i := 0; i < len(s); i++ {
			// ASCII码对照 a在ASCII码中是97 b为98 ...
			k[s[i]-'a'] += 1
		}
		keyMaps[k] = append(keyMaps[k], s)
	}

	result := [][]string{}
	for _, v := range keyMaps {
		result = append(result, v)
	}
	return result

}

// @lc code=end

