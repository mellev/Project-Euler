package main

import "fmt"
import "math"

func main() {
	var primes []int64
	var delimiters []int64
	var n int64 = 600851475143

Loop:
	for number := int64(2); float64(number) <= math.Ceil(math.Sqrt(float64(n))); number++ {
		for i := 0; i < len(primes); i++ {
			if number%primes[i] == 0 {
				continue Loop
			}
		}

		primes = append(primes, number)

		if n%number == 0 {
			delimiters = append(delimiters, number)
		}
	}

	fmt.Println(delimiters[len(delimiters)-1])
}
