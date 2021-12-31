/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var ans string
	buff := bytes.NewBufferString(ans)

	z := [][]byte{}
	for i := 0; i < numRows; i++ {
		z = append(z, []byte{})
	}
	down := true
	currentRow := 0

	for i := 0; i < len(s); i++ {
		z[currentRow] = append(z[currentRow], s[i])
		if currentRow == numRows-1 {
			currentRow--
			down = false
		} else if currentRow == 0 {
			currentRow++
			down = true
		} else {
			if down {
				currentRow++
			} else {
				currentRow--
			}
		}
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < len(z[i]); j++ {
			buff.WriteByte(z[i][j])
		}
	}

	return buff.String()
}

// @lc code=end

