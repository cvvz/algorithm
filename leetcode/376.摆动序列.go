/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	// 方法一：贪心
	n := len(nums)
	if n == 0 {
		return 0
	}

	length := 1

	prevDifference := 0
	for i := 1; i < n; i++ {
		difference := nums[i] - nums[i-1]
		if isWiggle(difference, prevDifference) {
			length++
			prevDifference = difference
		}
	}

	return length
}

func isWiggle(a, b int) bool {
	if a == 0 {
		return false
	}

	if a > 0 && b <= 0 {
		return true
	}

	if a < 0 && b >= 0 {
		return true
	}

	return false
}

// @lc code=end

