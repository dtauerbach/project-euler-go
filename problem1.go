// If we list all the natural numbers below 10 that are multiples
// of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import "fmt"

func main() {
	fmt.Printf("Result is: %d\n", dansum(3, 5, 10))
	fmt.Printf("Result is: %d\n", dansum(3, 5, 1000))
}

func dansum(firstDivisor int, secondDivisor int, total int) int {
	sum := 0
	for i := firstDivisor; i < total; i++ {
		if i%firstDivisor == 0 {
			sum += i
			continue
		} else if i%secondDivisor == 0 {
			sum += i
		}
	}
	return sum
}
