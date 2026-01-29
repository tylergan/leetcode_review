package binary_tree

import (
	"strconv"
	"strings"
)

/*
Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

A subtree of a binary tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.

Example 1:

Input: root = [1,2,3,4,5], subRoot = [2,4,5]

Output: true
Example 2:

Input: root = [1,2,3,4,5,null,null,6], subRoot = [2,4,5]

Output: false
Constraints:

1 <= The number of nodes in both trees <= 100.
-100 <= root.val, subRoot.val <= 100
*/

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var sameBinaryTree func(*TreeNode, *TreeNode) bool
	sameBinaryTree = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		} else if p == nil || q == nil || p.Val != q.Val {
			return false
		}
		return sameBinaryTree(p.Left, q.Left) && sameBinaryTree(p.Right, q.Right)
	}
	if root == nil && subRoot == nil {
		return true
	} else if root == nil || subRoot == nil {
		return false
	} else if sameBinaryTree(root, subRoot) { // if sameBinaryTree fails, possible that there still exists a subtree e.g., root = [1, 1], sub-root = [1]
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

/*
SERIALISATION APPROACH
*/

func serialisation(node *TreeNode) string {
	if node == nil {
		return "$#"
	}
	return "$" + strconv.Itoa(node.Val) + serialisation(node.Left) + serialisation(node.Right)
}

func isSubtreeSerialisation(root *TreeNode, subRoot *TreeNode) bool {
	return strings.Contains(serialisation(root), serialisation(subRoot))
}

func isSubtreeKMP(root *TreeNode, subRoot *TreeNode) bool {
	kmpSearch := func(text string, pattern string) bool {
		buildLPS := func(pattern string) []int {
			m := len(pattern)
			lps := make([]int, m)
			lps[0] = 0

			length := 0
			i := 1

			for i < m {
				if pattern[i] == pattern[length] {
					length++
					lps[i] = length
					i++
				} else {
					if length != 0 {
						length = lps[length-1]
					} else {
						lps[i] = 0
						i++
					}
				}
			}
			return lps
		}

		n, m := len(text), len(pattern)
		if m == 0 {
			return true
		}
		if n < m {
			return false
		}

		lps := buildLPS(pattern)
		i, j := 0, 0
		for i < n {
			if text[i] == pattern[j] {
				i++
				j++
				if j == m {
					return true
				}
			} else {
				if j != 0 {
					j = lps[j-1]
				} else {
					i++
				}
			}
		}
		return false
	}
	return kmpSearch(serialisation(root), serialisation(subRoot))
}

func isSubtreeZFunction(root *TreeNode, subRoot *TreeNode) bool {
	zFunction := func(s string) bool {
		n := len(s)
		z := make([]int, n)

		m := strings.Index(s, "|")

		l, r := 0, 0
		for i := 1; i < n; i++ {
			if i <= r {
				z[i] = min(r-i+1, z[i-l])
			}
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				z[i]++
			}
			if i+z[i]-1 > r {
				l = i
				r = i + z[i] - 1
			}
			if z[i] == m {
				return true
			}
		}
		return false
	}
	return zFunction(serialisation(subRoot) + "|" + serialisation(root))
}
