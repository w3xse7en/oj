package _022_07_31

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestClosestMeetingNode(t *testing.T) {
	fmt.Println(closestMeetingNode([]int{2, 3, 1, 0}, 0, 1))
	fmt.Println(closestMeetingNode([]int{2, 2, 3, -1}, 0, 1))
	fmt.Println(closestMeetingNode([]int{1, 2, -1}, 0, 2))
}

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n1 := map[int]int{}
	n2 := map[int]int{}
	tmp := make([]int, len(edges))
	n1 = loop(edges, node1, 0, tmp, n1)
	tmp = make([]int, len(edges))
	n2 = loop(edges, node2, 0, tmp, n2)
	mmax := math.MaxInt32
	matchNode := []int{}
	for i1, node := range n1 {
		if _, ok := n2[node]; !ok {
			continue
		}
		d1 := i1
		d2 := n2[node]
		max := 0
		if d2 > d1 {
			max = d2
		} else {
			max = d1
		}
		if max < mmax {
			matchNode = []int{node}
			mmax = max
		} else if max == mmax {
			matchNode = append(matchNode, node)
		}
	}
	if len(matchNode) == 0 {
		return -1
	}
	sort.Ints(matchNode)
	return matchNode[0]
}
func loop(edges []int, node, deep int, tmp []int, index map[int]int) map[int]int {
	if node == -1 {
		return index
	}
	if tmp[node] >= 1 {
		return index
	}
	tmp[node] += 1
	index[node] = deep
	return loop(edges, edges[node], deep+1, tmp, index)
}
