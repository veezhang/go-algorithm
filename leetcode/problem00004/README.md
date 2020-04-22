# [4. Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/)

## 题目

There are two sorted arrays `nums1` and `nums2` of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume `nums1` and `nums2` cannot be both empty.

Example 1:

```c
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
```

Example 2:

```c
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
```

## 题目大意

给定两个大小为 m 和 n 的有序数组 `nums1` 和 `nums2`。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

中位数：将一个集合划分为两个长度相等的子集，其中一个子集中的元素总是大于另一个子集中的元素。

如果为奇数个，则为中间那个数；如果为偶数个，则为中间那两个数的均值。

## 解题思路

### K值法

如果 m+n为奇数，则是找第(m+n+1)/2个数；如果是偶数，则是找第(m+n)/2个数和第(m+n)/2+1个数，再取均值。

转换一下，偶数的时候(m+n+1)/2 = (m+n)/2，也就是我们只要求第(m+n+1)/2个数，如果是偶数，再找一下下一个最小的数，再取均值，而不用再找第(m+n)/2+1个数。

求最小K值：过滤掉两个数组中K/2中较小的，如此反复

递归有函数堆栈的开销，我们这里不采用递归

### 二分法

用两把刀切两个数组，要满足两把刀左边的全部小于其右边的，使得左边有(m+n+1)/2个数

用一把刀切其中一个，另一把刀的位置也就知道了，因为要切的总数固定

如果切到左边了（也就是切小了），再切右边的一半；如果切到右边了（也就是切大了），再切左边的一半；

为了切的次数少一些，我们向短一些的切

假设最终如下：

```c
// nums1[0] ... nums1[mid1-1] | nums1[mid1] ... nums1[m - 1]
// nums2[0] ... nums2[mid2-1] | nums2[mid2] ... nums2[n - 1]

```

// 那么左边总数为： ((mid1 - 1) - 0 + 1) + ((mid2 - 1) - 0 + 1) = mid1 + mid2 = (m+n+1)/2

// 所以 mid2 = (m+n+1)/2 - mid1

如果是奇数个，那么取 max(nums1[mid1-1],nums2[mid2-1])

如果是偶数个，那么取 (max(nums1[mid1-1],nums2[mid2-1]) + min(nums1[mid1],nums2[mid1]))/2

注意刀的边界问题，mid1取值范围是[0,m]，可以取值m，表示其左边都是小的数字；另外取值0或者m的时候，是已经可以算出来数字的
