package main

import (
	"fmt"
	"strconv"
)

func main() {

	firstWirePlot := []string{"R8", "U5", "L5", "D3"}
	secondWirePlot := []string{"U7", "R6", "D4", "L4"}

	w1 := WirePath{}
	w2 := WirePath{}
	w1.FollowPath(firstWirePlot)
	w2.FollowPath(secondWirePlot)
	//fmt.Println(firstWire[0])
	// for _, element := range firstWire {
	// 	fmt.Println(element)
	// }
}

// WirePath keeps track of a wire's path starting from 0,0.
type WirePath struct {
	start, end, max, min Point32
	path                 []Point32
}

// Point32 represents and x and y coordinate in 2D space using int32
type Point32 struct {
	x, y int32
}

// FollowPath accepts a path represented as strings and follows it from current x,y
func (r *WirePath) FollowPath(path []string) {

	fmt.Println(fmt.Sprintf("Starting coordinates: %d,%d", r.end.x, r.end.y))
	for _, element := range path {
		direction := element[:1]
		m, _ := strconv.ParseInt(element[1:len(element)], 10, 32)
		magnitude := int32(m)

		if direction == "U" {
			fmt.Println(fmt.Sprintf("Going up %d", magnitude))
			r.end.y += magnitude
			if r.end.y > r.max.y {
				r.max.y = r.end.y
			}
		} else if direction == "D" {
			fmt.Println(fmt.Sprintf("Going down %d", magnitude))
			r.end.y -= magnitude
			if r.end.y < r.min.y {
				r.min.y = r.end.y
			}
		} else if direction == "L" {
			fmt.Println(fmt.Sprintf("Going left %d", magnitude))
			r.end.x -= magnitude
			if r.end.x < r.min.x {
				r.min.x = r.end.x
			}
		} else if direction == "R" {
			fmt.Println(fmt.Sprintf("Going right %d", magnitude))
			r.end.x += magnitude
			if r.end.x > r.max.x {
				r.max.x = r.end.x
			}
		} else {
			panic(fmt.Sprintf("invalid direction %s", direction))
		}

		fmt.Println(fmt.Sprintf("%d,%d", r.end.x, r.end.y))
	}
	fmt.Println(fmt.Sprintf("Ending coordinates: %d,%d  [max x,y = %d,%d] [min x,y = %d,%d]", r.end.x, r.end.y, r.max.x, r.max.y, r.min.x, r.min.y))
}
