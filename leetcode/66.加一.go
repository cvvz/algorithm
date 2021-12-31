/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			break
		}
	}

	if digits[0] == 0 {
		digits = append(digits, 0)
		digits[0] = 1
	}

	return digits
}

// 全9的情况不要处理的太复杂

// 空间复杂度：O(1)
// 时间复杂度：O(n)

// @lc code=end

