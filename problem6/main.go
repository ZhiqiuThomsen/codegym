package main

import (
	"fmt"
)

func main() {
	result := lookAndSay(20)
	fmt.Println(result)
}
func lookAndSay(n int) int {
	if n == 1 {
		return 1
	}

	var seq []int
	seq = append(seq, 1)

	for i := 2; i <= n; i++ {
		seq = nextLookAndSaySeq(seq)
		fmt.Println(i, " line:", seq)
	}

	return sum(seq)
}

func nextLookAndSaySeq(seq []int) []int {
	describor := 0
	number := seq[0]
	var nextSeq []int
	for index, decribed := range seq {

		if number == decribed {
			describor++

		} else {
			//add to nextSeq
			nextSeq = append(nextSeq, describor)
			nextSeq = append(nextSeq, number)

			//reset number and describor
			number = decribed
			describor = 1
		}
		// last number
		if index == len(seq)-1 {
			nextSeq = append(nextSeq, describor)
			nextSeq = append(nextSeq, number)
		}
	}

	return nextSeq
}
func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
