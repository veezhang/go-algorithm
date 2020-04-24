# [10. Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/)

## 题目

Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

```c
'.' Matches any single character.
'*' Matches zero or more of the preceding element.
```

The matching should cover the entire input string (not partial).

Note:

* s could be empty and contains only lowercase letters a-z.
* p could be empty and contains only lowercase letters a-z, and characters like . or *.

Example1:

```c
Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```

Example2:

```c
Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
```

Example3:

```c
Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
```

Example4:

```c
Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
```

Example5:

```c
Input:
s = "mississippi"
p = "mis*is*p*."
Output: false
```

## 题目大意

给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

```c
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
```

## 解题思路

`'.'` 可以可任何字符相等，只需要判断 s和p中字符是否相等，或者p中字符是否为 `'.'`。

`'*'` 字符比较特殊，我们重点需要关注这个。

### 回溯法

`'*'` 能够匹配 0 次或者多次，我们假定为 0 次或者 1 次，因为多次也就是 1 次 1 次的匹配的。

首先让其匹配 0 次，如果失败了回溯回来匹配 1 次试试。

### 动态规划

再回溯法里面，我们其实可以发现有很多重复计算，也就是重叠子问题，就是典型的动态规划。

dp[i][j] 表示 s[i:] (i 以及后面的) 和 p[j:] (j 以及后面的)  已经匹配了。这里是从后往前推进的。

根据依赖关系，确保不会在使用之前已经被修改了，可以对 dp 进行空间优化。
