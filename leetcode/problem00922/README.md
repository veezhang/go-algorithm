# [922. 按奇偶排序数组 II](https://leetcode-cn.com/problems/sort-array-by-parity-ii/)

## 题目

给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。

对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。

你可以返回任何满足上述条件的数组作为答案。

示例:

```c
输入：[4,2,5,7]
输出：[4,5,2,7]
解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
```

提示：

* 2 <= A.length <= 20000
* A.length % 2 == 0
* 0 <= A[i] <= 1000

## 解题思路

只要把偶数下标中所有的基数迁移到基数下标就好了

### 类似插入法

后面找到一个基数下标的偶数，插入到前面偶数下标的基数

### 双指针

前面的指向偶数，后面的指向基数，这里步长为 2

当前面发下不是偶数，后面不是基数的时候，互相换一下