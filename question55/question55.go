package main

import "fmt"

func main() {
	count := 0
	for i := 0; i < 10000; i++ {
		if isNotLychrel(i) == false {
			count++
		}
	}
	fmt.Printf("There are %d Lychrel numbers below ten-thousand\n", count)

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

func isNotLychrel(num int) bool {
	i := 1
	for i < 50 {
		if isPalindrome(num + reverseNumber(num)) {
			return true
		}
		num = num + reverseNumber(num)
		i++
	}
	return false
}
