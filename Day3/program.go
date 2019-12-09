package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	w1 := NewWirePath(Point32{x: 0, y: 0}, "R8,U5,L5,D3")
	w2 := NewWirePath(Point32{x: 0, y: 0}, "U7,R6,D4,L4")

	fmt.Println(fmt.Sprintf("Wire 1 traverses %d lengths", len(w1.path)-1))
	fmt.Println(fmt.Sprintf("Wire 2 traverses %d lengths", len(w2.path)-1))

	//w2 := NewWirePath(Point32{x: 0, y: 0}, "")
	//w1.FollowPath(firstWirePlot)
	//w2.FollowPath(secondWirePlot)
	grid, _ := DrawPaths(w1, w2)

	for _, line := range grid {
		fmt.Printf("%s\n", string(line))
	}
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

// DrawPaths returns a 2d array of of runes containing the wires paths and also an array of Point32 of overlaps
func DrawPaths(wirePath1 WirePath, wirePath2 WirePath) ([][]rune, []Point32) {
	minPoint := Point32{x: min32(wirePath1.min.x, wirePath2.min.x) - 1, y: min32(wirePath1.min.y, wirePath2.min.y) - 1}
	maxPoint := Point32{x: max32(wirePath1.max.x, wirePath2.max.x) + 1, y: max32(wirePath1.max.y, wirePath2.max.y) + 1}

	fmt.Println(fmt.Sprintf("grid should go from : %d,%d to %d,%d", minPoint.x, minPoint.y, maxPoint.x, maxPoint.y))

	grid := [][]rune{}
	for i := minPoint.y; i <= maxPoint.y; i++ {
		grid = append(grid, []rune(strings.Repeat(".", int(maxPoint.x-minPoint.x))))
	}

	for i := 0; i <= len(wirePath1.path)-1; i++ {
		if wirePath1.path[i].x == wirePath1.path[i+1].x {
			// vertical
			// x := (wirePath1.path[i].x + 1) * -1
			// minY := min32(wirePath1.path[i].y, wirePath2.path[i+1].y) // + minPoint.y
			// maxY := max32(wirePath1.path[i].y, wirePath2.path[i+1].y) // + minPoint.y
			// for y := minY; y < maxY+1; y++ {
			// 	grid[y][x] = 'x'
			// }
		} else {
			// horizontal
		}
	}

	grid[maxPoint.y][minPoint.x*-1] = 'o'

	return grid, nil
}

func min32(a, b int32) int32 {
	if a <= b {
		return a
	}
	return b
}

func max32(a, b int32) int32 {
	if a >= b {
		return a
	}
	return b
}

// NewWirePath returns a new WirePath following the provided path, which
// should be a comma-delimited string of directions starting from the origin.
func NewWirePath(origin Point32, path string) WirePath {
	p := WirePath{start: origin, end: origin, max: origin, min: origin, path: []Point32{origin}}
	p.FollowPath(strings.Split(path, ","))
	return p
}

// FollowPath accepts a path represented as strings and follows it from current x,y
func (r *WirePath) FollowPath(path []string) {

	//fmt.Println(fmt.Sprintf("Starting coordinates: %d,%d", r.end.x, r.end.y))
	for _, element := range path {
		direction := element[:1]
		m, _ := strconv.ParseInt(element[1:len(element)], 10, 32)
		magnitude := int32(m)

		if direction == "U" {
			//fmt.Println(fmt.Sprintf("Going up %d", magnitude))
			r.end.y += magnitude
			if r.end.y > r.max.y {
				r.max.y = r.end.y
			}
		} else if direction == "D" {
			//fmt.Println(fmt.Sprintf("Going down %d", magnitude))
			r.end.y -= magnitude
			if r.end.y < r.min.y {
				r.min.y = r.end.y
			}
		} else if direction == "L" {
			//fmt.Println(fmt.Sprintf("Going left %d", magnitude))
			r.end.x -= magnitude
			if r.end.x < r.min.x {
				r.min.x = r.end.x
			}
		} else if direction == "R" {
			//fmt.Println(fmt.Sprintf("Going right %d", magnitude))
			r.end.x += magnitude
			if r.end.x > r.max.x {
				r.max.x = r.end.x
			}
		} else {
			panic(fmt.Sprintf("invalid direction %s", direction))
		}

		//fmt.Println(fmt.Sprintf("%d,%d", r.end.x, r.end.y))
		r.path = append(r.path, r.end)
	}
	//fmt.Println(fmt.Sprintf("Ending coordinates: %d,%d  [max x,y = %d,%d] [min x,y = %d,%d]", r.end.x, r.end.y, r.max.x, r.max.y, r.min.x, r.min.y))
}
