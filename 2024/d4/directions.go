package main

type Direction struct {
	x int
	y int
}

func (d Direction) Next(increment int) (int, int) {
	var x, y int
	if d.x == 0 {
		x = 0
	} else if d.x == 1 {
		x = increment
	} else {
		x = -increment
	}
	if d.y == 0 {
		y = 0
	} else if d.y == 1 {
		y = increment
	} else {
		y = -increment
	}
	return x, y
}

var Directions = []Direction{
	{0, 1},
	{0, -1},
	{1, 0},
	{1, 1},
	{1, -1},
	{-1, 0},
	{-1, 1},
	{-1, -1},
}
