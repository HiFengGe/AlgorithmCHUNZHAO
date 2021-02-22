/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {

	var res string
	cnt := math.MaxInt32
	hashMap := make(map[byte]int)
	l := 0
	r := 0
	for i := 0; i < len(t); i++ {
		hashMap[t[i]]++
	}
	for r < len(s) {
		hashMap[s[r]]--
		for check(hashMap) {
			if r-l+1 < cnt {
				cnt = r - l + 1
				res = s[l : r+1]
			}
			hashMap[s[l]]++
			l++
		}
		r++
	}
	return res
}

func check(hashMap map[byte]int) bool {
	for _, v := range hashMap {
		if v > 0 {
			return false
		}
	}
	return true
}

// @lc code=end

