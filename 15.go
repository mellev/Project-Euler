package main

import "fmt"

func main() {
	const t = 20
	grid := [t + 1][t + 1]int{}

	grid[0][0] = 1

	for x := 1; x < (t+1)*2; x++ {
		y := 0
		z := x

		if x > t+1 {
			y = x - t - 1
		}

		if x >= t+1 {
			z = t
		}

		for z != -1 && y != t+1 {
			if z == 0 {
				grid[z][y] = grid[0][y-1]
			} else if y == 0 {
				grid[z][y] = grid[z-1][0]
			} else {
				grid[z][y] = grid[z][y-1] + grid[z-1][y]
			}

			y++
			z--
		}
	}

	fmt.Println(grid[t][t])
}
