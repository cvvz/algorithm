/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	n := len(nums)
	longest := 0

	for i := 0; i < n; i++ {
		if longest >= i && i+nums[i] > longest {
			longest = i + nums[i]
		}
		if longest >= n-1 {
			return true
		}
	}

	return false
}

// @lc code=end

