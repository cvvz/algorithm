/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	ans := []string{}
	combination := []byte{}
	leftCount, rightCount := 0, 0

	var bt func()
	bt = func() {
		if leftCount >= n+1 || rightCount > leftCount {
			return
		}

		if len(combination) == 2*n {
			temp := make([]byte, 2*n, 2*n)
			copy(temp, combination)
			ans = append(ans, string(temp))
			return
		}

		combination = append(combination, ')')
		rightCount++
		bt()
		combination = combination[:len(combination)-1]
		rightCount--

		combination = append(combination, '(')
		leftCount++
		bt()
		combination = combination[:len(combination)-1]
		leftCount--
	}

	bt()

	return ans
}

// @lc code=end

