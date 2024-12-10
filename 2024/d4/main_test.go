package main

import "testing"

var field = readFileNOTWORKING("./test.txt")

func TestRowInc(t *testing.T) {
	if findString(0, 5, Direction{0, 1}, &field) == false {
		t.Fatalf("did not find")
	}
}
func TestRowreverse(t *testing.T) {
	if findString(1, 4, Direction{0, -1}, &field) == false {
		t.Fatalf("did not find")
	}
}
func TestCol(t *testing.T) {
	if findString(3, 9, Direction{1, 0}, &field) == false {
		t.Fatalf("did not find")
	}
}
func TestColrev(t *testing.T) {
	if findString(9, 9, Direction{-1, 0}, &field) == false {
		t.Fatalf("did not find")
	}
}
func TestDiag(t *testing.T) {
	if findString(0, 4, Direction{1, 1}, &field) == false {
		t.Fatalf("did not find")
	}
}
func TestDiagrev(t *testing.T) {
	if findString(9, 9, Direction{-1, -1}, &field) == false {
		t.Fatalf("did not find")
	}
}
