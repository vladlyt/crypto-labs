package main

type MTimproved struct {
	inputs []uint32
}

func rightShift(input uint32, num uint32) uint32 {
	var result uint32
	result = num
	for i := 0; i < 32; i++ {
		result = input ^ result>>num
	}
	return result
}

func leftShift(input uint32, num uint32, bitmask uint32) uint32 {
	var result uint32
	result = num
	for i := 0; i < 32; i++ {
		result = input ^ (result << num & bitmask)
	}
	return result
}

func initMTImproved(inputs []uint32) *MTimproved {
	return &MTimproved{
		inputs: inputs,
	}
}

func (mt *MTimproved) Backtrack() *MT19937 {
	for i, n := range mt.inputs {
		mt.inputs[i] = mt.Unstep(n)
	}
	newMT := initMT19937()
	newMT.state = mt.inputs[:]
	newMT.index = 0
	return newMT
}

func (mt *MTimproved) MakeRange() *MT19937 {
	newMT := mt.Backtrack()
	for i := 0; i < 624; i++ {
		newMT.Next()
	}
	return newMT
}

func (mt *MTimproved) Unstep(n uint32) uint32 {
	var result uint32
	result = n
	result = rightShift(result, 18)
	result = leftShift(result, 15, 0xefc60000)
	result = leftShift(result, 7, 0x9d2c5680)
	result = rightShift(result, 11)
	return result
}
