# [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)

## 题目

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

* Open brackets must be closed by the same type of brackets.
* Open brackets must be closed in the correct order.

Note that an empty string is also considered valid.

Example1:

```c
Input: "()"
Output: true
```

Example2:

```c
Input: "()[]{}"
Output: true
```

Example3:

```c
Input: "(]"
Output: false
```

Example4:

```c
Input: "([)]"
Output: false
```

Example5:

```c
Input: "{[]}"
Output: true
```

## 题目大意

20\. 有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

## 解题思路

栈
