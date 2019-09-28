package main

import (
	"fmt"
	"math"
)

func main() {

	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}

	m := make(map[int]int)

	for i := 0; i <= 7; i++ {
		m[primes[i]] = 0
	}

	for j := 1; j <= 20; j++ {
		for i := range primes {
			if primeFact(j)[primes[i]] > m[primes[i]] {
				m[primes[i]] = primeFact(j)[primes[i]]
			}
		}
	}

	var finalval float64 = 1
	for k := range primes {
		onek := (math.Pow(float64(primes[k]), float64(m[primes[k]])))
		finalval = finalval * onek
	}
	fmt.Println(finalval)

}

func primeFact(n int) map[int]int {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}

	m := make(map[int]int)

	for i := 0; i <= 7; i++ {
		m[primes[i]] = 0
	}

OuterLoop:
	for i := range primes {
		counter := 0
	InnerLoop:
		for {
			if n%primes[i] != 0 {
				break InnerLoop
			} else {
				n = n / primes[i]
				if n == 1 {
					counter = counter + 1
					m[primes[i]] = counter
					break OuterLoop
				}
			}
			counter = counter + 1
			m[primes[i]] = counter
		}
	}
	return m

}
