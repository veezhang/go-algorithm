# [22. Generate Parentheses](https://leetcode.com/problems/generate-parentheses/)

## 题目

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

```c
[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
```

## 题目大意

22\. 括号生成

数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

## 解题思路

回溯法，使用slice分配固定长度，不再append，添加结果的时候才拷贝一份
