/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	result := 0
	for x != 0 {
		digit := x % 10
		x /= 10

		if (math.MaxInt32-digit)/10 < result ||
			(math.MinInt32+digit)/10 > result {
			return 0
		}

		result = result*10 + digit
	}
	return result
}

// -2147483648 ～ 2147483647
// 做加法或者乘法会造成溢出错误，所以需要做减法和除法
// 负数取模仍然等于负数！

//时间复杂度：O(log∣x∣)
//空间复杂度：O(1)

// @lc code=end

