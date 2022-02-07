/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	// 面对10美元，手上必须要有5美元
	// 面对20美元，手上必须要有10 + 5 或者 3个5
	fives, tens := 0, 0

	for _, bill := range bills {
		switch bill {
		case 5:
			fives++
		case 10:
			if fives == 0 {
				return false
			}
			fives--
			tens++
		case 20:
			if tens > 0 && fives > 0 {
				tens--
				fives--
			} else if fives >= 3 {
				fives -= 3
			} else {
				return false
			}
		}
	}
	return true
}

// @lc code=end

