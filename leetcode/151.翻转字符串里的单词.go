/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
func reverseWords(s string) string {
	ss := []byte(s)
	n := len(s)

	reverse(ss)

	i, j := 0, 0
	for j < n {
		if ss[j] != ' ' {
			j++
		} else {
			reverse(ss[i:j])
			j++
			i = j
		}
	}
	reverse(ss[i:j])
	return string(ss)
}

func reverse(s []byte) {
	n := len(s)

	for start, end := 0, n-1; start < end; start, end = start+1, end-1 {
		s[start], s[end] = s[end], s[start]
	}
}

// @lc code=end

