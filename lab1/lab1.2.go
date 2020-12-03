package main

import (
	"bytes"
	"math"
)

type resultStruct struct {
	idx    float64
	key    string
	result string
}

type resultSliceStruct []resultStruct

func stringRepeatBy(secret []byte, key int) []byte {
	res := []byte{}
	for i := 0; i < len(secret); i += key {
		res = append(res, secret[i])
	}
	return res
}

func indexOfCoincidence(secret []byte) float64 {
	counter := map[byte]int{}
	for _, c := range secret {
		counter[c]++
	}

	sum := 0
	for _, val := range counter {
		sum += val * (val - 1)
	}

	return float64(sum) / float64(len(secret)*(len(secret)-1))
}

func nextProduct(a []byte, r int) func() []byte {
	p := make([]byte, r)
	x := make([]int, len(p))
	return func() []byte {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = a[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(a) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return p
	}
}

func vigenereWithoutDest(secret []byte, key []byte) []byte {
	dest := make([]byte, len(secret))
	vigenere(dest, secret, key)
	return dest
}

func vigenere(dest []byte, secret []byte, key []byte) {
	for i, c := range secret {
		dest[i] = c ^ key[i%len(key)]
	}
}

func chiSquare(text []byte, letter string, tableFrq map[string]float64) float64 {
	lowerText := bytes.ToLower(text)
	expectedFrq := tableFrq[letter] / float64(len(lowerText))
	letterCount := float64(bytes.Count(lowerText, []byte(letter)))
	return (math.Pow(letterCount-expectedFrq, 2.)) / expectedFrq
}
