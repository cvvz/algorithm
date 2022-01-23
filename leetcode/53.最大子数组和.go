/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// 方法一：贪心
	// 一旦变成负数，就重新计算
	n := len(nums)
	max := -1000000
	sum := 0
	for i := 0; i < n; i++ {
		sum = nums[i] + sum
		if sum < 0 {
			if nums[i] > max {
				max = nums[i]
			}
			sum = 0
		} else {
			if sum > max {
				max = sum
			}
		}
	}

	return max
}

// 方法二：动态规划
// func maxSubArray(nums []int) int {
// 	sumSubArray := nums[0]
// 	ans := sumSubArray

// 	max := func(a, b int) int {
// 		if a > b {
// 			return a
// 		}

// 		return b
// 	}

// 	for i := 1; i < len(nums); i++ {
// 		sumSubArray = max(sumSubArray+nums[i], nums[i])
// 		if sumSubArray > ans {
// 			ans = sumSubArray
// 		}
// 	}

// 	return ans
// }

// @lc code=end

