# [62. 不同路径](https://leetcode-cn.com/problems/unique-paths/)

## 题目

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？

示例 1：

```c
输入：m = 3, n = 7
输出：28
```

示例 2：

```c
输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
```

示例 3：

```c
输入：m = 7, n = 3
输出：28
```

示例 3：

```c
输入：m = 3, n = 3
输出：6
```

## 解题思路

### 组合

组合问题， 总共要走 `m+n-2` 步，其中有 `m-1` 次向下移动，`n-1` 次向右移动。

组合公式：

C(r,n) = A(r,n)/r! = n!/(r!(n-r)!)

递推公式：

C(r,n)=C(r-1,n-1)+C(r-1,n)

### 动态规划

我们用dp[i][j]表示到坐标(i，j)这个格内有多少条不同的路径，所以最终的答案就是求dp[m-1][n-1]。

转换方程： dp[i][j] = dp[i-1][j] + dp[i][j-1]

当前坐标的值只和左边与上面的值有关，和其他的无关，所以我们可以把它改为一维数组。