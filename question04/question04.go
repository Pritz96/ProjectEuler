package main

import "fmt"

func main() {
	var largestPalindrome int
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			if isPalindrome(i*j) && i*j > largestPalindrome {
				largestPalindrome = i * j
			}
		}
	}
	fmt.Println(largestPalindrome)

}

func isPalindrome(integer int) bool {
	if integer == reverseNumber(integer) {
		return true
	}
	return false
}

func reverseNumber(num int) int {
	result := 0
	for num > 0 {
		remainder := num % 10
		result = (result * 10) + remainder
		num /= 10
	}
	return result
}
