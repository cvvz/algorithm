// https://leetcode-cn.com/problems/FortPu/

type RandomizedSet struct {
	set    map[int]int
	values []int
	size   int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		set:    make(map[int]int),
		values: []int{},
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, exist := this.set[val]; exist {
		return false
	}
	this.set[val] = this.size
	this.values = append(this.values, val)
	this.size++
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, exist := this.set[val]; !exist {
		return false
	}
	this.set[this.values[this.size-1]] = this.set[val]
	this.values[this.set[val]] = this.values[this.size-1]
	delete(this.set, val)
	this.values = this.values[:this.size-1]
	this.size--
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.values[rand.Intn(this.size)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */