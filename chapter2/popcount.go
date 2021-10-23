package main

import "fmt"

// pc[i] is the population of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var result = 0

	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}

	return result
}

func main() {
	fmt.Println(PopCount(69))
}

