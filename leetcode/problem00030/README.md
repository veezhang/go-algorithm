# [30. Substring with Concatenation of All Words](https://leetcode.com/problems/substring-with-concatenation-of-all-words/)

## 题目

You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.

Example1:

```c
Input:
  s = "barfoothefoobarman",
  words = ["foo","bar"]
Output: [0,9]
Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
The output order does not matter, returning [9,0] is fine too.
```

Example1:

```c
Input:
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
Output: []
```

## 题目大意

30\. 串联所有单词的子串

给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

## 解题思路

用 map 存 words 中的字符串，注意： words 中字符串可能重复

### 循环遍历

循环遍历字符串，找到一个匹配后，判断接下来的是否匹配。

### 滑动窗口 (Sliding Window)

在循环遍历中可能出现很多重复的计算，可以用滑动窗口来每次移出去一个 word 并添加一个 word 。

然后对 word 长度都滑动窗口算一遍。

比如：

```go
s = "barfoothefoobarman"
words = ["foo","bar"]
// 滑动窗口宽度是： 6 ，每次移动 3
// 那么需要从：[0,3) , [1,4) , [2,5) 为起始，分别向后滑动
// [0,6)   [3,9)   [6,12)  [9,15)  [12,18)
// [1,7)   [4,10)  [7,13)  [10,16)
// [2,8)   [5,11)  [8,14)  [11,16)
// 这样不会少也不会多
```
