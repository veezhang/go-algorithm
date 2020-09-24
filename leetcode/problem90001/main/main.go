package main

import (
	"fmt"
	"math/rand"
	"time"

	"github/veezhang/go-algorithm/leetcode/problem90001"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// 测试不准确，随机数，只是大概的
	tests := []struct {
		MinO  int
		MaxO  int
		Min   int
		Max   int
		Count int
	}{
		// {
		// 	MinO:  1,
		// 	MaxO:  10,
		// 	Min:   10,
		// 	Max:   120,
		// 	Count: 100000000,
		// },
		{
			MinO:  1,
			MaxO:  7,
			Min:   1,
			Max:   5,
			Count: 100000000,
		},
		{
			MinO:  1,
			MaxO:  5,
			Min:   1,
			Max:   7,
			Count: 100000000,
		}, {
			MinO:  3,
			MaxO:  8,
			Min:   10,
			Max:   20,
			Count: 100000000,
		},
	}

	for _, t := range tests {
		countMap := make([]int, t.Max-t.Min+1)
		for i := 0; i < t.Count; i++ {
			rb := problem90001.Random{Min: t.MinO, Max: t.MaxO}
			r := rb.Random(t.Min, t.Max)
			if r-t.Min >= t.Max-t.Min+1 {
				println("r = ", r)
			}
			countMap[r-t.Min]++
		}
		fmt.Printf("[%d,%d]->[%d,%d]\n", t.MinO, t.MaxO, t.Min, t.Max)
		for i, v := range countMap {
			fmt.Printf("%d : %d\n", i+t.Min, v)
		}
	}
}
