/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start

// 首先做个位数 * 字符串 ，存到string中，末尾加“0”
// 再做字符串的加法
func multiply(num1 string, num2 string) string {
	temp := make([]string, len(num2), len(num2))

	for i := 0; i < len(num2); i++ {
		zeroNum := len(num2) - i - 1
		zeroString := ""

		for j := 0; j < zeroNum; j++ {
			zeroString += "0"
		}

		temp[i] = singleMultiply(num2[i], num1)
		if strings.Compare(temp[i], "0") != 0 {
			temp[i] += zeroString
		}
	}

	result := "0"

	for i := 0; i < len(temp); i++ {
		result = stringAdd(result, temp[i])
	}

	return result
}

func singleMultiply(num1 byte, num2 string) string {
	if num1 == '0' || strings.Compare(num2, "0") == 0 {
		return "0"
	}
	carry := 0
	result := ""
	i := len(num2) - 1
	for i >= 0 {
		singleResult := string(((int)((num2[i]-'0')*(num1-'0'))+carry)%10 + '0')
		carry = ((int)((num2[i]-'0')*(num1-'0')) + carry) / 10
		result = singleResult + result
		i--
	}

	if carry >= 1 {
		result = fmt.Sprint(carry) + result
	}

	return result
}

func stringAdd(num1, num2 string) string {
	if strings.Compare(num1, "0") == 0 {
		return num2
	}
	if strings.Compare(num2, "0") == 0 {
		return num1
	}

	result := ""
	carry := 0
	i := len(num1) - 1
	j := len(num2) - 1

	for i >= 0 || j >= 0 {
		if j < 0 {
			singleResult := ((int)(num1[i]-'0')+carry)%10 + '0'
			result = (string)(singleResult) + result
			carry = ((int)(num1[i]-'0') + carry) / 10
			i--
		} else if i < 0 {
			singleResult := ((int)(num2[j]-'0')+carry)%10 + '0'
			result = (string)(singleResult) + result
			carry = ((int)(num2[j]-'0') + carry) / 10
			j--
		} else {
			singleResult := ((int)(num1[i]-'0'+num2[j]-'0')+carry)%10 + '0'
			result = (string)(singleResult) + result
			carry = ((int)(num1[i]-'0'+num2[j]-'0') + carry) / 10
			i--
			j--
		}
	}

	if carry == 1 {
		return "1" + result
	}

	return result
}

// @lc code=end

