/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 */

// @lc code=start
func calculate(s string) int {
	var numStack []int
	var opStack []byte
	sLen := len(s)
	var val int
	for i := 0; i < sLen; i++ {
		switch s[i] {
		case '+':
			opStack = append(opStack, '+')
		case '-':
			opStack = append(opStack, '-')
		case '(':
			opStack = append(opStack, '(')
		case ')':
			if opStack[len(opStack)-1] == '(' {
				opStack = opStack[:len(opStack)-1]
				break
			}
			j := len(opStack) - 1
			k := len(numStack) - 1
			for opStack[j] != '(' {
				j--
				k--
			}
			tmp := numStack[k]
			jj := j + 1
			kk := k + 1
			for jj <= len(opStack)-1 {
				if opStack[jj] == '+' {
					tmp += numStack[kk]
				} else {
					tmp -= numStack[kk]
				}
				kk++
				jj++
			}
			numStack[k] = tmp
			numStack = numStack[:k+1]
			opStack = opStack[:j]

		case ' ':
		default:
			val = val*10 + int(s[i]-'0')
			if i+1 < sLen && s[i+1] >= '0' && s[i+1] <= '9' {
				break
			}
			numStack = append(numStack, val)
			val = 0
		}
	}
	val = numStack[0]
	for i := 0; i < len(opStack); i++ {
		if opStack[i] == '+' {
			val += numStack[i+1]
		} else {
			val -= numStack[i+1]
		}
	}
	return val
}

// @lc code=end

