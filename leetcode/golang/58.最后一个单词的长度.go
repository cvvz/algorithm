/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	length, last := 0, 0
	for i, _ := range s {
		if s[i] != ' ' {
			length++
			last = length
		} else {
			length = 0
		}
	}
	return last
}

// @lc code=end

