package main

import "testing"

var rules, updates = parseFile("test.txt")

func TestFirstSampeUpdate(t *testing.T) {
	if updateIsValid(updates[0], rules) == false {
		t.Fatalf("should be a valid update")
	}
}

func TestSecondSampeUpdate(t *testing.T) {
	if updateIsValid(updates[1], rules) == false {
		t.Fatalf("should be a valid update")
	}
}

func TestThirdSampeUpdate(t *testing.T) {
	if updateIsValid(updates[2], rules) == false {
		t.Fatalf("should be a valid update")
	}
}
func TestFourthSampeUpdate(t *testing.T) {
	if updateIsValid(updates[3], rules) != false {
		t.Fatalf("not a valid update")
	}
}
func TestFifthSampeUpdate(t *testing.T) {
	if updateIsValid(updates[4], rules) != false {
		t.Fatalf("not a valid update")
	}
}
func TestSixedSampeUpdate(t *testing.T) {
	if updateIsValid(updates[5], rules) != false {
		t.Fatalf("not a valid update")
	}
}
