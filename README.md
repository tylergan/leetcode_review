# LeetCode Review

A collection of LeetCode problems solved in Go with comprehensive test coverage.

## 📚 Problem Categories

### Arrays

📄 **[View Cheat Sheet](https://tylergan.github.io/LeetcodeReview/arrays/cheatsheet.html)** - Quick reference guide for arrays problems

| Problem | Description | Solution | Tests |
|---------|-------------|----------|-------|
| Encode Decode String | Design an algorithm to encode a list of strings to a single string. The encoded ... | [encode_decode_string.go](arrays/encode_decode_string.go) | [✓](arrays/encode_decode_string_test.go) |
| Find Duplicates | Given an integer array nums, return true if any value appears more than once in ... | [find_duplicates.go](arrays/find_duplicates.go) | [✓](arrays/find_duplicates_test.go) |
| Group Anagrams | Given an array of strings strs, group all anagrams together into sublists. You m... | [group_anagrams.go](arrays/group_anagrams.go) | [✓](arrays/group_anagrams_test.go) |
| Longest Consecutive Subsequence | Given an array of integers nums, return the length of the longest consecutive se... | [longest_consecutive_subsequence.go](arrays/longest_consecutive_subsequence.go) | [✓](arrays/longest_consecutive_subsequence_test.go) |
| Product Of Array Except Self | Given an integer array nums, return an array output where output[i] is the produ... | [product_of_array_except_self.go](arrays/product_of_array_except_self.go) | [✓](arrays/product_of_array_except_self_test.go) |
| Top Freq K Elems | Given an integer array nums and an integer k, return the k most frequent element... | [top_freq_k_elems.go](arrays/top_freq_k_elems.go) | [✓](arrays/top_freq_k_elems_test.go) |
| Two Sum | Given an array of integers nums and an integer target, return the indices i and ... | [two_sum.go](arrays/two_sum.go) | [✓](arrays/two_sum_test.go) |
| Valid Anagram | Given two strings s and t, return true if the two strings are anagrams of each o... | [valid_anagram.go](arrays/valid_anagram.go) | [✓](arrays/valid_anagram_test.go) |
| Valid Sudoku | You are given a 9 x 9 Sudoku board. A Sudoku board is valid if the following rul... | [valid_sudoku.go](arrays/valid_sudoku.go) | [✓](arrays/valid_sudoku_test.go) |

## 🚀 Running Tests

Run all tests:
```bash
go test ./...
```

Run tests for a specific category:
```bash
go test ./arrays
```

Run tests for a specific problem:
```bash
go test ./arrays -run TestTwoSum
```

## 🎯 Pre-commit Hook

This repository includes a pre-commit hook that:
- Formats Go code with `go fmt`
- Runs all tests before allowing commits

The hook is automatically set up in `.git/hooks/pre-commit`.

## 📖 Cheat Sheets

Each category includes a comprehensive cheat sheet with:
- Common patterns and techniques
- Time and space complexity analysis
- Key insights and gotchas
- Code snippets

Cheat sheets are hosted on GitHub Pages. Click the links above to view them online!

## 🔄 Updating This README

This README is automatically generated. To update it, run:
```bash
python3 update_readme.py
```
