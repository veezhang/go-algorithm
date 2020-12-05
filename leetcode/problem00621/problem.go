package problem00621

func leastInterval(tasks []byte, n int) int {
	length := len(tasks)
	if length == 1 {
		// 只有一个任务，返回 1
		return 1
	}
	if n == 0 {
		// 没有间隔，返回长度
		return length
	}

	// 统计各个任务的个数
	mCount := make(map[byte]int, 0)
	for _, task := range tasks {
		mCount[task]++
	}

	if lenKind := len(mCount); 1 == lenKind {
		// 只有一种任务
		return (length-1)*(n+1) + 1
	}

	// nBucket 为桶的个数，最长的相同的任务决定桶的个数
	// nTaskOfLastBucket 为最后一个桶中的任务数，也就是：与最长相同任务数量相同的任务种类
	nBucket, nTaskOfLastBucket := 0, 0
	for _, count := range mCount {
		if count > nBucket {
			nBucket, nTaskOfLastBucket = count, 1
		} else if count == nBucket {
			nTaskOfLastBucket++
		}
	}

	if t := (nBucket-1)*(n+1) + nTaskOfLastBucket; t > length {
		return t
	}

	return length
}
