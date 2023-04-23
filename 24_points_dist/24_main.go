package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := &point{x: 0, y: 4}
	p2 := &point{x: 3, y: 0}
	fmt.Println(p1.distanceTo(p2))
}

type point struct {
	x, y int
}

func (p1 *point) distanceTo(p2 *point) float64 {
	xVec := p2.x - p1.x
	yVec := p2.y - p1.y
	dist := math.Sqrt(float64(xVec*xVec + yVec*yVec))
	return dist
}
