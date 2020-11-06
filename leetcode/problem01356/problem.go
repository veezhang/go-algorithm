package problem01356

import (
	"sort"
)

func sortByBitsInterval(arr []int, bitCount func(int) int) []int {
	sort.Slice(arr, func(i, j int) bool {
		ni, nj := arr[i], arr[j]
		ci, cj := bitCount(ni), bitCount(nj)
		return ci < cj || ci == cj && ni < nj
	})
	return arr
}

func sortByBits1(arr []int) []int {
	return sortByBitsInterval(arr, func(num int) (count int) {
		for b := 1; b <= 1e4; b <<= 1 {
			if 0 != num&b {
				count++
			}
		}
		return
	})
}

func sortByBits2(arr []int) []int {
	return sortByBitsInterval(arr, func(num int) (count int) {
		for ; 0 != num; num >>= 1 {
			count += num & 1
		}
		return
	})
}

func sortByBits3(arr []int) []int {
	return sortByBitsInterval(arr, func(num int) (count int) {
		// 每次消去最右边的 1，直到消完为止
		for ; num != 0; num &= (num - 1) {
			count++
		}
		return
	})
}

var bit = [1e4 + 1]int{}

func init() {
	for i := range bit {
		// 递推
		bit[i] = bit[i>>1] + i&1
	}
}

func sortByBits4(arr []int) []int {
	return sortByBitsInterval(arr, func(num int) int {
		return bit[num]
	})
}

func sortByBits5(arr []int) []int {
	bitTable := [256]int{
		0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
		1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
		1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
		2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
		1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
		2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
		2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
		3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
		1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
		2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
		2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
		3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
		2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
		3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
		3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
		4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8,
	}
	return sortByBitsInterval(arr, func(num int) int {
		return bitTable[num&0xff] + bitTable[(num>>8)&0xff] + bitTable[(num>>16)&0xff] + bitTable[(num>>24)&0xff]
	})
}

func sortByBits6(arr []int) []int {
	return sortByBitsInterval(arr, func(num int) int {
		num = ((num >> 1) & 0x55555555) + (num & 0x55555555)
		num = ((num >> 2) & 0x33333333) + (num & 0x33333333)
		num = (((num >> 4) & 0x0f0f0f0f) + (num & 0x0f0f0f0f))
		num = (((num >> 8) & 0x00ff00ff) + (num & 0x00ff00ff))
		num = (((num >> 16) & 0x0000ffff) + (num & 0x0000ffff))
		return num
	})
}

func sortByBits7(arr []int) []int {
	return sortByBitsInterval(arr, func(num int) int {
		num = num&0x55555555 + (num>>1)&0x55555555
		num = num&0x33333333 + (num>>2)&0x33333333
		num = num&0x0f0f0f0f + (num>>4)&0x0f0f0f0f
		num = num&0x00ff00ff + (num>>8)&0x00ff00ff
		num = num&0x0000ffff + (num>>16)&0x0000ffff
		return num & 0x3f
	})
}

// 其他二进制参考 problem70015
