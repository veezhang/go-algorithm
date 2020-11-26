package problem00164

import (
	"math"

	"github.com/RoaringBitmap/roaring"
)

// bitmap
// 只能记录 [min,max] 直接的值
type bitmap struct {
	bits []byte
	min  int
	max  int
}

func newBitmap(min, max int) *bitmap {
	return &bitmap{
		bits: make([]byte, (max-min+1+7)>>3),
		min:  min,
		max:  max,
	}
}

func (b *bitmap) Add(x int) {
	if x < b.min && x > b.max {
		panic("Out of range")
	}
	x -= b.min
	b.bits[x>>3] |= 1 << (x & 0x7)
}

func (b *bitmap) AddMany(xs ...int) {
	for _, x := range xs {
		b.Add(x)
	}
}

func (b *bitmap) Contains(x int) bool {
	if x < b.min && x > b.max {
		panic("Out of range")
	}
	x -= b.min
	return 0 != b.bits[x>>3]&(1<<(x&0x7))
}

func maximumGap1(nums []int) (ans int) {
	// 小于等于一个数据，间隙为 0
	if len(nums) <= 1 {
		return
	}
	// 统计
	minNum, maxNum := math.MaxUint32, 0
	for _, n := range nums {
		if n < minNum {
			minNum = n
		}
		if n > maxNum {
			maxNum = n
		}
	}
	// 全部是同一个数
	if minNum == maxNum {
		return
	}

	b := newBitmap(minNum, maxNum)
	b.AddMany(nums...)

	// i 为第一个数，j 为第二个数
	for i, j := minNum, minNum+1; j <= maxNum; j++ {
		// 如果不包含，继续找
		if !b.Contains(j) {
			continue // 找到下一个数字
		}
		// 找到了
		// 计算间隙
		gap := j - i
		// 如果大于当前的，则更新
		if gap > ans {
			ans = gap
		}
		// 更新第一个数
		i = j
	}
	return
}

func maximumGap2(nums []int) (ans int) {
	// 小于等于一个数据，间隙为 0
	if len(nums) <= 1 {
		return
	}

	// 统计
	minNum, maxNum := math.MaxUint32, 0
	rb := roaring.NewBitmap()
	for _, n := range nums {
		rb.AddInt(n)
		if n < minNum {
			minNum = n
		}
		if n > maxNum {
			maxNum = n
		}
	}
	// 全部是同一个数
	if minNum == maxNum {
		return
	}

	// i 为第一个数，j 为第二个数
	it := rb.Iterator()
	for i := it.Next(); it.HasNext(); {
		j := it.Next()
		// 计算间隙
		gap := int(j - i)
		// 如果大于当前的，则更新
		if gap > ans {
			ans = gap
		}
		// 更新第一个数
		i = j
	}
	return
}

func maximumGap3(nums []int) (ans int) {
	// 小于等于一个数据，间隙为 0
	if len(nums) <= 1 {
		return
	}

	// 统计
	minNum, maxNum := math.MaxUint32, 0
	for _, n := range nums {
		if n < minNum {
			minNum = n
		}
		if n > maxNum {
			maxNum = n
		}
	}
	// 全部是同一个数
	if minNum == maxNum {
		return
	}

	// 分配桶的步长和个数是桶排序的关键
	// 在 n 个数下，形成的两两相邻区间是 n - 1 个，比如 [2,4,6,8] 这里
	// 有 4 个数，但是只有 3 个区间，[2,4], [4,6], [6,8]
	// 因此，桶步长 = 区间总长度 / 区间总个数 = (maxNum - minNum) / (len(nums) - 1)
	bucketStep := (maxNum - minNum) / (len(nums) - 1)
	if bucketStep < 1 {
		bucketStep = 1
	}

	// 上面得到了桶的步长，我们就可以以此来确定桶的个数
	// 桶个数 = 区间长度 / 桶长度
	// 这里考虑到实现的方便，多加了一个桶，为什么？
	// 还是举上面的例子，[2,4,6,8], 桶的步长 = (8 - 2) / (4 - 1) = 2
	//                             桶的个数 = (8 - 2) / 2 = 3
	// 已知一个元素，需要定位到桶的时候，一般是 (当前元素 - 最小值) / 桶步长
	// 这里其实利用了整数除不尽向下取整的性质
	// 但是上面的例子，如果当前元素是 8 的话 (8 - 2) / 2 = 3，对应到 3 号桶
	//                如果当前元素是 2 的话 (2 - 2) / 2 = 0，对应到 0 号桶
	// 你会发现我们有 0,1,2,3 号桶，实际用到的桶是 4 个，而不是 3 个
	// 透过例子应该很好理解，但是如果要说根本原因，其实是开闭区间的问题
	// 这里其实 0,1,2 号桶对应的区间是 [2,4),[4,6),[6,8)
	// 那 8 怎么办？多加一个桶呗，3 号桶对应区间 [8,10)
	bucketCount := (maxNum-minNum)/bucketStep + 1

	buckets := make([]struct {
		min int
		max int
	}, bucketCount)
	for i := range buckets {
		// -1 为标志，表示此桶的内部没有数据
		buckets[i].min = -1
	}

	// 计算桶内部的最小值，和最大值
	for _, n := range nums {
		bucketIndex := (n - minNum) / bucketStep
		if -1 == buckets[bucketIndex].min {
			buckets[bucketIndex].min = n
			buckets[bucketIndex].max = n
		} else {
			if n < buckets[bucketIndex].min {
				buckets[bucketIndex].min = n
			}
			if n > buckets[bucketIndex].max {
				buckets[bucketIndex].max = n
			}
		}
	}

	// 桶间计算最大间隙
	bucketPreMax := buckets[0].max
	for _, bucket := range buckets[1:] {
		if -1 == bucket.min {
			continue
		}
		gap := bucket.min - bucketPreMax
		// 如果大于当前的，则更新
		if gap > ans {
			ans = gap
		}
		bucketPreMax = bucket.max
	}

	return
}
