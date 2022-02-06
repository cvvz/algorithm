/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	abs := func(x int) int {
		if x < 0 {
			return -1 * x
		}
		return x
	}

	ans := nums[0] + nums[1] + nums[2]

	for i, num1 := range nums {
		start, end := i+1, len(nums)-1
		for start < end {
			num2, num3 := nums[start], nums[end]
			sum := num1 + num2 + num3
			if target-sum == 0 {
				return sum
			}
			if abs(target-sum) < abs(target-ans) {
				ans = sum
			}
			// 移动
			if target < sum {
				for end-1 >= 0 && nums[end] == nums[end-1] {
					end--
				}
				end--
				continue
			}

			for start+1 < len(nums) && nums[start] == nums[start+1] {
				start++
			}
			start++
		}
	}

	return ans
}

// @lc code=end

