/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[j] + nums[k] + nums[i]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				// 跳过重复数
				for j+1 <= len(nums)-1 && nums[j] == nums[j+1] {
					j++
				}
				j++
				// 跳过重复数
				for k-1 >= 0 && nums[k] == nums[k-1] {
					k--
				}
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}

func sort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 排序 -> 两数之和（跳过重复数）

// @lc code=end

