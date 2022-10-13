package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	result := lookAndSay(100)
	fmt.Println(result)
}
func lookAndSay(n int) int {
	if n == 1 {
		return 1
	}

	var seq []int
	seq = append(seq, 1)
	type chunkSeqResult struct {
		index int
		seq   []int
	}
	for i := 2; i <= n; i++ {
		// if the sequence is large, break the sequence into chunks
		if len(seq) > 1000 {

			chunks := chunkSlice(seq, 400)
			// all next sequßences from chunk with index
			var wg sync.WaitGroup
			nextSequences := make([]chunkSeqResult, len(chunks))

			for j, chunk := range chunks {
				wg.Add(1)
				go func(index int, v []int) {

					//find next sequence from the chunks, keep track of the index
					newseq := nextLookAndSaySeq(v)
					chunkSeq := chunkSeqResult{
						index: index,
						seq:   newseq,
					}

					nextSequences[index] = chunkSeq

					defer wg.Done()
				}(j, chunk)
			}
			//wait for go routin to finish
			wg.Wait()

			// sort nextSequences by index
			sort.Slice(nextSequences, func(a, b int) bool {
				return nextSequences[a].index < nextSequences[b].index
			})
			combinedSeq := []int{}
			previousSeq := []int{}
			//compare the last number of previous sequence with the second number of the next sequence.
			for _, nexseq := range nextSequences {
				// combine the descriptor
				if len(previousSeq) > 0 && nexseq.seq[1] == previousSeq[len(previousSeq)-1] {
					combinedSeq[len(combinedSeq)-2] = previousSeq[len(previousSeq)-2] + nexseq.seq[0]
					// remove the first two next sequence
					combinedSeq = append(combinedSeq, nexseq.seq[2:]...)
				} else if len(previousSeq) == 0 || nexseq.seq[1] != previousSeq[len(previousSeq)-1] {
					combinedSeq = append(combinedSeq, nexseq.seq...)
				}

				previousSeq = combinedSeq
			}

			seq = combinedSeq
			// fmt.Println(i, "*****final******** ", seq)
		} else {
			seq = nextLookAndSaySeq(seq)
		}
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
func chunkSlice(slice []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
