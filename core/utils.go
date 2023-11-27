package core

import (
	"log"
)

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

func CheckErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func MaxOverArray(a []int) int {
	mx := MinInt
	for _, x := range a {
		if x > mx {
			mx = x
		}
	}
	return mx
}
