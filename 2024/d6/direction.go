package main

type Orientation struct {
	x int
	y int
}

const (
	North int = iota
	West
	South
	East
)

var orientations = []Orientation{
	{-1, 0},
	{0, 1},
	{-1, 0},
	{1, 0},
}

func (d *Orientation) Turn() {

}
