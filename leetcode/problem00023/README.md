# [23. Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/)

## 题目

Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

```c
Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
```

## 题目大意

23\. 合并K个排序链表

合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

## 解题思路

### K 个链表中找最小

根据21题一样，只是这里找 k 个中最小的来合并

### 一个一个的合并

根据21题一样，一个一个的合并到总的合并好的链表中

### 分治法 Divide and Conquer

每两个两个的合并

### 小根堆

优化 "K 个链表中找最小" 算法， K 个链表用小根堆来查找最小的
