/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	pathArr := strings.Split(path, "/")
	stack := new(stringStack)

	for _, subPath := range pathArr {
		switch subPath {
		case ".", "":
		case "..":
			stack.Pop()
		default:
			stack.Push(subPath)
		}
	}

	return "/" + strings.Join(*stack, "/")
}

// 用栈保存路径
type stringStack []string

func (s *stringStack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s *stringStack) Push(str string) {
	*s = append(*s, str)
}

func (s *stringStack) IsEmpty() bool {
	return len(*s) == 0
}

// @lc code=end

