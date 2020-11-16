package problem0146

import (
	"fmt"
	"testing"
)

func TestLRUCache_search(t *testing.T) {
	//fmt.Println(searchv2([]int{4, 3, 2, 1}, 1))
	//fmt.Println(searchv2([]int{4, 3, 2, 1}, 2))
	//fmt.Println(searchv2([]int{4, 3, 2, 1}, 3))
	//fmt.Println(searchv2([]int{4, 3, 2, 1}, 4))
	//fmt.Println(searchv2([]int{4, 3, 2, 1}, -1))
	//fmt.Println(searchv2([]int{1, 2, 3, 4}, 1))
	//fmt.Println(searchv2([]int{1, 2, 3, 4}, 2))
	//fmt.Println(searchv2([]int{1, 2, 3, 4}, 3))
	//fmt.Println(searchv2([]int{1, 2, 3, 4}, 4))
	//fmt.Println(searchv2([]int{1, 2, 3, 4}, 5))
	//fmt.Println(searchv2([]int{1, 2, 3, 4}, -1))
}
func TestLRUCache_All(t *testing.T) {
	lruCache := Constructor(1)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	fmt.Println(lruCache.Get(1))
	lruCache.Put(3, 3)
	fmt.Println(lruCache.Get(2))
	lruCache.Put(4, 4)
	fmt.Println(lruCache.Get(1))
	fmt.Println(lruCache.Get(3))
	fmt.Println(lruCache.Get(4))
}
