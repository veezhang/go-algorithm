# [240. Search a 2D Matrix II](https://leetcode.com/problems/search-a-2d-matrix-ii/)

## 题目

Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

* Integers in each row are sorted in ascending from left to right.
* Integers in each column are sorted in ascending from top to bottom.

Example:

Consider the following matrix:

```go
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
```

Given target = `5`, return `true`.

Given target = `20`, return `false`.

## 题目大意

240\. 搜索二维矩阵 II

编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

* 每行的元素从左到右升序排列。
* 每列的元素从上到下升序排列。

## 解题思路

### 二分查找法

因为是有序的，可以使用二分查找法

### 步进法

根据题意得知，数据是有顺序性质的，可以从右上或者左下开始查找，比如 target = `20`

```go
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
// 从 15 开始
// 15 比 20 小，往下走
// 19 比 20 小，往下走
// 22 比 20 大，往左走
// 16 比 20 小，往下走
// 17 比 20 小，往下走
// 26 比 20 大，往左走
// 23 比 20 大，往左走
// 21 比 20 大，往左走
// 18 比 20 小，往下走
// 越界了，未找到
```

时间复杂度： O(m+n)

### 子矩阵法

```go
[
  [1,   4,  7, | 11, 15],
  [2,   5,  8, | 12, 19],
  [3,   6,  9, | 16, 22],
  [10, 13, 14, | 17, 24],
---------------|----------
  [18, 21, 23, | 26, 30]
]

// 比如查找 18
// 取中间列， 第 3 列
// 从行来获取小于等于 target 最大的行
// 则左上角和右下角不可能包含 target 了
```
