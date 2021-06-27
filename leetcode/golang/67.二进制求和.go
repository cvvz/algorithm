/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	aLen, bLen := len(a), len(b)
	ans := ""
	carry := 0

	for i, j := aLen-1, bLen-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		if i >= 0 {
			carry += int(a[i] - '0')
		}
		if j >= 0 {
			carry += int(b[j] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry != 0 {
		return "1" + ans
	}
	return ans
}

// 注意多循环变量的for循环写法

// 时间复杂度：O(n)
// 空间复杂度：O(1)
// @lc code=end

