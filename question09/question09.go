package main

import (
	"fmt"
	"math"
)

func main() {
	squares := []int{}
	for i := 1; i < 1000; i++ {
		s := i * i
		squares = append(squares, s)
	}
All:
	for j := range squares {
		for k := range squares {
			cSquared := squares[j] + squares[k]
			if isInSlice(cSquared, squares) == true {
				if (math.Sqrt(float64(squares[j])) + math.Sqrt(float64(squares[k])) + math.Sqrt(float64(cSquared))) == 1000 {
					fmt.Printf("The pythagorean triple is a = %v, b = %v, c = %v, a+b+c =%v", math.Sqrt(float64(squares[j])), math.Sqrt(float64(squares[k])), math.Sqrt(float64(cSquared)), math.Sqrt(float64(squares[j]))+math.Sqrt(float64(squares[k]))+math.Sqrt(float64(cSquared)))
					fmt.Printf("\nThe product of a,b and c is abc=%v", math.Sqrt(float64(squares[j]))*math.Sqrt(float64(squares[k]))*math.Sqrt(float64(cSquared)))
					break All
				}
			}
		}
	}
}

func isInSlice(num int, s []int) bool {
	for i := range s {
		if num == s[i] {
			return true
		}
	}
	return false
}
