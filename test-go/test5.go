package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	loop([]*TreeNode{root}, 0)
	return max
}

var max int

func loop(roots []*TreeNode, y int) {
	y++
	XNodes := make([]*TreeNode, 0)

	exist := false
	cnt := 0
	for _, v := range roots {
		if v != nil {
			XNodes = append(XNodes, v.Left)
			XNodes = append(XNodes, v.Right)
			if v.Left != nil {
				cnt++
				exist = true
			}
			if v.Right != nil {
				cnt++
				exist = true
			}
		} else {
			XNodes = append(XNodes, nil)
			XNodes = append(XNodes, nil)
		}
	}

	if exist {
		if cnt >= 2 {
			if len(XNodes) > max {
				max = len(XNodes)
			}
		}
		loop(XNodes, y)
	}
}
func ca(x, y int) int {
	fmt.Println(x, y, x+y)
	return x + y
}

func main() {
	wa := sync.WaitGroup{}
	wa.Add(1)
	go func() {
		for {
			fmt.Println(1)
		}
	}()
	wa.Wait()

	// treeNode := &TreeNode{Val: 0}
	// treeNode1 := &TreeNode{Val: 1}
	// treeNode2 := &TreeNode{Val: 2}
	// treeNode3 := &TreeNode{Val: 3}
	// treeNode4 := &TreeNode{Val: 4}
	// treeNode5 := &TreeNode{Val: 5}
	// treeNode6 := &TreeNode{Val: 6}
	// treeNode.Left = treeNode1
	// treeNode.Right = treeNode2
	// treeNode2.Right = treeNode4
	// treeNode1.Left = treeNode3
	// treeNode3.Left = treeNode5
	// treeNode4.Right = treeNode6
	// treeNode3.Right = treeNode4
	// fmt.Println(widthOfBinaryTree(treeNode))
}
