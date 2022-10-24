package main

import (
	"fmt"
	"math/big"
)

func fibNumber() func(i int) *big.Int {
	a := big.NewInt(1)
	b := big.NewInt(0)
	return func(i int) *big.Int {
		//swap values and return new b
		b.Add(a, b)
		b, a = a, b
		return b
	}
}

func main() {
	f := fibNumber()
	i, fabN := 0, big.NewInt(0)

	var limit big.Int
	// smallest integer with 10000 digits
	limit.Exp(big.NewInt(10), big.NewInt(9999), nil)

	for fabN.Cmp(&limit) < 0 {
		fabN = f(i)
		i++
	}
	fmt.Println(i)
}
