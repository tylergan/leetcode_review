# LeetCode Review

A collection of LeetCode problems solved in Go with comprehensive test coverage.

## 📚 Problem Categories

### Arrays

📄 **[View Cheat Sheet](https://tylergan.github.io/leetcode_review/arrays/cheatsheet.html)** - Quick reference guide for arrays problems

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

### Sliding_Window

| Problem | Description | Solution | Tests |
|---------|-------------|----------|-------|
| Best Time Buy Sell Stock | You are given an integer array prices where prices[i] is the price of NeetCoin o... | [best_time_buy_sell_stock.go](sliding_window/best_time_buy_sell_stock.go) | [✓](sliding_window/best_time_buy_sell_stock_test.go) |
| Longest Substr Wo Repeating Chars | Given a string s, find the length of the longest substring without duplicate cha... | [longest_substr_wo_repeating_chars.go](sliding_window/longest_substr_wo_repeating_chars.go) | [✓](sliding_window/longest_substr_wo_repeating_chars_test.go) |

### Two_Pointers

📄 **[View Cheat Sheet](https://tylergan.github.io/leetcode_review/two_pointers/cheatsheet.html)** - Quick reference guide for two_pointers problems

| Problem | Description | Solution | Tests |
|---------|-------------|----------|-------|
| 3Sum | Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]... | [3Sum.go](two_pointers/3Sum.go) | [✓](two_pointers/3Sum_test.go) |
| Container With Most Water | You are given an integer array heights where heights[i] represents the height of... | [container_with_most_water.go](two_pointers/container_with_most_water.go) | [✓](two_pointers/container_with_most_water_test.go) |
| Trapping Rain Water | You are given an array of non-negative integers height which represent an elevat... | [trapping_rain_water.go](two_pointers/trapping_rain_water.go) | [✓](two_pointers/trapping_rain_water_test.go) |
| Two Integer Sum 2 | Given an array of integers numbers that is sorted in non-decreasing order. | [two_integer_sum_2.go](two_pointers/two_integer_sum_2.go) | [✓](two_pointers/two_integer_sum_2_test.go) |
| Valid Palindrome | Given a string s, return true if it is a palindrome, otherwise return false. | [valid_palindrome.go](two_pointers/valid_palindrome.go) | [✓](two_pointers/valid_palindrome_test.go) |

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
- Updates this README automatically
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

## ⚙️ Configuration

Customize the README generator by editing `.leetcode-config.yml`:
- Add folders to ignore when scanning for categories
- Configure GitHub Pages settings

## 🔄 Updating This README

This README is automatically generated. To update it manually, run:
```bash
python3 scripts/update_readme.py
```
