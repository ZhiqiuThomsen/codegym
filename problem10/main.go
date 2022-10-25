package main

import (
	"fmt"
	"math/big"
)

func fibNumber() func(i int) *big.Int {
	a := big.NewInt(1)
	b := big.NewInt(0)
	return func(i int) *big.Int {
		//add and swap
		b, a = a, b.Add(a, b)
		return b
	}
}

func main() {
	f := fibNumber()
	i, fibN := 0, big.NewInt(0)

	var edge big.Int
	// smallest integer with 10000 digits
	edge.Exp(big.NewInt(10), big.NewInt(9999), nil)

	//loop until the fib number is larger than the edge number
	for fibN.Cmp(&edge) < 0 {
		fibN = f(i)
		i++
	}
	fmt.Println(i)
}
