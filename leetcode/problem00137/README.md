# [137. 只出现一次的数字 II](https://leetcode-cn.com/problems/single-number-ii/)

## 题目

给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

```c
输入: [2,2,3,2]
输出: 3
```

示例 2:

```c
输入: [0,1,0,1,0,1,99]
输出: 99
```

## 解题思路

使用位运算。

* 任何数和 `0` 做异或运算，结果仍然是原来的数，即 `0^a = a`
* 任何数和其自身做异或运算，结果是 `0` ，`a^a=0`
* 异或运算满足交换律和结合律，即 `a^b^a=a^a^b=0^b=b`

### 方案1

用 `bits1` 记录到当前处理的元素为止，二进制 `1` 出现 `1` 次的有哪些二进制位

用 `bits2` 记录到当前处理的元素为止，二进制 `2` 出现 `2` 次的有哪些二进制位

当 `bits1` 和 `bits2` 中的某一位同时为 `1` 时表示该二进制位上 `1` 出现了 `3` 次，此时需要清零。

### 方案2

见实现