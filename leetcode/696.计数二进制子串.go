/*
 * @lc app=leetcode.cn id=696 lang=golang
 *
 * [696] 计数二进制子串
 */

// @lc code=start
func countBinarySubstrings(s string) int {
	n := len(s)

	counter := []int{}
	count := 1
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			count++
		} else {
			counter = append(counter, count)
			count = 1
		}
	}
	counter = append(counter, count)

	n = len(counter)
	ans := 0
	for i := 0; i < n-1; i++ {
		ans += min(counter[i], counter[i+1])
	}

	return ans
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

// @lc code=end

