package main

import "fmt"

func main() {

	fmt.Println(squaresOfTheSum(100) - sumOfTheSquares(100))

}

func sumOfTheSquares(upperLimit int) int {
	sum := 0
	for i := 0; i <= upperLimit; i++ {
		sum = sum + (i * i)
	}
	return sum
}

func squaresOfTheSum(upperLimit int) int {
	sum := 0
	for i := 0; i <= upperLimit; i++ {
		sum = sum + i
	}
	return sum * sum
}
