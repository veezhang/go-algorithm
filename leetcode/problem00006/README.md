# [6. ZigZag Conversion](https://leetcode.com/problems/zigzag-conversion/)

## 题目

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```c
P   A   H   N
A P L S I I G
Y   I   R
```

And then read line by line: `"PAHNAPLSIIGYIR"`

Write the code that will take a string and make this conversion given a number of rows:

```c
string convert(string s, int numRows);
```

Example 1:

```c
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
```

Example 2:

```c
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
```

## 题目大意

将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

之后，你的输出需要从左往右逐行读取，产生出一个新的字符串。

## 解题思路

### 模拟法

申请numRows个数组，模拟Z字形，将数据插入到对应的行里面，最后将所有行连接

### 找规律

```c
numRows = 5
row=0    00              08              16              24
row=1    01          07  09          15  17          23  25
row=2    02      06      10      14      18      22
row=3    03  05          11  13          19  21
row=4    04              12              20

row = 0, 1, 2, 3, 4

所有竖排的等差都是：8，也就是： 2 * numRows -2

中间部分的也可以看成小的N，比如：
row=0    00              08
row=1    01          07
row=2    02      06
row=3    03  05
row=4    04

row = 1: 01 - 07: 6 = 2 * (numRows - row) -2
row = 2: 01 - 06: 4 = 2 * (numRows - row) -2
row = 3: 01 - 06: 2 = 2 * (numRows - row) -2

第一行和最后一行：
    按照row + i * (2 * numRows - 2)获得下标（竖着的），记为：index
中间的行一次循环有两个字符：
    按照row + i * (2 * numRows - 2)获得第一个下标（竖着的），记为：index
    再加上中间的 index + (2 * (numRows - row) - 2)
所有的下标，如果超出字符串长度，进行下一行
```
