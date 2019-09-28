package main

import "fmt"

func main() {
	maxLimit := 1000
	multiplesSum := 0
	for i := 0; i < maxLimit; i++ {
		if i%3 == 0 || i%5 == 0 {
			multiplesSum = multiplesSum + i
		}
	}
	fmt.Println(multiplesSum)

}
