package main

import "fmt"

func main() {

	amicable := []int{}

	for i := 0; i < 10000; i++ {
		dofI := d(i)
		if d(dofI) == i && i != dofI && !contains(amicable, dofI) {
			amicable = append(amicable, i, dofI)
		}
	}

	fmt.Printf("The sum of all the amicable numbers under 10000 is %v", sliceSum(amicable))

}

func contains(s []int, e int) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == e {
			return true
		}
	}
	return false
}

func d(num int) int {

	sum := 0

	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sum = sum + i
		}
	}
	return sum
}

func sliceSum(s []int) int {
	sum := 0
	for _, element := range s {
		sum = sum + element
	}
	return sum
}
