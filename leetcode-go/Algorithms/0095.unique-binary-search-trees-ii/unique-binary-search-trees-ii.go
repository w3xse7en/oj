package problem0095

import (
	"github.com/w3xse7en/oj/leetcode-go/kit"
)

type TreeNode = kit.TreeNode

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return RecGenTrees(1, n)
}

func RecGenTrees(s, e int) []*TreeNode {
	if e < s {
		return []*TreeNode{nil}
	}
	result := []*TreeNode{}
	for i := s; i <= e; i++ {
		lefts := RecGenTrees(s, i-1)
		rights := RecGenTrees(i+1, e)
		for _, l := range lefts {
			for _, r := range rights {
				result = append(result, &TreeNode{
					Val:   i,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return result
}
