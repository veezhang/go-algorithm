# [9. Palindrome Number](https://leetcode.com/problems/palindrome-number/)

## 题目

Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

```c
Input: 121
Output: true
```

Example 2:

```c
Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```

Example 3:

```c
Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```

Follow up:

Coud you solve it without converting the integer to a string?

## 题目大意

判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

## 解题思路

### 最高位和个位比

获取最高位和个位比较，然后去除最高位再除以10（去除个位），如此循环

### 倒转一半数字

比如： 123321， 倒转后面的 123 再与 前面的 123 相比较
