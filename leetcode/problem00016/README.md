# [16. 3Sum Closest](https://leetcode.com/problems/3sum-closest/)

## 题目

Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

```c
Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
```

## 题目大意

16\. 最接近的三数之和

给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

## 解题思路

跟第15题类似，我们这里采用排序后，用双指针

排序后，先固定一个数，然后双指针去找另外的两个数。如果小了，右移左边的指针；如果大了，左移右边的指针。
