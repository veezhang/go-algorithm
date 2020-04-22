# [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

## 题目

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

```c
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```

Example 2:

```c
Input: "cbbd"
Output: "bb"
```

## 题目大意

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

回文是一个正读和反读都相同的字符串。

## 解题思路

### 最长公共子串

根据回文串的定义，正着和反着读一样，那我们是不是把原来的字符串倒置了，然后找最长的公共子串就可以了。例如 S = "caba" ，S = "abac"，最长公共子串是 "aba"，所以原字符串的最长回文串就是 "aba"。

但是， 当S="abc435cba"，S="abc534cba"，最长公共子串是 "abc" 和 "cba"，但很明显这两个字符串都不是回文串。所以我们求出最长公共子串后，并不一定是回文串，我们还需要判断该字符串倒置前的下标和当前的字符串下标是不是匹配。

### 动态规划（DP）

动态规划问题的一般形式就是`求最值`， 这里也就是最长回文。
存在`重叠子问题`。
能够通过子问题堆到出原问题，也就是要找出`状态转移方程`。
确认一些，`边界问题`，
是否可以`剪枝`。
具体看实现及注释。

### 扩展中心

中心扩散法的思路是：遍历每一个索引，以这个索引为中心，利用“回文串”中心对称的特点，往两边扩散，看最多能扩散多远。

### Manacher's Algorithm 马拉车算法

在每个字符间插入 "#"，并且在两端分别插入 "^" 和 "$"，中心扩展不需要边界检查。经过处理，字符串的长度永远都是奇数了。
利用回文的对称性，根据左边的可以推断出右边的值。
