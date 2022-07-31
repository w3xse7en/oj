package _022_07_03

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	var matrix = make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = -1
		}
	}
	x, y := 0, 0

	top, bottom, left, right := 0, m-1, 0, n-1
	xp, xm, yp, ym := true, false, false, false

	node := head
	for node != nil {
		if x > right {
			x--
			y++
			top++
			xp, xm, yp, ym = false, false, true, false
		}
		if x < left {
			x++
			y--
			bottom--
			xp, xm, yp, ym = false, false, false, true
		}
		if y > bottom {
			y--
			x--
			right--
			xp, xm, yp, ym = false, true, false, false
		}
		if y < top {
			y++
			x++
			left++
			xp, xm, yp, ym = true, false, false, false
		}
		matrix[y][x] = node.Val
		if xp {
			x++
		}
		if xm {
			x--
		}
		if yp {
			y++
		}
		if ym {
			y--
		}
		node = node.Next
	}
	return matrix
}

func loop(node *ListNode) {
	if node != nil {
		loop(node.Next)
	}
}

func buildNode(list []int) *ListNode {
	head := &ListNode{}
	tail := head
	for i, v := range list {
		if i == 0 {
			head.Val = v
		} else {
			node := &ListNode{}
			node.Val = v
			tail.Next = node
			tail = node
		}
	}
	return head
}

func TestSpiralMatrix(t *testing.T) {
	fmt.Println(spiralMatrix(3, 5, buildNode([]int{3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0})))
	fmt.Println(spiralMatrix(1, 4, buildNode([]int{0, 1, 2})))
}
