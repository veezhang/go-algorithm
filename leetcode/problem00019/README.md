# [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

## 题目

Given a linked list, remove the n-th node from the end of list and return its head.

Example:

```c
Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
```

Note:

Given n will always be valid.

## 题目大意

19\. 删除链表的倒数第N个节点

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

## 解题思路

快慢指针，快指针先行 n + 1 步，然后快慢指针一起移动
