# [11. Container With Most Water](https://leetcode.com/problems/container-with-most-water/)

## 题目

Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example:

```c
Input: [1,8,6,2,5,4,8,3,7]
Output: 49
```

## 题目大意

给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

## 解题思路

### 双指针

面积为： (right - left) * min(height[right], height[left])

假设移动长柱子，则 right - left 肯定会变小， min(height[right], height[left]) 也不可能超过之前的值，所以我们移动矮柱子

如果 height[left] <= height[right] 则 left++ ；否则 right--
