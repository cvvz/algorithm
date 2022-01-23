/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] 压缩字符串
 */

// @lc code=start
func compress(chars []byte) int {
	i := 0
	prev := chars[0]
	count := 1

	n := len(chars)
	for j := 1; j < n; j++ {
		current := chars[j]

		if current != prev {
			prev = current
			i++
			if count > 1 {
				setAsString(chars, &i, count)
				count = 1
			}
			chars[i] = current
		} else {
			count++
		}
	}

	i++
	if count > 1 {
		setAsString(chars, &i, count)
	}

	return i
}

func setAsString(chars []byte, i *int, count int) {
	str := []byte(strconv.Itoa(count))
	for _, c := range str {
		chars[*i] = c
		*i++
	}
}

// @lc code=end

