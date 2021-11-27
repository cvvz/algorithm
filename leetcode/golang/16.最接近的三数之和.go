/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var ans = nums[0] + nums[1] + nums[2]
	for i, num1 := range nums {
		j, k := i+1, len(nums)-1
		for j < k {
			num2, num3 := nums[j], nums[k]
			sum := num1 + num2 + num3
			if target-sum == 0 {
				return sum
			}
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-ans)) {
				ans = sum
			}
			// 移动
			if target < sum {
				for k-1 >= 0 && nums[k] == nums[k-1] {
					k--
				}
				k--
				continue
			}

			for j+1 < len(nums) && nums[j] == nums[j+1] {
				j++
			}
			j++
		}
	}

	return ans
}

// @lc code=end

