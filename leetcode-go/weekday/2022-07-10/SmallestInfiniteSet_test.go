package _022_07_10

import (
	"fmt"
	"testing"
)

type SmallestInfiniteSet struct {
	list []int
}

func Constructor() SmallestInfiniteSet {
	list := make([]int, 1024)
	for i := 0; i < 1024; i++ {
		list[i] = i + 1
	}
	return SmallestInfiniteSet{list: list}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(this.list) > 0 {
		t := this.list[0]
		this.list = this.list[1:]
		return t
	}
	return 0
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	list := this.list

	l, r := 0, len(list)-1
	for {
		mid := (l + r) / 2
		if list[mid] > num {
			r = mid - 1
		}
		if list[mid] < num {
			l = mid + 1
		}
		if list[mid] == num {
			return
		}
		if l > r {
			break
		}
	}
	if l >= len(list) {
		l = len(list) - 1
	}
	list = append(list[:l+1], list[l:]...)
	list[l] = num
	this.list = list
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

func TestSmallestInfiniteSet(t *testing.T) {
	obj := Constructor()
	fmt.Println(obj.PopSmallest())
	fmt.Println(obj.PopSmallest())
	obj.AddBack(9999)
	fmt.Println(obj.PopSmallest())
	obj.AddBack(2)
	fmt.Println(obj.PopSmallest())
	fmt.Println(obj.PopSmallest())
}
