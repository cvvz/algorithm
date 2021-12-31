// https://leetcode-cn.com/problems/1fGaJU/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var ans [][]int

	for first, _ := range nums {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		third := len(nums) - 1
		second := first + 1

		for third > second {
			if second > first+1 && nums[second] == nums[second-1] {
				second++
				continue
			}
			if third < len(nums)-1 && nums[third] == nums[third+1] {
				third--
				continue
			}

			sum := nums[first] + nums[second] + nums[third]
			if sum == 0 {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
				second++
				third--
			} else if sum < 0 {
				second++
			} else {
				third--
			}
		}
	}

	return ans
}