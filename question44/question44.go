package main

import (
	"fmt"
	"math"
)

func main() {

Loop:
	for j := 1; j < 1021; j++ {
		for k := 1; k < 2168; k++ {
			query1 := pentagonal(j) + pentagonal(k)
			query2 := pentagonal(k) - pentagonal(j)
			if isPentagonal(query1) == true && isPentagonal(query2) == true {
				fmt.Printf("The value of Pj is %d", j)
				fmt.Printf("\nThe value of Pk is %v", k)
				fmt.Printf("\nThe value of D is %d", pentagonal(k)-pentagonal(j))
				break Loop
			}

		}
	}

}

func isPentagonal(n int) bool {

	var temp = (1 + math.Sqrt(1+24*float64(n))) / 6
	if temp == float64(int(temp)) {
		return true
	}
	return false
}

func pentagonal(n int) int {
	return n * ((3 * n) - 1) / 2
}
