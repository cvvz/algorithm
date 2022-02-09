/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	numsCopy = nums

	pq := &PriorityQueue{
		IntSlice: make([]int, k),
	}
	for i := 0; i < k; i++ {
		pq.IntSlice[i] = i
	}
	heap.Init(pq)

	res := make([]int, 0)
	res = append(res, nums[pq.IntSlice[0]])
	for i, j := k, 0; i < n; i, j = i+1, j+1 {
		heap.Push(pq, i)
		// 如果最大值不在窗口内，一直吐到最大值在窗口内为止
		for pq.IntSlice[0] > i || pq.IntSlice[0] <= j {
			heap.Pop(pq)
		}

		res = append(res, nums[pq.IntSlice[0]])
	}
	return res
}

var numsCopy []int

type PriorityQueue struct {
	sort.IntSlice
}

// 优先级队列中存的是nums的下标
func (pq *PriorityQueue) Less(i, j int) bool {
	return numsCopy[pq.IntSlice[i]] > numsCopy[pq.IntSlice[j]]
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.IntSlice = append(pq.IntSlice, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(pq.IntSlice)
	if n == 0 {
		return nil
	}
	x := pq.IntSlice[n-1]
	pq.IntSlice = pq.IntSlice[:n-1]
	return x
}

// @lc code=end

