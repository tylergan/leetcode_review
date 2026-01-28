package binary_tree

import "strconv"

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
Approach 2: Serialization + KMP Pattern Matching
Time: O(m + n) where m = nodes in root, n = nodes in subRoot
Space: O(m + n) for serialized strings

Instead of comparing trees recursively (O(m*n) worst case), we:
1. Serialize both trees to strings using preorder traversal
2. Check if serialized subRoot is a substring of serialized root using KMP

Why separators matter:
- Without separators, tree [12] and tree [2] would have serializations "12##" and "2##"
- "2##" is a substring of "12##" but [2] is NOT a subtree of [12]
- Using "$" prefix: "$12$#$#" vs "$2$#$#" - no false match!
*/

// serialize converts a tree to a string using preorder traversal.
// Each node is prefixed with "$" to prevent false substring matches.
// Null nodes are represented as "$#".
func serialize(root *TreeNode) string {
	if root == nil {
		return "$#"
	}
	return "$" + strconv.Itoa(root.Val) + serialize(root.Left) + serialize(root.Right)
}

/*
KMP (Knuth-Morris-Pratt) Algorithm

The key insight: when a mismatch occurs, we've already matched some characters.
The LPS array tells us how much of that matched portion is ALSO a prefix of the pattern,
so we don't need to re-compare those characters.

LPS = Longest Proper Prefix which is also Suffix
- "Proper" means the prefix can't be the entire string

Example: pattern = "ABABC"
Index:    0  1  2  3  4
Pattern:  A  B  A  B  C
LPS:      0  0  1  2  0

LPS[3] = 2 means: for "ABAB", the longest proper prefix that's also a suffix is "AB" (length 2)
- Prefixes of "ABAB": "A", "AB", "ABA"
- Suffixes of "ABAB": "B", "AB", "BAB"
- Longest match: "AB" → length 2

Why this helps:
If we're matching pattern "ABABC" against text and fail at index 4 (the 'C'):
  Text:    ...A B A B X...
  Pattern:    A B A B C
                      ^ mismatch at pattern[4]

We know text has "ABAB" before the mismatch. LPS[3]=2 tells us "AB" at the end
of our matched portion is also "AB" at the start of the pattern. So we can
continue matching from pattern[2] instead of starting over:

  Text:    ...A B A B X...
  Pattern:        A B A B C
                  ^ continue from pattern[2]
*/

// buildLPS constructs the Longest Proper Prefix Suffix array for the pattern.
// lps[i] = length of longest proper prefix of pattern[0..i] that is also a suffix
func buildLPS(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)
	lps[0] = 0 // first character has no proper prefix

	length := 0 // length of previous longest prefix suffix
	i := 1

	for i < m {
		if pattern[i] == pattern[length] {
			// Characters match: extend the prefix-suffix
			length++
			lps[i] = length
			i++
		} else {
			// Mismatch
			if length != 0 {
				// Don't increment i, try shorter prefix-suffix
				// Key insight: lps[length-1] gives us the next longest prefix-suffix to try
				length = lps[length-1]
			} else {
				// No prefix-suffix exists
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

// kmpSearch returns true if pattern exists as a substring in text.
// Uses the LPS array to skip redundant comparisons.
func kmpSearch(text, pattern string) bool {
	n, m := len(text), len(pattern)
	if m == 0 {
		return true
	}
	if n < m {
		return false
	}

	lps := buildLPS(pattern)

	i := 0 // index for text
	j := 0 // index for pattern

	for i < n {
		if text[i] == pattern[j] {
			// Characters match, advance both
			i++
			j++

			if j == m {
				// Found complete pattern match!
				return true
			}
		} else {
			// Mismatch
			if j != 0 {
				// Use LPS to skip: we know pattern[0..j-1] matched,
				// and lps[j-1] characters of that are also a prefix.
				// So continue matching from pattern[lps[j-1]]
				j = lps[j-1]
			} else {
				// j == 0, can't fall back further, advance text pointer
				i++
			}
		}
	}
	return false
}

func isSubtreeKMP(root *TreeNode, subRoot *TreeNode) bool {
	serializedRoot := serialize(root)
	serializedSubRoot := serialize(subRoot)
	return kmpSearch(serializedRoot, serializedSubRoot)
}

/*
Approach 3: Serialization + Z-Function Pattern Matching
Time: O(m + n) where m = nodes in root, n = nodes in subRoot
Space: O(m + n) for serialized strings and Z-array

Z-Function Algorithm

The Z-array for a string s is defined as:
z[i] = length of the longest substring starting at s[i] that matches a prefix of s

Example: s = "aabxaab"
Index:  0  1  2  3  4  5  6
String: a  a  b  x  a  a  b
Z:      -  1  0  0  3  1  0

z[0] is undefined (or set to 0/n depending on convention)
z[1] = 1: s[1..1] = "a" matches prefix s[0..0] = "a"
z[4] = 3: s[4..6] = "aab" matches prefix s[0..2] = "aab"

How to use Z-function for pattern matching:
1. Create combined string: pattern + "|" + text
   ("|" is a separator that doesn't appear in pattern or text)
2. Compute Z-array for combined string
3. If any z[i] == len(pattern) for i > len(pattern), pattern exists in text!

Example: Find "aab" in "xaabx"
Combined: "aab|xaabx"
Index:     0 1 2 3 4 5 6 7 8
String:    a a b | x a a b x
Z:         - 1 0 0 0 3 1 0 0
                     ^ z[5]=3 equals len("aab"), so pattern found!

The Z-function optimization (the "Z-box"):
We maintain [l, r] = the rightmost interval where s[l..r] matches s[0..r-l].
For each new position i:
- If i > r: compute z[i] naively from scratch, update [l,r] if we extend past r
- If i <= r: we already know s[i..r] = s[i-l..r-l], so z[i] >= min(z[i-l], r-i+1)
  Then extend if possible and update [l,r]

This ensures each character is compared at most twice: once when extending r,
once when being the starting point of a new z[i] computation.
*/

// zFunction computes the Z-array for string s.
// z[i] = length of longest substring starting at s[i] that matches a prefix of s
func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	// z[0] is undefined; we leave it as 0

	l, r := 0, 0 // [l, r] is the rightmost Z-box (interval matching a prefix)

	for i := 1; i < n; i++ {
		if i <= r {
			// i is inside the current Z-box
			// s[i..r] is known to equal s[i-l..r-l]
			// So z[i] is at least min(z[i-l], r-i+1)
			z[i] = min(r-i+1, z[i-l])
		}

		// Try to extend z[i] by comparing characters
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}

		// Update Z-box if we extended past r
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}

func isSubtreeZFunction(root *TreeNode, subRoot *TreeNode) bool {
	serializedRoot := serialize(root)
	serializedSubRoot := serialize(subRoot)

	// Combine: pattern + separator + text
	combined := serializedSubRoot + "|" + serializedRoot

	zValues := zFunction(combined)
	patternLen := len(serializedSubRoot)

	// Check if pattern appears in text portion (after the separator)
	for i := patternLen + 1; i < len(combined); i++ {
		if zValues[i] == patternLen {
			return true
		}
	}
	return false
}

/*
KMP vs Z-Function Comparison:

Both achieve O(n + m) time for pattern matching. The key differences:

KMP:
- Builds LPS array on pattern only (O(m) space for pattern)
- Then scans text with two pointers
- More intuitive for "searching while reading" scenarios
- The LPS concept (prefix = suffix) is used during both preprocessing AND searching

Z-Function:
- Builds Z-array on combined string (O(n + m) space)
- Single pass algorithm - preprocessing IS the search
- More intuitive once you understand the Z-box optimization
- Each position directly tells you "how much of prefix matches here"

Mental Model:
- KMP: "When I fail, where in the pattern can I resume?"
- Z-Function: "At each position, how much of the beginning do I see?"
*/
