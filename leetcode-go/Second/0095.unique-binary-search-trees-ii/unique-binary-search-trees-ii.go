package problem0095

import (
	"fmt"
	"github.com/w3xse7en/oj/leetcode-go/kit"
	"strconv"
	"strings"
)

type TreeNode = kit.TreeNode

var trees [][]int

func generateTrees(n int) []*TreeNode {
	var bsts []*TreeNode
	if n == 0 {
		return bsts
	}
	trees = [][]int{}
	list := []int{}
	for i := 1; i <= n; i++ {
		list = append(list, i)
	}
	buildAllSubList([]int{}, list)
	fmt.Println(trees)
	noDup := map[string]bool{}
	for _, tree := range trees {
		bst := buildBST(tree)
		ty := []string{}
		loop(bst, &ty)
		typetree := strings.Join(ty, ",")
		if !noDup[typetree] {
			noDup[typetree] = true
			bsts = append(bsts, bst)
		}
	}
	return bsts
}

func remove(list []int, k int) []int {
	newList := make([]int, len(list))
	copy(newList, list)
	return append(newList[:k], newList[k+1:]...)
}

func buildAllSubList(tree, list []int) {
	if len(list) == 0 {
		nt := make([]int, len(tree))
		copy(nt, tree)
		trees = append(trees, nt)
	}
	for k, v := range list {
		newList := remove(list, k)
		if len(tree) == 0 {
			buildAllSubList(append(tree, v), newList)
		} else {
			if v > tree[len(tree)-1] {
				buildAllSubList(append(tree, v), newList)
			} else {
				buildAllSubList(append(tree, v), newList)
			}
		}
	}
}

func buildBST(list []int) (root *TreeNode) {
	for k, v := range list {
		if k == 0 {
			root = &TreeNode{Val: v}
		} else {
			addNode(v, root)
		}
	}
	return root
}

func addNode(v int, root *TreeNode) {
	if root != nil {
		if v > root.Val {
			if root.Right == nil {
				root.Right = &TreeNode{Val: v}
			} else {
				addNode(v, root.Right)
			}
		} else {
			if root.Left == nil {
				root.Left = &TreeNode{Val: v}
			} else {
				addNode(v, root.Left)
			}
		}
	}
}

func loop(root *TreeNode, treeType *[]string) {
	if root != nil {
		*treeType = append(*treeType, strconv.Itoa(root.Val))
		loop(root.Left, treeType)
		loop(root.Right, treeType)
	} else {
		*treeType = append(*treeType, "null")
	}
}
