# [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## 题目

Given a string, find the length of the longest substring without repeating characters.

Example 1:

```c
Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

Example 2:

```c
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

Example 3:

```c
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

## 题目大意

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

## 解题思路

滑动窗口
如果 s[j] 在 s[i,j]中，其下标是j1, 则i滑动到 j1 + 1, 计算值，j滑动j++
如果 s[j] 不在 s[i,j]中，计算值，j滑动j++
