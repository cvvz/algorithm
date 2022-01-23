/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var numMap = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	n := len(digits)

	ans := []string{}
	subAns := []byte{}

	var backtracking func(int)
	backtracking = func(i int) {
		if i == n {
			ans = append(ans, string(subAns))
			return
		}

		letters := []byte(numMap[digits[i]])
		for _, letter := range letters {
			subAns = append(subAns, letter)
			backtracking(i + 1)
			subAns = subAns[:len(subAns)-1]
		}
	}

	backtracking(0)

	return ans
}

// @lc code=end

