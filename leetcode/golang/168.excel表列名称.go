/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
func convertToTitle(columnNumber int) string {
	ans := ""
	for columnNumber > 26 {
		if columnNumber%26 == 0 {
			ans = fmt.Sprintf("%c%s", 'Z', ans)
			columnNumber = columnNumber/26 - 1
		} else {
			ans = fmt.Sprintf("%c%s", 'A'+columnNumber%26-1, ans)
			columnNumber /= 26
		}
	}
	return fmt.Sprintf("%c%s", 'A'+columnNumber-1, ans)
}

// 类似于26进制，但是考虑52这种特殊情况
// @lc code=end

