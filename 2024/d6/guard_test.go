package main

import "testing"

func TestGuardMoveUp(t *testing.T) {
	start_loc := Position{1, 1}
	orientation := Orientation{-1, 0}
	guard := Guard{start_loc, orientation}

	guard.Move()

	if guard.location.row != 0 && guard.location.col != 1 {
		t.Fatalf("wrong location")
	}
}
