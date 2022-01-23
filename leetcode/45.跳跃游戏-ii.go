/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	step := 0
	n := len(nums)
	maxIndex := 0
	endIndex := 0

	for i := 0; i < n-1; i++ {
		if i+nums[i] > maxIndex {
			maxIndex = i + nums[i]
		}
		if i == endIndex {
			endIndex = maxIndex
			step++
		}
	}

	return step
}

// @lc code=end

