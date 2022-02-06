/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	result := ""
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += (int)(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += (int)(num2[j] - '0')
			j--
		}
		result = strconv.Itoa(sum%10) + result
		carry = sum / 10
	}

	if carry > 0 {
		return "1" + result
	}

	return result
}

// @lc code=end

