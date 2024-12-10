package main

type Direction struct {
	x int
	y int
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
