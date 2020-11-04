# [57. 插入区间](https://leetcode-cn.com/problems/insert-interval/)

## 题目

给出一个无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

示例 1：

```c
输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]
```

示例 2：

```c
输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
```

## 解题思路

```golang
func insert(intervals [][]int, newInterval []int) (ans [][]int) {
  // 和 intervals 中每一个 interval 存在三种关系
  // 1. 没有交集，interval 在 newInterval 右边
  // 2. 没有交集，interval 在 newInterval 左边
  // 3. 否则，有交集，这里计算并集
}
```
