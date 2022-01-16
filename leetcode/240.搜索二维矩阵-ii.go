/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	// 解：逐行二分查找
	// for i := 0; i < len(matrix); i++ {
	// 	if binSearch(matrix[i], 0, len(matrix[i])-1, target) {
	// 		return true
	// 	}
	// }

	// return false

	// 最优解：从右上角开始查找
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		current := matrix[i][j]
		if current == target {
			return true
		} else if current < target {
			i++
		} else {
			j--
		}
	}

	return false
}

func binSearch(arr []int, start, end, target int) bool {
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return false
}

// @lc code=end

