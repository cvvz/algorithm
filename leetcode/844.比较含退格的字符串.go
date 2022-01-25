/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
// func backspaceCompare(S string, T string) bool {
// 	var s, t []byte
// 	s_len := len(S)
// 	t_len := len(T)

// 	i, j := 0, 0
// 	for i < s_len || j < t_len {
// 		if i < s_len {
// 			if S[i] != '#' {
// 				s = append(s, S[i])
// 			} else {
// 				if len(s) > 0 {
// 					s = s[:len(s)-1]
// 				}
// 			}
// 			i++
// 		}
// 		if j < t_len {
// 			if T[j] != '#' {
// 				t = append(t, T[j])
// 			} else {
// 				if len(t) > 0 {
// 					t = t[:len(t)-1]
// 				}
// 			}
// 			j++
// 		}
// 	}
// 	s_len = len(s)
// 	t_len = len(t)
// 	if s_len != t_len {
// 		return false
// 	}
// 	for i := 0; i < s_len; i++ {
// 		if s[i] != t[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
func backspaceCompare(S string, T string) bool {
	s_len, t_len := len(S), len(T)
	s_back, t_back := 0, 0
	i, j := s_len-1, t_len-1
	for i >= 0 || j >= 0 {
		if i >= 0 {
			if S[i] == '#' {
				s_back++
				i--
				continue
			} else {
				if s_back > 0 {
					i--
					s_back--
					continue
				}
			}
		}
		if j >= 0 {
			if T[j] == '#' {
				t_back++
				j--
				continue
			} else {
				if t_back > 0 {
					j--
					t_back--
					continue
				}
			}
		}
		if i < 0 || j < 0 {
			return false
		} else {
			if S[i] != T[j] {
				return false
			} else {
				i--
				j--
			}
		}
	}
	return true
}

// @lc code=end

