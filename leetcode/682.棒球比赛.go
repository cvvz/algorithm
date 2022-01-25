/*
 * @lc app=leetcode.cn id=682 lang=golang
 *
 * [682] 棒球比赛
 */

// @lc code=start
func calPoints(ops []string) int {
	var sum int
	var stack []int
	for _, val := range ops {
		switch val[0] {
		case '+':
			score := stack[len(stack)-1] + stack[len(stack)-2]
			stack = append(stack, score)
		case 'D':
			score := 2 * stack[len(stack)-1]
			stack = append(stack, score)
		case 'C':
			stack = stack[:len(stack)-1]
		default:
			num, _ := strconv.Atoi(val)
			stack = append(stack, num)
		}
	}
	for _, val := range stack {
		sum += val
	}
	return sum
}

// @lc code=end

