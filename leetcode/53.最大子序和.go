/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	numsLen := len(nums)
	pre, max := nums[0], nums[0]
	for i := 1; i < numsLen; i++ {
		if pre+nums[i] > nums[i] {
			pre += nums[i]
		} else {
			pre = nums[i]
		}

		if max < pre {
			max = pre
		}
	}
	return max
}

// 动态规划
// f(i)：以i结尾的最大和
// f(i) = max{f(i-1) + nums[i], nums[i]}

// 时间复杂度：O(n)
// 空间复杂度：O(1)
// @lc code=end

