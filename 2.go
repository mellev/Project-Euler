package main

import "fmt"

func main() {
	n1 := 1
	n2 := 2
	sum := n2
	fibo := 0

	for fibo <= 4000000 {
		if fibo%2 == 0 {
			sum += fibo
		}

		fibo = n1 + n2
		n1 = n2
		n2 = fibo
	}

	fmt.Println(sum)
}
