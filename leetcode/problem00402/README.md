# [402. 移掉K位数字](https://leetcode-cn.com/problems/remove-k-digits/)

## 题目

给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。

注意:

* num 的长度小于 10002 且 ≥ k。
* num 不会包含任何前导零。

示例1:

```c
输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
```

示例2:

```c
输入: num = "10200", k = 1
输出: "200"
解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
```

示例3:

```c
输入: num = "10", k = 2
输出: "0"
解释: 从原数字移除所有的数字，剩余为空就是0。
```

## 解题思路

对于两个相同长度的数字序列，最左边不同的数字决定了这两个数字的大小。

例如，对于 A = 1axxx ，B = 1bxxx ，如果 a > b 则 A > B。

若要使得剩下的数字最小，需要保证靠前的数字尽可能小。

使用单调栈，如果栈顶元素比当前的大，则丢弃。
