/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	ans := [][]int{}

	sort.Ints(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 两数之和
			twoSum := target - nums[i] - nums[j]
			begin, end := j+1, n-1
			for begin < end {
				if begin > j+1 && nums[begin] == nums[begin-1] {
					begin++
					continue
				}
				if end < n-1 && nums[end] == nums[end+1] {
					end--
					continue
				}

				sum := nums[begin] + nums[end]
				if sum == twoSum {
					ans = append(ans, []int{nums[i], nums[j], nums[begin], nums[end]})
					begin++
					end--
				} else if sum < twoSum {
					begin++
				} else {
					end--
				}
			}
		}
	}

	return ans
}

// @lc code=end

