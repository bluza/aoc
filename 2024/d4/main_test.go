package main

import "testing"

var field = readFile("./test.txt")

func TestRowInc(t *testing.T) {
	if FindString(0, 5, Direction{0, 1}, &field) == false {
		t.Fatalf("did not find")
	}
}
func TestRowreverse(t *testing.T) {
	if FindString(1, 4, Direction{0, -1}, &field) == false {
		t.Fatalf("did not find")
	}
}
func TestCol(t *testing.T) {
	if FindString(3, 9, Direction{1, 0}, &field) == false {
		t.Fatalf("did not find")
	}
}
func TestColrev(t *testing.T) {
	if FindString(9, 9, Direction{-1, 0}, &field) == false {
		t.Fatalf("did not find")
	}
}
func TestDiag(t *testing.T) {
	if FindString(0, 4, Direction{1, 1}, &field) == false {
		t.Fatalf("did not find")
	}
}
func TestDiagrev(t *testing.T) {
	if FindString(9, 9, Direction{-1, -1}, &field) == false {
		t.Fatalf("did not find")
	}
}
