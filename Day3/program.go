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
	w1.SetPath(firstWirePlot)
	w2.SetPath(secondWirePlot)
	//fmt.Println(firstWire[0])
	// for _, element := range firstWire {
	// 	fmt.Println(element)
	// }
}

// WirePath keeps track of a wire's path
type WirePath struct {
	x, y, xMax, yMax, xMin, yMin int32
}

// SetPath accepts a path represented as strings and follows it from current x,y
func (r *WirePath) SetPath(path []string) {
	fmt.Println(fmt.Sprintf("Starting coordinates: %d,%d", r.x, r.y))
	for _, element := range path {
		direction := element[:1]
		m, _ := strconv.ParseInt(element[1:len(element)], 10, 32)
		magnitude := int32(m)

		if direction == "U" {
			fmt.Println(fmt.Sprintf("Going up %d", magnitude))
			r.y += magnitude
			if r.y > r.yMax {
				r.yMax = r.y
			}
		} else if direction == "D" {
			fmt.Println(fmt.Sprintf("Going down %d", magnitude))
			r.y -= magnitude
			if r.y < r.yMin {
				r.yMin = r.y
			}
		} else if direction == "L" {
			fmt.Println(fmt.Sprintf("Going left %d", magnitude))
			r.x -= magnitude
			if r.x < r.xMin {
				r.xMin = r.x
			}
		} else if direction == "R" {
			fmt.Println(fmt.Sprintf("Going right %d", magnitude))
			r.x += magnitude
			if r.x > r.xMax {
				r.xMax = r.x
			}
		} else {
			panic(fmt.Sprintf("invalid direction %s", direction))
		}
		fmt.Println(magnitude)
		fmt.Println(direction)
		fmt.Println(fmt.Sprintf("%d,%d", r.x, r.y))
	}
	fmt.Println(fmt.Sprintf("Ending coordinates: %d,%d  [max x,y = %d,%d] [min x,y = %d,%d]", r.x, r.y, r.xMax, r.yMax, r.xMin, r.yMin))
}
