/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	lettersCode := map[byte]int64{}

	for i := 0; i <= 25; i++ {
		lettersCode[(byte)('a'+i)] = 1 << i
	}

	codeStrs := map[int64][]string{}

	for i := 0; i < len(strs); i++ {
		str := strs[i]
		var code int64 = 0
		for j := 0; j < len(str); j++ {
			code += lettersCode[str[j]]
		}
		codeStrs[code] = append(codeStrs[code], str)
	}

	ans := [][]string{}

	for _, strs := range codeStrs {
		ans = append(ans, strs)
	}

	return ans
}

// @lc code=end

