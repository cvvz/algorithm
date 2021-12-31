/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	stack := []byte{}
	match := map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
	for i, _ := range s {
		switch s[i] {
		case '(', '{', '[':
			stack = append(stack, s[i])
		default:
			if len(stack) == 0 {
				return false
			}
			if match[s[i]] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}

	return true
}

// 左括号入栈，右括号出栈

// 时间复杂度：O(n)
// 空间复杂度：O(n)
// @lc code=end

