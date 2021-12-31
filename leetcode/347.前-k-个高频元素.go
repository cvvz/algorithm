/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start

// 堆一定是通过数组实现，因为堆是一个完全二叉树
type minHeap []NumFrequece

type NumFrequece struct {
	num      int
	frequece int
}

// 🌟Pop弹出来的一定是slice的最后一个元素，因为Heap Pop前会Swap
func (h *minHeap) Pop() interface{} {
	top := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return top
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(NumFrequece))
}

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i].frequece < (*h)[j].frequece
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func topKFrequent(nums []int, k int) []int {
	numFrequence := map[int]int{}

	for _, num := range nums {
		if _, exist := numFrequence[num]; !exist {
			numFrequence[num] = 1
			continue
		}

		numFrequence[num]++
	}

	// 构造小顶堆
	// ❌var mp *minHeap
	mp := new(minHeap)
	for num, frequece := range numFrequence {
		heap.Push(mp, NumFrequece{
			num:      num,
			frequece: frequece,
		})
		if mp.Len() > k {
			_ = heap.Pop(mp)
		}
	}

	var ans []int
	for i := 0; i < k; i++ {
		ans = append(ans, (*mp)[i].num)
	}

	return ans
}

// @lc code=end

