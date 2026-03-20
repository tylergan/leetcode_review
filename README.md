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

### Backtrack

📄 **[View Cheat Sheet](https://tylergan.github.io/leetcode_review/backtrack/cheatsheet.html)** - Quick reference guide for backtrack problems

| Problem | Description | Solution | Tests |
|---------|-------------|----------|-------|
| Combination Sum | You are given an array of distinct integers nums and a target integer target. Yo... | [combination_sum.go](backtrack/combination_sum.go) | [✓](backtrack/combination_sum_test.go) |
| Subset | Given an array nums of unique integers, return all possible subsets of nums. | [subset.go](backtrack/subset.go) | [✓](backtrack/subset_test.go) |

### Binary_Search

📄 **[View Cheat Sheet](https://tylergan.github.io/leetcode_review/binary_search/cheatsheet.html)** - Quick reference guide for binary_search problems

| Problem | Description | Solution | Tests |
|---------|-------------|----------|-------|
| Binary Search | You are given an array of distinct integers nums, sorted in ascending order, and... | [binary_search.go](binary_search/binary_search.go) | [✓](binary_search/binary_search_test.go) |
| Finding Min Rotated Sorted | You are given an array of length n which was originally sorted in ascending orde... | [finding_min_rotated_sorted.go](binary_search/finding_min_rotated_sorted.go) | [✓](binary_search/finding_min_rotated_sorted_test.go) |
| Koko Eating Bananas | You are given an integer array piles where piles[i] is the number of bananas in ... | [koko_eating_bananas.go](binary_search/koko_eating_bananas.go) | [✓](binary_search/koko_eating_bananas_test.go) |
| Median Of Two Sorted Arrs | You are given two integer arrays nums1 and nums2 of size m and n respectively, w... | [median_of_two_sorted_arrs.go](binary_search/median_of_two_sorted_arrs.go) | [✓](binary_search/median_of_two_sorted_arrs_test.go) |
| Search 2D Matrix | You are given an m x n 2-D integer array matrix and an integer target. | [search_2d_matrix.go](binary_search/search_2d_matrix.go) | [✓](binary_search/search_2d_matrix_test.go) |
| Time Based Kv Store | Implement a time-based key-value data structure that supports: | [time_based_kv_store.go](binary_search/time_based_kv_store.go) | [✓](binary_search/time_based_kv_store_test.go) |

### Binary_Tree

📄 **[View Cheat Sheet](https://tylergan.github.io/leetcode_review/binary_tree/cheatsheet.html)** - Quick reference guide for binary_tree problems

| Problem | Description | Solution | Tests |
|---------|-------------|----------|-------|
| Balanced Binary Tree | Given a binary tree, return true if it is height-balanced and false otherwise. | [balanced_binary_tree.go](binary_tree/balanced_binary_tree.go) | [✓](binary_tree/balanced_binary_tree_test.go) |
| Binary Tree Max Path Sum | Given the root of a non-empty binary tree, return the maximum path sum of any no... | [binary_tree_max_path_sum.go](binary_tree/binary_tree_max_path_sum.go) | [✓](binary_tree/binary_tree_max_path_sum_test.go) |
| Btree Level Order Traversal | Given a binary tree root, return the level order traversal of it as a nested lis... | [btree_level_order_traversal.go](binary_tree/btree_level_order_traversal.go) | [✓](binary_tree/btree_level_order_traversal_test.go) |
| Btree Right Side View | You are given the root of a binary tree. Return only the values of the nodes tha... | [btree_right_side_view.go](binary_tree/btree_right_side_view.go) | [✓](binary_tree/btree_right_side_view_test.go) |
| Construct Tree From Pre To In | You are given two integer arrays preorder and inorder. | [construct_tree_from_pre_to_in.go](binary_tree/construct_tree_from_pre_to_in.go) | [✓](binary_tree/construct_tree_from_pre_to_in_test.go) |
| Diameter Of Binary Tree | The diameter of a binary tree is defined as the length of the longest path betwe... | [diameter_of_binary_tree.go](binary_tree/diameter_of_binary_tree.go) | [✓](binary_tree/diameter_of_binary_tree_test.go) |
| Good Nodes Btree | Within a binary tree, a node x is considered good if the path from the root of t... | [good_nodes_btree.go](binary_tree/good_nodes_btree.go) | [✓](binary_tree/good_nodes_btree_test.go) |
| Invert Binary Tree | You are given the root of a binary tree root. Invert the binary tree and return ... | [invert_binary_tree.go](binary_tree/invert_binary_tree.go) | [✓](binary_tree/invert_binary_tree_test.go) |
| Kth Smallest | Given the root of a binary search tree, and an integer k, return the kth smalles... | [kth_smallest.go](binary_tree/kth_smallest.go) | [✓](binary_tree/kth_smallest_test.go) |
| Lowest Common Ancestor | Given a binary search tree (BST) where all node values are unique, and two nodes... | [lowest_common_ancestor.go](binary_tree/lowest_common_ancestor.go) | [✓](binary_tree/lowest_common_ancestor_test.go) |
| Max Depth Binary Tree | Given the root of a binary tree, return its depth. | [max_depth_binary_tree.go](binary_tree/max_depth_binary_tree.go) | [✓](binary_tree/max_depth_binary_tree_test.go) |
| Same Binary Tree | Given the roots of two binary trees p and q, return true if the trees are equiva... | [same_binary_tree.go](binary_tree/same_binary_tree.go) | [✓](binary_tree/same_binary_tree_test.go) |
| Serialise And Deserialise Btree | Implement an algorithm to serialize and deserialize a binary tree. | [serialise_and_deserialise_btree.go](binary_tree/serialise_and_deserialise_btree.go) | [✓](binary_tree/serialise_and_deserialise_btree_test.go) |
| Subtree Of Another | Given the roots of two binary trees root and subRoot, return true if there is a ... | [subtree_of_another.go](binary_tree/subtree_of_another.go) | [✓](binary_tree/subtree_of_another_test.go) |
| Valid Btree | Given the root of a binary tree, return true if it is a valid binary search tree... | [valid_btree.go](binary_tree/valid_btree.go) | [✓](binary_tree/valid_btree_test.go) |

### Heap_Pq

📄 **[View Cheat Sheet](https://tylergan.github.io/leetcode_review/heap_pq/cheatsheet.html)** - Quick reference guide for heap_pq problems

| Problem | Description | Solution | Tests |
|---------|-------------|----------|-------|
| Design Twitter | Implement a simplified version of Twitter which allows users to post tweets, fol... | [design_twitter.go](heap_pq/design_twitter.go) | [✓](heap_pq/design_twitter_test.go) |
| K Closest Points To Origin | You are given an 2-D array points where points[i] = [xi, yi] represents the coor... | [k_closest_points_to_origin.go](heap_pq/k_closest_points_to_origin.go) | [✓](heap_pq/k_closest_points_to_origin_test.go) |
| Kth Largest Elem | Given an unsorted array of integers nums and an integer k, return the kth larges... | [kth_largest_elem.go](heap_pq/kth_largest_elem.go) | [✓](heap_pq/kth_largest_elem_test.go) |
| Kth Largest Element In Stream | Design a class to find the kth largest integer in a stream of values, including ... | [kth_largest_element_in_stream.go](heap_pq/kth_largest_element_in_stream.go) | [✓](heap_pq/kth_largest_element_in_stream_test.go) |
| Last Stone Weight | You are given an array of integers stones where stones[i] represents the weight ... | [last_stone_weight.go](heap_pq/last_stone_weight.go) | [✓](heap_pq/last_stone_weight_test.go) |
| Median In Data Stream | The median is the middle value in a sorted list of integers. For lists of even l... | [median_in_data_stream.go](heap_pq/median_in_data_stream.go) | [✓](heap_pq/median_in_data_stream_test.go) |
| Task Scheduler | You are given an array of CPU tasks tasks, where tasks[i] is an uppercase englis... | [task_scheduler.go](heap_pq/task_scheduler.go) | [✓](heap_pq/task_scheduler_test.go) |

### Linked_List

📄 **[View Cheat Sheet](https://tylergan.github.io/leetcode_review/linked_list/cheatsheet.html)** - Quick reference guide for linked_list problems

| Problem | Description | Solution | Tests |
|---------|-------------|----------|-------|
| Add Two Numbers | You are given two non-empty linked lists, l1 and l2, where each represents a non... | [add_two_numbers.go](linked_list/add_two_numbers.go) | [✓](linked_list/add_two_numbers_test.go) |
| Copy Ll W Random Ptr | You are given the head of a linked list of length n. Unlike a singly linked list... | [copy_ll_w_random_ptr.go](linked_list/copy_ll_w_random_ptr.go) | [✓](linked_list/copy_ll_w_random_ptr_test.go) |
| Cycle Detection | Given the beginning of a linked list head, return true if there is a cycle in th... | [cycle_detection.go](linked_list/cycle_detection.go) | [✓](linked_list/cycle_detection_test.go) |
| Find Duplicate | You are given an array of integers nums containing n + 1 integers. Each integer ... | [find_duplicate.go](linked_list/find_duplicate.go) | [✓](linked_list/find_duplicate_test.go) |
| Lru Cache | Implement the Least Recently Used (LRU) cache class LRUCache. The class should s... | [lru_cache.go](linked_list/lru_cache.go) | [✓](linked_list/lru_cache_test.go) |
| Merge K Sorted Lists | You are given an array of k linked lists lists, where each list is sorted in asc... | [merge_k_sorted_lists.go](linked_list/merge_k_sorted_lists.go) | [✓](linked_list/merge_k_sorted_lists_test.go) |
| Merge Two Sorted Ll | You are given the heads of two sorted linked lists list1 and list2. | [merge_two_sorted_ll.go](linked_list/merge_two_sorted_ll.go) | [✓](linked_list/merge_two_sorted_ll_test.go) |
| Remove Nth | You are given the beginning of a linked list head, and an integer n. | [remove_nth.go](linked_list/remove_nth.go) | [✓](linked_list/remove_nth_test.go) |
| Reorder Ll | You are given the head of a singly linked-list. | [reorder_ll.go](linked_list/reorder_ll.go) | [✓](linked_list/reorder_ll_test.go) |
| Reverse Ll | Given the beginning of a singly linked list head, reverse the list, and return t... | [reverse_ll.go](linked_list/reverse_ll.go) | [✓](linked_list/reverse_ll_test.go) |
| Reverse Nodes In K Group | You are given the head of a singly linked list head and a positive integer k. | [reverse_nodes_in_k_group.go](linked_list/reverse_nodes_in_k_group.go) | [✓](linked_list/reverse_nodes_in_k_group_test.go) |

### Sliding_Window

📄 **[View Cheat Sheet](https://tylergan.github.io/leetcode_review/sliding_window/cheatsheet.html)** - Quick reference guide for sliding_window problems

| Problem | Description | Solution | Tests |
|---------|-------------|----------|-------|
| Best Time Buy Sell Stock | You are given an integer array prices where prices[i] is the price of NeetCoin o... | [best_time_buy_sell_stock.go](sliding_window/best_time_buy_sell_stock.go) | [✓](sliding_window/best_time_buy_sell_stock_test.go) |
| Longest Repeating Char Replacement | You are given a string s consisting of only uppercase english characters and an ... | [longest_repeating_char_replacement.go](sliding_window/longest_repeating_char_replacement.go) | [✓](sliding_window/longest_repeating_char_replacement_test.go) |
| Longest Substr Wo Repeating Chars | Given a string s, find the length of the longest substring without duplicate cha... | [longest_substr_wo_repeating_chars.go](sliding_window/longest_substr_wo_repeating_chars.go) | [✓](sliding_window/longest_substr_wo_repeating_chars_test.go) |
| Minimim Window Substring | Given two strings s and t, return the shortest substring of s such that every ch... | [minimim_window_substring.go](sliding_window/minimim_window_substring.go) | [✓](sliding_window/minimim_window_substring_test.go) |
| Permutation In String | You are given two strings s1 and s2. | [permutation_in_string.go](sliding_window/permutation_in_string.go) | [✓](sliding_window/permutation_in_string_test.go) |
| Sliding Window Maximum | You are given an array of integers nums and an integer k. There is a sliding win... | [sliding_window_maximum.go](sliding_window/sliding_window_maximum.go) | [✓](sliding_window/sliding_window_maximum_test.go) |

### Stack

📄 **[View Cheat Sheet](https://tylergan.github.io/leetcode_review/stack/cheatsheet.html)** - Quick reference guide for stack problems

| Problem | Description | Solution | Tests |
|---------|-------------|----------|-------|
| Car Fleet | There are n cars traveling to the same destination on a one-lane highway. | [car_fleet.go](stack/car_fleet.go) | [✓](stack/car_fleet_test.go) |
| Daily Temperatures | You are given an array of integers temperatures where temperatures[i] represents... | [daily_temperatures.go](stack/daily_temperatures.go) | [✓](stack/daily_temperatures_test.go) |
| Eval Rpn | You are given an array of strings tokens that represents a valid arithmetic expr... | [eval_RPN.go](stack/eval_RPN.go) | [✓](stack/eval_RPN_test.go) |
| Largest Rectangle | You are given an array of integers heights where heights[i] represents the heigh... | [largest_rectangle.go](stack/largest_rectangle.go) | [✓](stack/largest_rectangle_test.go) |
| Minimum Stack | Design a stack class that supports the push, pop, top, and getMin operations. | [minimum_stack.go](stack/minimum_stack.go) | [✓](stack/minimum_stack_test.go) |
| Valid Parantheses | You are given a string s consisting of the following characters: '(', ')', '{', ... | [valid_parantheses.go](stack/valid_parantheses.go) | [✓](stack/valid_parantheses_test.go) |

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
