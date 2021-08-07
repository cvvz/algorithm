/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start

func intToRoman(num int) string {
	var dic = map[int]map[int]string{
		0: { // 个位
			1: "I",
			4: "IV",
			5: "V",
			9: "IX",
		},
		1: { //十位
			1: "X",
			4: "XL",
			5: "L",
			9: "XC",
		},
		2: { //百位
			1: "C",
			4: "CD",
			5: "D",
			9: "CM",
		},
		3: { //千位
			1: "M",
		},
	}

	ans := ""

	for digit := 0; num != 0; digit++ {
		number := num % 10
		switch {
		case number < 4:
			for i := 0; i < number; i++ {
				ans = dic[digit][1] + ans
			}
		case number > 4 && number < 9:
			for i := 0; i < number-5; i++ {
				ans = dic[digit][1] + ans
			}
			ans = dic[digit][5] + ans
		default:
			ans = dic[digit][number] + ans
		}
		num = num / 10
	}

	return ans
}

// 拆分成个十百千四个部分
// 每个部分和4、9做对比；是否在4、9之间，是的话先减5

// @lc code=end

