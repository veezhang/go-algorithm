# [242. 有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)

## 题目

给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

```c
输入: s = "anagram", t = "nagaram"
输出: true
```

示例 2:

```c
输入: s = "rat", t = "car"
输出: false
```

说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

## 解题思路

两个字符串中字符出现的种类和次数均相等

哈希，统计个数
