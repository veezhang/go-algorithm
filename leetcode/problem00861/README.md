# [861. 翻转矩阵后的得分](https://leetcode-cn.com/problems/score-after-flipping-matrix/)

## 题目

有一个二维矩阵 A 其中每个元素的值为 0 或 1 。

移动是指选择任一行或列，并转换该行或列中的每一个值：将所有 0 都更改为 1，将所有 1 都更改为 0。

在做出任意次数的移动后，将该矩阵的每一行都按照二进制数来解释，矩阵的得分就是这些数字的总和。

返回尽可能高的分数。

示例:

```c
输入：[[0,0,1,1],[1,0,1,0],[1,1,0,0]]
输出：39
解释：
转换为 [[1,1,1,1],[1,0,0,1],[1,1,1,1]]
0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39
```

提示：

* 1 <= A.length <= 20
* 1 <= A[0].length <= 20
* A[i][j] 是 0 或 1

## 解题思路

要想使得总和最大，那么每一行都需要尽可能的大

所以先行进行翻转，让最左边为 1

然后再进行列的翻转，让每一列对应的数最大，也就是让每一列中 1 的数量多一些