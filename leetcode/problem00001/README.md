# [1. Two Sum](https://leetcode.com/problems/two-sum/)

## 题目

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

```c
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## 题目大意

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

## 解题思路

给定数 n, 目标是target，那么另一个数就是target - n
用一个map映射起来，如果存在map[target - n]，则找到了，否则存入map[n]=i的下标
