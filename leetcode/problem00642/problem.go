package problem00860

func predictPartyVictory(senate string) string {
	n := len(senate)
	// radiant 为 R 的队列，dire 为 D 的队列
	var radiant, dire []int
	for i, s := range []byte(senate) {
		if 'R' == s {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	// 互相投票，直到某个队列完了
	// 每个投票都投对方接下来第一个 “禁止一名参议员的权利”
	for len(radiant) > 0 && len(dire) > 0 {
		// 谁先投票就投对方接下来第一个 “禁止一名参议员的权利”
		// 那么自己就 + n ，可以第二轮投票，对方就被永久禁止了
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+n)
		} else {
			dire = append(dire, radiant[0]+n)
		}
		// 第 0 个要么被禁止了，要么 + n 放到后面进行第二轮的选票
		radiant = radiant[1:]
		dire = dire[1:]
	}

	// 哪方还有，则为胜利者
	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
