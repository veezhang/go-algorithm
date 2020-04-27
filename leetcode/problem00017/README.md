# [17. Letter Combinations of a Phone Number](https://leetcode.com/problems/two-sum/)

## 题目

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

```c
| 1 !@#  | 2 abc | 3 def  |
| 4 ghi  | 5 jkl | 6 mno  |
| 7 pqrs | 8 tuv | 9 wxyz |
| *      | 0     | #      |
```

Example:

```c
Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
```

Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.

## 题目大意

17\. 电话号码的字母组合

给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

## 解题思路

### 直接定位

其实能够直到总共都少种情况，每种情况多少字符，也能直接定位字符位置，预申请，没有 append 消耗

### 回溯

处理第一个字符， 递归调用接下来的字符，处理完成后添加到结果列表。
