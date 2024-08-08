package _380_RandomizedSet

import (
	"math/rand"
)

type RandomizedSet struct {
	hash map[int]int
	list []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{map[int]int{}, []int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hash[val]; ok {
		return false
	}
	this.hash[val] = len(this.list)
	this.list = append(this.list, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.hash[val]
	if !ok {
		return false
	}
	last := len(this.list) - 1
	tail := this.list[last]
	this.list[idx] = tail
	this.hash[tail] = idx
	this.list = append(this.list[:last])
	delete(this.hash, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
