package problem00204

func countPrimes1(n int) (count int) {
	// 第一个质数位 2
	if n <= 2 {
		return 0
	}

	isPrime := func(x int) bool {
		for i := 2; i*i <= x; i++ {
			if 0 == x%i {
				return false
			}
		}
		return true
	}

	count = 1 // 质数 2
	for i := 3; i < n; i += 2 {
		if isPrime(i) {
			count++
		}
	}
	return
}

// bitmap
type bitmap struct {
	n    int
	bits []byte
}

func newBitmap(n int) *bitmap {
	if n < 0 {
		panic("n must be non-negative number.")
	}
	return &bitmap{
		n:    n,
		bits: make([]byte, (n+7)>>3),
	}
}

func (b *bitmap) Add(x int) {
	if x < 0 || x > b.n {
		panic("Out of range")
	}
	b.bits[x>>3] |= 1 << (x & 0x7)
}

func (b *bitmap) AddMany(xs ...int) {
	for _, x := range xs {
		b.Add(x)
	}
}

func (b *bitmap) Contains(x int) bool {
	if x < 0 || x > b.n {
		panic("Out of range")
	}
	return 0 != b.bits[x>>3]&(1<<(x&0x7))
}

func countPrimes2(n int) (count int) {
	// 第一个质数位 2
	if n <= 2 {
		return 0
	}

	// bit 位为 0 表示质数， 非零表示不是质数
	b := newBitmap(n)

	for i := 2; i < n; i++ {
		if !b.Contains(i) {
			count++
			// 质数的倍数全部为非质数
			for j := 2 * i; j < n; j += i {
				b.Add(j)
			}
		}
	}
	return
}
