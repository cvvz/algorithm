/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	start := false
	flag := 1
	ret := 0
	//找到第一个合法字符
	for _, b := range s {
		if !start {
			switch {
			case b == ' ':
				continue
			case b >= '0' && b <= '9':
				start = true
			case b == '-':
				start = true
				flag = -1
				continue
			case b == '+':
				start = true
				continue
			default:
				return 0
			}
		}

		if b >= '0' && b <= '9' {
			ret = ret*10 + int(b-'0')
			if ret*flag > 2147483647 {
				return 2147483647
			}
			if ret*flag < -2147483648 {
				return -2147483648
			}
		} else {
			return ret * flag
		}
	}
	return ret * flag
}

// @lc code=end

