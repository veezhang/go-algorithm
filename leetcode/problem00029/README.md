# [29. Divide Two Integers](https://leetcode.com/problems/divide-two-integers/)

## 题目

Given two integers `dividend` and `divisor`, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing `dividend` by `divisor`.

The integer division should truncate toward zero, which means losing its fractional part. For example, `truncate(8.345) = 8` and `truncate(-2.7335) = -2`.

Example1:

```c
Input: dividend = 10, divisor = 3
Output: 3
Explanation: 10/3 = truncate(3.33333..) = 3.
```

Example2:

```c
Input: dividend = 7, divisor = -3
Output: -2
Explanation: 7/-3 = truncate(-2.33333..) = -2.
```

Note:

* Both dividend and divisor will be 32-bit signed integers.
* The divisor will never be 0.
* Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.

## 题目大意

29\. 两数相除

给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2。

* 被除数和除数均为 32 位有符号整数。
* 除数不为 0。
* 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。

## 解题思路

不能使用乘法、除法和 mod 运算符，那么想到的就是 加减法，位移？

dividend / divisor = result

其实就是相当于 dividend - divisor 执行 result 次，统计减少了多少次就好

但是这样的话， dividend 很大， divisor 很小，那么次数 result 会很大，比如： dividend = 2^31 , divisor = 1 的时候

那么每次按照尽可能大的数来减，怎么来计算呢？

```c
// dividend 可以表示为，每次判断还没有到极限，就翻倍， a = a << 1
dividend = 2^0 * divisor + 2^1 * divisor + ... + 2^n * divisor
```
