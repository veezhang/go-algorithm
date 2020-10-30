# [287. 寻找重复数](https://leetcode-cn.com/problems/find-the-duplicate-number/)

## 题目

给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

示例1:

```c
输入: [1,3,4,2,2]
输出: 2
```

示例2:

```c
输入: [3,1,3,4,2]
输出: 3
```

说明：

* 不能更改原数组（假设数组是只读的）。
* 只能使用额外的 O(1) 的空间。
* 时间复杂度小于 O(n2) 。
* 数组中只有一个重复的数字，但它可能不止重复出现一次。

## 解题思路

### 二分查找

```go
// 定义 cnt[i] 表示 nums 数组中小于等于 i 的数有多少个
// 假设 target 为重复的数：
//      [1, target - 1]之间的所有数字满足 cnt[i] <= i
//      [target + 1, n]之间的数满足 cnt[i] > i
// cnt[i] 具有单调性
// 那么可以二分查找法来查找 target
```

### 二进制

将所有数二进制展开按位考虑如何找出重复的数，如果我们能确定重复数每一位是 `1` 还是 `0` 就可以按位还原出重复的数。

对于第 `i` 位， 假设

* `nums` 数组中二进制展开后第 `i` 位为 `1` 的数有 `x` 个
* `[1,n]` 这 `n` 个数二进制展开后第 `i` 位为 `1` 的数有 `y` 个

那么重复的数第 `i` 位为 `1` 当且仅当 `x > y` 。

### 快慢指针

把 `nums` 看成一个链表，数组的下标就是指向元素的指针，把数组的元素也看作指针。

假设有这样一个样例：`[1,2,3,4,5,6,7,8,9,5]`。

如果我们按照上面的循环下去就会得到这样一个路径: `1 2 3 4 5 [6 7 8 9] [6 7 8 9] [6 7 8 9] . . .`

这样就有了一个环，也就是6 7 8 9。point 会一直在环中循环的前进。

然后就演变成判断链表是否存在环了。