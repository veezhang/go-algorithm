package problem00127

func ladderLengthBFS1(beginWord string, endWord string, wordList []string) int {
	// wordMap word 对于的哈希map
	wordMap := make(map[string]struct{}, len(wordList)+1)

	for _, word := range wordList {
		wordMap[word] = struct{}{}
	}

	if _, ok := wordMap[endWord]; !ok {
		return 0
	}
	// 如果有 beginWord ，将其删除，避免重复计算
	delete(wordMap, beginWord)

	visited := make(map[string]struct{})

	// 队列
	queue := []string{beginWord}
	ans := 1 // beginWord 已经是第一个单词了
	visited[beginWord] = struct{}{}

	// 开始BFS
	for len(queue) > 0 {
		qlen := len(queue)
		ans++ // 接下来开始处理第 ans 个单词，如果找到了则返回 ans
		for i := 0; i < qlen; i++ {
			// 先进先出
			word := []byte(queue[0])
			queue = queue[1:]
			// 将 work 中每一个字符替换成 a-z 26个英文字母，看看 wordMap 中是否存在对应的
			for j, b := range word {
				for k := byte('a'); k <= byte('z'); k++ {
					if k == b {
						continue
					}
					word[j] = k
					nextWord := string(word)
					word[j] = b

					if _, ok := wordMap[nextWord]; ok {
						// 如果存在，且是 endWord 则找到了，返回
						if nextWord == endWord {
							return ans
						}
						// 如果没有遍历过，才加入队列
						if _, ok := visited[nextWord]; !ok {
							visited[nextWord] = struct{}{}
							queue = append(queue, nextWord)
						}
					}
				}
			}
		}
	}

	return 0
}

func ladderLengthBFS2(beginWord string, endWord string, wordList []string) int {
	// wordMap word 对于的哈希map
	wordMap := make(map[string]struct{}, len(wordList)+1)

	for _, word := range wordList {
		wordMap[word] = struct{}{}
	}

	if _, ok := wordMap[endWord]; !ok {
		return 0
	}
	// 如果有 beginWord ，将其删除，避免重复计算
	delete(wordMap, beginWord)

	// 双向BFS ，需要记录两边遍历过的数据
	visitedBegin := make(map[string]struct{})
	visitedEnd := make(map[string]struct{})

	// 双向BFS ，这里两个队列
	queueBegin := []string{beginWord}
	queueEnd := []string{endWord}
	ans := 1 // beginWord 已经是第一个单词了
	visitedBegin[beginWord] = struct{}{}
	visitedEnd[endWord] = struct{}{}

	// 开始双向BFS
	for len(queueBegin) > 0 && len(queueEnd) > 0 {
		// 去一个小点儿的队列来找，后面就一直操作 queueBegin
		if len(queueBegin) > len(queueEnd) {
			queueBegin, queueEnd = queueEnd, queueBegin
			visitedBegin, visitedEnd = visitedEnd, visitedBegin
		}

		qlen := len(queueBegin)
		ans++ // 接下来开始处理第 ans 个单词，如果找到了则返回 ans
		for i := 0; i < qlen; i++ {
			// 先进先出
			word := []byte(queueBegin[0])
			queueBegin = queueBegin[1:]
			// 将 work 中每一个字符替换成 a-z 26个英文字母，看看 wordMap 中是否存在对应的
			for j, b := range word {
				for k := byte('a'); k <= byte('z'); k++ {
					if k == b {
						continue
					}
					word[j] = k
					nextWord := string(word)
					word[j] = b

					if _, ok := wordMap[nextWord]; ok {
						// 如果另一边已经遍历过，则找到了，返回
						if _, ok := visitedEnd[nextWord]; ok {
							return ans
						}
						// 如果这边还没有遍历过，才加入队列
						if _, ok := visitedBegin[nextWord]; !ok {
							visitedBegin[nextWord] = struct{}{}
							queueBegin = append(queueBegin, nextWord)
						}
					}
				}
			}
		}
	}

	return 0
}

func ladderLengthBFS3(beginWord string, endWord string, wordList []string) int {
	// wordMap word 对于的哈希map

	wordMap := make(map[string]int, len(wordList)+1)
	graph := make([][]int, 0)

	// 添加节点，返回节点 id
	addWord := func(word string) int {
		id, ok := wordMap[word]
		if !ok {
			id = len(wordMap)
			wordMap[word] = id
			graph = append(graph, []int{})
		}
		return id
	}

	addEdge := func(word string) int {
		// 真实单词
		id := addWord(word)
		wordBytes := []byte(word)
		for i, b := range wordBytes {
			wordBytes[i] = '*'
			// 虚拟单词，带有 * 号的
			idVirtual := addWord(string(wordBytes))
			// 添加相互连接关系
			graph[id] = append(graph[id], idVirtual)
			graph[idVirtual] = append(graph[idVirtual], id)

			wordBytes[i] = b
		}
		return id
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, endWordExists := wordMap[endWord]
	// 不存在 endWord
	if !endWordExists {
		return 0
	}

	queue := []int{beginId}

	// beginWord 到每一个 word 的距离
	// 0 表示无穷远，不用初始化为一个很大的值，保存的是 距离 + 1 （为了将 0 预留出来）
	dist := make([]int, len(wordMap))
	dist[beginId] = 1
	// 因为有带 * 的中间 word ，则距离 = (dict[id] - 1)/2
	// 节点数 = 距离 + 1  =  (dict[id] - 1)/2 + 1 = (dict[id] + 1)/2
	// 如下：
	// word  距离  dict
	// hot   0		1
	// *ot   1 		2
	// dot 	 2		3
	// do*   3		4
	// dog   4		5

	for len(queue) > 0 {
		id := queue[0]
		queue = queue[1:]

		// 如果找到 endId 了，则返回
		if id == endId {
			return (dist[endId] + 1) >> 1
		}

		// 遍历 id 的边的另一端
		for _, idPeer := range graph[id] {
			// idPeer 点之前不可达，则更新下距离，并把其加入到队列中
			if dist[idPeer] == 0 {
				dist[idPeer] = dist[id] + 1
				queue = append(queue, idPeer)
			}
		}
	}

	return 0
}

func ladderLengthBFS4(beginWord string, endWord string, wordList []string) int {
	// wordMap word 对于的哈希map

	wordMap := make(map[string]int, len(wordList)+1)
	graph := make([][]int, 0)

	// 添加节点，返回节点 id
	addWord := func(word string) int {
		id, ok := wordMap[word]
		if !ok {
			id = len(wordMap)
			wordMap[word] = id
			graph = append(graph, []int{})
		}
		return id
	}

	addEdge := func(word string) int {
		// 真实单词
		id := addWord(word)
		wordBytes := []byte(word)
		for i, b := range wordBytes {
			wordBytes[i] = '*'
			// 虚拟单词，带有 * 号的
			idVirtual := addWord(string(wordBytes))
			// 添加相互连接关系
			graph[id] = append(graph[id], idVirtual)
			graph[idVirtual] = append(graph[idVirtual], id)

			wordBytes[i] = b
		}
		return id
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, endWordExists := wordMap[endWord]
	// 不存在 endWord
	if !endWordExists {
		return 0
	}

	queueBegin := []int{beginId}
	queueEnd := []int{endId}

	// beginWord 到每一个 word 的距离
	// 0 表示无穷远，不用初始化为一个很大的值，保存的是 距离 + 1 （为了将 0 预留出来）
	distBegin := make([]int, len(wordMap))
	distEnd := make([]int, len(wordMap))
	distBegin[beginId] = 1
	distEnd[endId] = 1
	// 因为有带 * 的中间 word ，则距离 = (dict[id] - 1)/2 = ((distBegin[id] - 1) + (distBegin[id] - 1))/2 = (distBegin[id]+distBegin[id])/2 - 1
	// 节点数 = 距离 + 1  = (distBegin[id]+distBegin[id])/2
	// 如下：
	// word  距离  dict
	// hot   0		1
	// *ot   1 		2
	// dot 	 2		3
	// do*   3		4
	// dog   4		5

	for len(queueBegin) > 0 && len(queueEnd) > 0 {
		// 去一个小点儿的队列来找，后面就一直操作 queueBegin
		if len(queueBegin) > len(queueEnd) {
			queueBegin, queueEnd = queueEnd, queueBegin
			distBegin, distEnd = distEnd, distBegin
		}

		id := queueBegin[0]
		queueBegin = queueBegin[1:]

		// 如果另一边已经遍历过，则找到了，返回
		if distEnd[id] != 0 {
			return (distBegin[id] + distEnd[id]) >> 1
		}

		// 遍历 id 的边的另一端
		for _, idPeer := range graph[id] {
			// idPeer 点之前不可达，则更新下距离，并把其加入到队列中
			if distBegin[idPeer] == 0 {
				distBegin[idPeer] = distBegin[id] + 1
				queueBegin = append(queueBegin, idPeer)
			}
		}
	}

	return 0
}
