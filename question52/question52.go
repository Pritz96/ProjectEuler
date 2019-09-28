package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

func main() {

	for i := 1; ; i++ {
		if sixLoop(i) == true {
			fmt.Printf("The smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits is %v", i)
			break
		}
	}
}

func numSplit(num int) []int {
	numLen := len(strconv.Itoa(num))
	digits := []int{}
	for i := 0; i < numLen; i++ {
		digits = append(digits, num%10)
		num = num / 10
	}
	sort.Ints(digits)
	return digits
}

func sixLoop(a int) bool {
	for i := 2; i < 7; i++ {
		if compare(numSplit(a), numSplit(a*i)) == false {
			return false
		}
	}
	return true
}

func compare(a []int, b []int) bool {
	return reflect.DeepEqual(a, b)
}
