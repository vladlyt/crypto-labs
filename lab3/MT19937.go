package main

const (
	UPPER  uint32 = 0x80000000
	LOWER  uint32 = 0x7fffffff
	FVALUE uint32 = 1812433253
)

type MT19937 struct {
	state []uint32
	index int
}

func initMT19937() *MT19937 {
	return &MT19937{
		state: make([]uint32, 624),
		index: 625,
	}

}

func (mt *MT19937) Seed(seed uint32) {
	mt.state[0] = seed
	for i := uint32(1); i < 624; i++ {
		mt.state[i] = FVALUE*(mt.state[i-1]^(mt.state[i-1]>>30)) + i
	}
	mt.index = 624
}

func (mt *MT19937) Twist() {
	n := 624
	m := 397
	if mt.index == 625 {
		mt.Seed(5489)
	}
	for i := 0; i < n-m; i++ {
		x := (mt.state[i] & UPPER) | (mt.state[i+1] & LOWER)
		mt.state[i] = mt.state[i+397] ^ (x >> 1) ^ ((x & 1) * 0x9908b0df)
	}
	for i := n - m; i < n-1; i++ {
		x := (mt.state[i] & UPPER) | (mt.state[i+1] & LOWER)
		mt.state[i] = mt.state[i+(m-n)] ^ (x >> 1) ^ ((x & 1) * 0x9908b0df)
	}
	x := (mt.state[n-1] & UPPER) | (mt.state[0] & LOWER)
	mt.state[n-1] = mt.state[m-1] ^ (x >> 1) ^ ((x & 1) * 0x9908b0df)
	mt.index = 0
}

func (mt *MT19937) Next() uint32 {
	if mt.index >= 624 {
		mt.Twist()
	}
	x := mt.state[mt.index]
	x ^= x >> 11
	x ^= (x << 7) & 0x9d2c5680
	x ^= (x << 15) & 0xefc60000
	x ^= x >> 18
	mt.index++
	return x
}

func (mt *MT19937) mtToFloat() float64 {
	var res float64
	res = (float64(mt.Next()) + 0.5) * (1.0 / 4294967296.0)
	return res
}
