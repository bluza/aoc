package main

type Position struct {
	row int
	col int
}

type Guard struct {
	location    Position
	orientation Orientation
}

func (g *Guard) Move() {
	g.location.row += g.orientation.x
	g.location.col += g.orientation.y
}
