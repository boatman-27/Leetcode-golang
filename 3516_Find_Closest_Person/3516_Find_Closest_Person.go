package main

import "math"

func findClosest(x int, y int, z int) int {
	distX := math.Abs(float64(x - z))
	distY := math.Abs(float64(y - z))

	if distX == distY {
		return 0
	}
	if distX < distY {
		return 1
	} else {
		return 2
	}
}
