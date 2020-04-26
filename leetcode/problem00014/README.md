# [14. Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/)

## 题目

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example1:

```c
Input: ["flower","flow","flight"]
Output: "fl"
```

Example2:

```c
Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
```

Note:

All given inputs are in lowercase letters a-z.

## 题目大意

编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""。

## 解题思路

遍历第一个字符串和其他的对比。
