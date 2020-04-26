# [13. Roman to Integer](https://leetcode.com/problems/roman-to-integer/)

## 题目

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

```c
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
```

For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

* I can be placed before V (5) and X (10) to make 4 and 9.
* X can be placed before L (50) and C (100) to make 40 and 90.
* C can be placed before D (500) and M (1000) to make 400 and 900.

Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.

Example1:

```c
Input: "III"
Output: 3
```

Example2:

```c
Input: "IV"
Output: 4
```

Example3:

```c
Input: "IX"
Output: 9
```

Example4:

```c
Input: "LVIII"
Output: 58
Explanation: L = 50, V = 5, III = 3.
```

Example5:

```c
Input: "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
```

## 题目大意

罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

```c
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
```

例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

* I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
* X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
* C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

## 解题思路

和第 12 题相反，这里是罗马数转整数。

遍历罗马数字，罗马数字有个特点：

包含 4 和 9 的情况下，都是在前面加一个小的值，比如: IV(5 - 1)， IX(10 - 1)， XL(50 - 10)，XC(100-10)，都是后面一个字符表示的值减去前面数字的值。其他情况就是累加的，1，2，3，5，6，7，8。

所以我们可以从第二个数字开始遍历，记住前面一个数字，视情况来决定加还是减。
