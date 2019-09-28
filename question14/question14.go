package main

import "fmt"

func main() {
	biggestChainLength := 0
	start := 0
	for i := 2; i < 1000000; i++ {
		if len(collatz(i)) > biggestChainLength {
			biggestChainLength = len(collatz(i))
			start = i
		}
	}
	fmt.Printf("The starting number that produces the longest chain under 1 million is %v", start)
}

func collatz(num int) []int {
	sequence := []int{num}

	for {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = (3 * num) + 1
		}
		sequence = append(sequence, num)
		if num == 1 {
			break
		}
	}
	return sequence
}
