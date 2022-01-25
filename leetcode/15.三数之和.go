/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start, end := i+1, len(nums)-1
		for start < end {
			sum := nums[start] + nums[end] + nums[i]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[start], nums[end]})
				// 跳过重复数
				for start+1 <= len(nums)-1 && nums[start] == nums[start+1] {
					start++
				}
				start++
				// 跳过重复数
				for end-1 >= 0 && nums[end] == nums[end-1] {
					end--
				}
				end--
			} else if sum < 0 {
				start++
			} else {
				end--
			}
		}
	}
	return ans
}

// 排序 -> 两数之和（跳过重复数）

// @lc code=end

