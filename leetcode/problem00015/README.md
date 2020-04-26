# [15. 3Sum](https://leetcode.com/problems/3sum/)

## 题目

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

```c
Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
```

## 题目大意

给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

## 解题思路

### Hash

twoSum 算法类似，这里使用两级 hash ， 注意不能重复，记得判断

### 双指针

排序后，先固定一个数，然后双指针去找另外的两个数。如果小了，右移左边的指针；如果大了，左移右边的指针。
