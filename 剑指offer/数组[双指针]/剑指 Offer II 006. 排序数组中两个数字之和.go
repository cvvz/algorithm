// https://leetcode-cn.com/problems/kLl5u1/

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i, j}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
}