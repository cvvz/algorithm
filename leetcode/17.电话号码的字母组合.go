/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
var ans []string

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

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	ans = []string{}
	backtracking(digits, 0, "")

	return ans
}

func backtracking(digits string, i int, combination string) {
	if i == len(digits) {
		ans = append(ans, combination)
		return
	}

	for _, c := range numMap[digits[i]] {
		backtracking(digits, i+1, combination+string(c))
	}
}

// @lc code=end

