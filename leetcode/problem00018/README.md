# [18. 4Sum](https://leetcode.com/problems/4sum/)

## 题目

Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:

```c
Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
```

## 题目大意

18\. 四数之和

给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

## 解题思路

跟第15题类似，我们这里采用排序后，用双指针

排序后，先固定一个数，然后双指针去找另外的两个数。如果小了，右移左边的指针；如果大了，左移右边的指针。
