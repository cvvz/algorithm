/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}

	tmp := 0

	for x > tmp {
		tmp = tmp*10 + x%10
		if x == tmp {
			break
		}
		x /= 10
	}

	if x == tmp {
		return true
	}

	return false
}

// 1. 全部反转可能遇到溢出——反转一半
// 2. 反转结束条件：大于原数
// 3. 考虑处于中间的数字

// 时间复杂度：O(logn)
// 空间复杂度：O(1)
// @lc code=end

