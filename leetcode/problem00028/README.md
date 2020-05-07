# [28. Implement strStr()](https://leetcode.com/problems/implement-strstr/)

## 题目

Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example1:

```c
Input: haystack = "hello", needle = "ll"
Output: 2
```

Example2:

```c
Input: haystack = "aaaaa", needle = "bba"
Output: -1
```

Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

## 题目大意

实现 strStr() 函数。

## 解题思路

### 子串

取子串来对比

### 双指针

给定数 n, 目标是target，那么另一个数就是target - n
用一个map映射起来，如果存在map[target - n]，则找到了，否则存入map[n]=i的下标

### BF （Brute Force）

BF算法，即蛮力(Brute Force)算法，是普通的模式匹配算法，BF算法的思想就是将目标串 T 的第一个字符与模式串 P 的第一个字符进行匹配，若相等，则继续比较 T 的第二个字符和 P 的第二个字符；若不相等，则比较 T 的第二个字符和 P 的第一个字符，依次比较下去，直到得出最后的匹配结果。蛮力解法，和前面的 子串 和  双指针 意思差不多。

### KMP

KMP 算法是一种改进的字符串匹配算法，由D.E.Knuth，J.H.Morris和V.R.Pratt提出的，因此人们称它为克努特—莫里斯—普拉特操作（简称KMP算法）。KMP算法的核心是利用匹配失败后的信息，尽量减少模式串与主串的匹配次数以达到快速匹配的目的。

主要是通过找公共前后缀来达到快速匹配的。

next 数组的值含义是：代表失配前的字符串中，有多大长度的相同的前缀后缀。比如 next[j] = k；表示 j 之前的字符串中有最大长度为 k 的公共前缀后缀。注意， next 表示的之 j 之前的字符串。

next 中表示的是长度，所以正好也就是前缀中最后一个字符的后一个字符的下标。当不匹配时，则后缀肯定已经匹配了，所以前缀不用比较，也就是移动前缀的后一个字符，正好也就是 next 的值，也就是 T[i] 与 P[next[j]] 来对齐。next[0] = -1 ，这个是标记。

假设：主串 T(Text) ，下标 i ；模式串 P(Pattern) ，下标 j 。

算法步骤：

```c
i       0  1  2  3  4  5  6  7  8  9
T       A  A  B  A  B  A  A  A  B  A            i = 0 + 2 (起始位置 + 移动了多少， 后同，不再写)
        √  √  ×
j       0  1  2  3  4
P       A  A  A  B  A                           j = 0 + 2
next   -1  0  1  2  0
```

在 i = 2 , j = 2 的时候失配， i 不变， j 移动到 j 之前最长公共前缀后缀中前缀的下一个字符，j = next[j] = 1 。 i = 2, j = 1 继续往后匹配。

```c
i       0  1  2  3  4  5  6  7  8  9
T       A  A  B  A  B  A  A  A  B  A            i = 2 + 0
           √  ×
j          0  1  2  3  4
P          A  A  A  B  A                        j = 1 + 0
next      -1  0  1  2  0
```

在 i = 2 , j = 1 的时候失配， i 不变， j 移动到 j 之前最长公共前缀后缀中前缀的下一个字符，j = next[j] = 0 。 i = 2, j = 0 继续往后匹配。

```c
i       0  1  2  3  4  5  6  7  8  9
T       A  A  B  A  B  A  A  A  B  A            i = 2 + 0
              ×
j             0  1  2  3  4
P             A  A  A  B  A                     j = 0 + 0
next         -1  0  1  2  0
```

在 i = 2 , j = 0 的时候失配， i 不变， j 移动到 j 之前最长公共前缀后缀中前缀的下一个字符，j = next[j] = -1 。 i++,j = 0 ,i = 3, j = 0 继续往后匹配。(注意，这里正好 i++,j++ 和匹配时候是一样的)。

```c
i       0  1  2  3  4  5  6  7  8  9
T       A  A  B  A  B  A  A  A  B  A            i = 3 + 1
                 √  ×
j                0  1  2  3  4
P                A  A  A  B  A                  j = 0 + 1
next            -1  0  1  2  0
```

在 i = 4 , j = 1 的时候失配， i 不变， j 移动到 j 之前最长公共前缀后缀中前缀的下一个字符，j = next[j] = 0 。 i = 4, j = 0 继续往后匹配。

```c
i       0  1  2  3  4  5  6  7  8  9
T       A  A  B  A  B  A  A  A  B  A            i = 4 + 0
                    ×
j                   0  1  2  3  4
P                   A  A  A  B  A               j = 0 + 0
next               -1  0  1  2  0
```

在 i = 4 , j = 0 的时候失配， i 不变， j 移动到 j 之前最长公共前缀后缀中前缀的下一个字符，j = next[j] = -1 。 i++,j = 0 ,i = 5, j = 0 继续往后匹配。

```c
i       0  1  2  3  4  5  6  7  8  9
T       A  A  B  A  B  A  A  A  B  A            i = 5 + 4
                       √  √  √  √  √
j                      0  1  2  3  4
P                      A  A  A  B  A            j = 0 + 4
next                  -1  0  1  2  0
```

在 i = 9 , j = 4 的时候匹配完成。

计算公共前后缀长度：
根据前一个子串的最长公共前后缀长度和当前的字符和之前最长公共前后缀后面一个字符（P[next[j]]）是否相等，可以算出当前的最长公共前后缀长度。

i 表示最长公共前后缀中后缀的下标
j 表示最长公共前后缀中前缀的下标，也正好是前一个最长公共前后缀的长度，也正好也就是前缀中最后一个字符的后一个字符的下标

1. 如果 j == -1 ，i++,j=0 (注意，这里正好 i++,j++ 和匹配时候是一样的)。
2. 如果 P[j] == P[next[j]] ， next[j] == j+1
3. 如果 P[j] != P[next[j]] ，比较其前缀的前缀

```c
A  A  A  B  A
next[0] = -1
i = 0,j = -1    : j == -1 所以 i++,j++,next[i]=j => i = 1,j = 0,next[1]=0
i = 1,j = 1     : next[i] == next[j]，所以 i++,j++,next[i]=j => i = 2,j = 1,next[2]=1
i = 2,j = 1     : next[i] == next[j]，所以 i++,j++,next[i]=j => i = 3,j = 2,next[3]=2
i = 3,j = 2     : next[i] != next[j], 所以 j = next[j] => j = next[2] = 1
i = 3,j = 1     : next[i] != next[j], 所以 j = next[j] => j = next[1] = 0
i = 3,j = 0     : next[i] != next[j], 所以 j = next[j] => j = next[0] = -1
i = 3,j = -1    : j == -1 所以 i++,j++,next[i]=j => i = 4,j = 0,next[4]=0

另外，当 P[j] = P[ next[j] ] 的时候可以优化下，因为 T[i] != P[j] 时，j = next[j] 然后继续比较
如果 P[j] = P[ next[j] ] ，T[i] != P[j] 仍然成立，所以在初始化 next 的时候就直接让其往前找，
让 next[i] = next[j] ，而不是 next[i] = j
```

### BM

BM后缀匹配算法构造了两个跳转表，分别叫做`坏后缀表`，和`好后缀表`。这两个表涉及了BM算法中的两个规则：

* 坏字符规则：
  * Case1: 坏字符没出现在模式串中，这时可以把模式串移动到坏字符的下一个字符，继续比较。
  * Case2: 坏字符出现在模式串中，这时可以把模式串最右边第一个出现的坏字符和主串的坏字符对齐，当然，这样可能造成模式串倒退移动，需要好后缀规则来调整。
* 好后缀规则：
  * Case1: 模式串中存在好后辍构成的子串，这种情况只需要将模式串中最靠右的子串与好后辍对齐即可。
  * Case2: 模式串中没有子串匹配上好后缀，但是存在与好后辍匹配的前辍，此时只需要将好后辍的后辍与模式串中对应的最长前辍匹配即可。
  * Case3: 模式串中没有子串匹配上好后缀，也不存在与好后辍匹配的前辍，则需要将模式串整个移动到目标串中不匹配位置之后即可。

算法步骤（假设：主串 T(Text) ，下标 i(i 为某一轮中主串 T 的起始位置，比较使用 T[i+j] ，这一轮中 i 不自增)；模式串 P(Pattern) ，下标 j 。）：

```c
在 i = 0,j=12 处开始比较

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                            ×
j       0  1  2  3  4  5  6  7  8  9 10 11 12
P       B  C  A  B  C  D  A  B  C  E  A  B  C
Bad     1  0  2  1  0  7  2  1  0  3  2  1  0
suffix  0  2  0  0  3  0  0  0  3  0  0  0 13
Good   11 11 11 11 11 11 11 11 11  4 11 13  1

在 i = 0 , j = 12 的时候失配
坏字符不在模式串中 Bad['F'] = 13 (坏字符 case1 )， 移动步数为： Bad['F'] - ((lenP - 1) - j) = 13 ；好后缀 Good[12] = 1
取其大者，移动 13 步， i = i + 13 = 13, j = 12 继续比较。

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                                                                   ×
j                                              0  1  2  3  4  5  6  7  8  9 10 11 12
P                                              B  C  A  B  C  D  A  B  C  E  A  B  C
Bad                                            1  0  2  1  0  7  2  1  0  3  2  1  0
suffix                                         0  2  0  0  3  0  0  0  3  0  0  0 13
Good                                          11 11 11 11 11 11 11 11 11  4 11 13  1

在 i = 13 , j = 12 的时候失配
坏字符在模式串中 Bad['B'] = 1 (坏字符 case2 )， 移动步数为： Bad['B'] - ((lenP - 1) - j) = 1 ；好后缀 Good[12] = 1
取其大者，移动 1 步， i = i + 1 = 14, j = 12 继续比较。

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                                                             ×  √  √  √
j                                                 0  1  2  3  4  5  6  7  8  9 10 11 12
P                                                 B  C  A  B  C  D  A  B  C  E  A  B  C
Bad                                               1  0  2  1  0  7  2  1  0  3  2  1  0
suffix                                            0  2  0  0  3  0  0  0  3  0  0  0 13
Good                                             11 11 11 11 11 11 11 11 11  4 11 13  1

在 i = 14 , j = 9 的时候失配
坏字符在模式串中 Bad['A'] = 2 (坏字符 case2 )， 移动步数为： Bad['A'] - ((lenP - 1) - j) = -1 ；好后缀 Good[12] = 4 (case1 存在2个好后缀，取右边的)
取其大者，移动 4 步， i = i + 11 = 18, j = 12 继续比较。

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                                                                            ×  √  √
j                                                             0  1  2  3  4  5  6  7  8  9 10 11 12
P                                                             B  C  A  B  C  D  A  B  C  E  A  B  C
Bad                                                           1  0  2  1  0  7  2  1  0  3  2  1  0
suffix                                                        0  2  0  0  3  0  0  0  3  0  0  0 13
Good                                                         11 11 11 11 11 11 11 11 11  4 11 13  1

在 i = 18 , j = 10 的时候失配
坏字符在模式串中 Bad['C'] = 0 (坏字符 case2 )， 移动步数为： Bad['C'] - ((lenP - 1) - j) = -2 ；好后缀 Good[10] = 11 (case2 存在好后缀的前缀 BC 匹配)
取其大者，移动 11 步， i = i + 1 = 29, j = 12 继续比较。

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                                                                                                                ×  √
j                                                                                              0  1  2  3  4  5  6  7  8  9 10 11 12
P                                                                                              B  C  A  B  C  D  A  B  C  E  A  B  C
Bad                                                                                            1  0  2  1  0  7  2  1  0  3  2  1  0
suffix                                                                                         0  2  0  0  3  0  0  0  3  0  0  0 13
Good                                                                                          11 11 11 11 11 11 11 11 11  4 11 13  1

在 i = 29 , j = 11 的时候失配
坏字符在模式串中 Bad['D'] = 0 (坏字符 case2 )， 移动步数为： Bad['D'] - ((lenP - 1) - j) = 6 ；好后缀 Good[11] = 13 (case3 不存在好后缀，也不存在好后缀匹配的前缀)
取其大者，移动 11 步， i = i + 13 = 42, j = 12 继续比较。

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                                                                                                                      √  √  √  √  √  √  √  √  √  √  √  √  √
j                                                                                                                                     0  1  2  3  4  5  6  7  8  9 10 11 12
P                                                                                                                                     B  C  A  B  C  D  A  B  C  E  A  B  C
Bad                                                                                                                                   1  0  2  1  0  7  2  1  0  3  2  1  0
suffix                                                                                                                                0  2  0  0  3  0  0  0  3  0  0  0 13
Good                                                                                                                                 11 11 11 11 11 11 11 11 11  4 11 13  1

在 i = 42 , j = 0 的时候匹配成功

```

### Horspool

Horspool 是 BM 算法中后缀匹配的一个改进。如果不匹配时，取主串 T 中与匹配串最后一个字符对应的字符，计算这个字符在模式串最右边的位置来移动。

算法步骤（假设：主串 T(Text) ，下标 i(i 为某一轮中主串 T 的起始位置，比较使用 T[i+j] ，这一轮中 i 不自增)；模式串 P(Pattern) ，下标 j 。）：

```c
在 i = 0,j=12 处开始比较

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                            ×
j       0  1  2  3  4  5  6  7  8  9 10 11 12
P       B  C  A  B  C  D  A  B  C  E  A  B  C
shift   1  4  2  1  4  7  2  1  4  3  2  1

在 i = 0 , j = 12 的时候失配，取最后一个字符 T[i+lenP-1] ， shift['F'] = 13 , i = i + 13 = 13, j = 12 继续比较。

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                                                                   ×
j                                              0  1  2  3  4  5  6  7  8  9 10 11 12
P                                              B  C  A  B  C  D  A  B  C  E  A  B  C
shift                                          1  4  2  1  4  7  2  1  4  3  2  1

在 i = 13 , j = 12 的时候失配，取最后一个字符 T[i+lenP-1] ， shift['B'] = 1 , i = i + 1 = 14, j = 12 继续比较。

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                                                             ×  √  √  √
j                                                 0  1  2  3  4  5  6  7  8  9 10 11 12
P                                                 B  C  A  B  C  D  A  B  C  E  A  B  C
shift                                             1  4  2  1  4  7  2  1  4  3  2  1

在 i = 14 , j = 9 的时候失配，取最后一个字符 T[i+lenP-1] ， shift['C'] = 4 , i = i + 4 = 18, j = 12 继续比较。

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                                                                            ×  √  √
j                                                             0  1  2  3  4  5  6  7  8  9 10 11 12
P                                                             B  C  A  B  C  D  A  B  C  E  A  B  C
shift                                                         1  4  2  1  4  7  2  1  4  3  2  1

在 i = 18 , j = 10 的时候失配，取最后一个字符 T[i+lenP-1] ， shift['C'] = 4 , i = i + 4 = 22, j = 12 继续比较。

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                                                                            ×  √  √  √  √  √  √
j                                                                         0  1  2  3  4  5  6  7  8  9 10 11 12
P                                                                         B  C  A  B  C  D  A  B  C  E  A  B  C
shift                                                                     1  4  2  1  4  7  2  1  4  3  2  1

在 i = 22 , j = 6 的时候失配，取最后一个字符 T[i+lenP-1] ， shift['C'] = 4 , i = i + 4 = 26, j = 12 继续比较。

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                                                                                                          ×
j                                                                                     0  1  2  3  4  5  6  7  8  9 10 11 12
P                                                                                     B  C  A  B  C  D  A  B  C  E  A  B  C
shift                                                                                 1  4  2  1  4  7  2  1  4  3  2  1

在 i = 26 , j = 12 的时候失配，取最后一个字符 T[i+lenP-1] ， shift['B'] = 1 , i = i + 1 = 27, j = 12 继续比较。

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                                                                                                    ×  √  √  √
j                                                                                        0  1  2  3  4  5  6  7  8  9 10 11 12
P                                                                                        B  C  A  B  C  D  A  B  C  E  A  B  C
shift                                                                                    1  4  2  1  4  7  2  1  4  3  2  1

在 i = 27 , j = 9 的时候失配，取最后一个字符 T[i+lenP-1] ， shift['C'] = 4 , i = i + 4 = 31, j = 12 继续比较。

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                                                                                                                ×  √  √  √
j                                                                                                    0  1  2  3  4  5  6  7  8  9 10 11 12
P                                                                                                    B  C  A  B  C  D  A  B  C  E  A  B  C
shift                                                                                                1  4  2  1  4  7  2  1  4  3  2  1

在 i = 31 , j = 9 的时候失配，取最后一个字符 T[i+lenP-1] ， shift['C'] = 4 , i = i + 4 = 35, j = 12 继续比较。

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                                                                                                                                     ×
j                                                                                                                0  1  2  3  4  5  6  7  8  9 10 11 12
P                                                                                                                B  C  A  B  C  D  A  B  C  E  A  B  C
shift                                                                                                            1  4  2  1  4  7  2  1  4  3  2  1

在 i = 35 , j = 12 的时候失配，取最后一个字符 T[i+lenP-1] ， shift['D'] = 7 , i = i + 7 = 42, j = 12 继续比较。

i       0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56
T       B  C  A  B  C  D  A  B  C  E  A  B  F  B  C  A  B  C  D  A  B  C  E  A  A  B  C  B  C  B  C  E  A  B  C  B  C  A  B  C  D  C  B  C  A  B  C  D  A  B  C  E  A  B  C  A  A
                                                                                                                 √  √  √  √  √  √  √  √  √  √  √  √  √
j                                                                                                                0  1  2  3  4  5  6  7  8  9 10 11 12
P                                                                                                                B  C  A  B  C  D  A  B  C  E  A  B  C
shift                                                                                                            1  4  2  1  4  7  2  1  4  3  2  1

在 i = 42 , j = 0 的时候匹配成功

```

### Sundy

Sunday算法是Daniel M.Sunday于1990年提出的字符串模式匹配。假设：主串 T(Text) ，下标 i(i 为某一轮中主串 T 的起始位置，比较使用 T[i+j] ，这一轮中 i 不自增)；模式串 P(Pattern) ，下标 j 。Sunday算法的思想是：

1. 从主串 T 的第 i 个字符开始，和模式串 T 的第 j 个字符进行比较 T[i+j] = P[j]，若相等，则主串和模式串都后移一个字符继续比较；
2. 若不相同，则将主串参与匹配的最末位字符的后一个字符 T[i+len(P)] 与模式串逆着查找。
3. 若模式串没有该字符，则模式串直接跳过，即移动位数 = 匹配串长度 + 1 ， i += len(P) + 1 。
4. 若模式串没该字符，则模式串中相同字符移动到主串该字符下，与该字符对齐。其移动位数 = 模式串中最右端的该字符到末尾的距离+1 ，i += (len(P) - j - 1) + 1 = len(P) - j 。

注意： 第 3 步中，也可以看成是 i += len(P) -(-1) = len(P) + 1 (j = -1)

算法步骤：

```c
i       0  1  2  3  4  5  6  7  8  9  10 11 12
T       c  b  a  b  d  c  b  a  c  b  b  a  d           i = 0
        √  √  ×
P       c  b  b  a
```

在 i = 0 , j = 2 的时候失配，后一个字符 i + len(P) = 4 的 d 不在 P 中(j = -1)，直接跳过， 将 T 移动到 i = i + len(P) + 1 = 5 处继续比较。

```c
i       0  1  2  3  4  5  6  7  8  9  10 11 12
T       c  b  a  b  d  c  b  a  c  b  b  a  d           i = 5
                       √  √  ×
T                      c  b  b  a
```

在 i = 5 , j = 2 的时候失配，后一个字符 i + len(P) = 9 的 b 在 P 中(j = 2)，T[9] 与 P[2] 对齐， 将 T 移动到 i = i + len(P) - j = 7 处继续比较。

```c
i       0  1  2  3  4  5  6  7  8  9  10 11 12
T       c  b  a  b  d  c  b  a  c  b  b  a  d           i = 7
                             ×
T                            c  b  b  a
```

在 i = 7 , j = 0 的时候失配，后一个字符 i + len(P) = 11 的 a 在 T 中(j = 3)，T[11] 与 P[3] 对齐， 将 T 移动到 i = i + len(P) - j = 8 处继续比较。

```c
i       0  1  2  3  4  5  6  7  8  9  10 11 12
T       c  b  a  b  d  c  b  a  c  b  b  a  d
                                √  √  √  √
T                               c  b  b  a
```

此时匹配完成。

### Rabin Karp

这里的实现，参考的 GO 源码 src/strings/strings.go 中的实现。这里不展开，在源码注释中详细解释了。

计算主串 T 中与 P 相同长度的字符串的 hash 值，如果相等再判断是否相等。使用滑动窗口计算 hash 值，调整 hash 值：

```go
primeRK = 16777619      // 质数 FNV hash
hash *= primeRK         // 相当于进位， primeRK 进制
hash += T[i]            // 添加头字符
hash -= T[i-lenP]       // 去掉尾字符
```

### Go 源码实现

go 源码 src/strings/strings.go 中的实现。这里不展开，在源码注释中详细解释了。
