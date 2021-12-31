// https://leetcode-cn.com/problems/2VG8Kg/

func minSubArrayLen(target int, nums []int) int {
	i, j := 0, 0
	sum := nums[0]
	ans := len(nums)

	notExist := true
	for j < len(nums) && i <= j {
		if sum >= target {
			if j-i+1 < ans {
				ans = j - i + 1
			}
			sum -= nums[i]
			i++
			notExist = false
		} else {
			j++
			if j < len(nums) {
				sum += nums[j]
			}
		}
	}

	if notExist {
		return 0
	}

	return ans
}