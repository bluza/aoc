package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(path string) [][]byte {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err, "could not find file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var field [][]byte
	for scanner.Scan() {
		field = append(field, scanner.Bytes())
	}
	return field
}

func outofbounds(row int, col int, side_length int) bool {
	if col < 0 || col > side_length-1 {
		return true
	}
	if row < 0 || row > side_length-1 {
		return true
	}
	return false
}

func FindString(row int, col int, dir Direction, field *[][]byte) bool {

	tofind := "XMAS"
	tfield := *field
	if tfield[row][col] != byte(tofind[0]) {
		return false
	}
	current_row := row
	current_col := col
	for _, c := range tofind[1:] {
		current_row += dir.x
		current_col += dir.y
		if outofbounds(current_row, current_col, len(tfield)) {
			return false
		}
		if tfield[current_row][current_col] != byte(c) {
			return false
		}
	}
	fmt.Println("found xmas: ", row, col, dir)
	return true
}

func main() {
	field := readFile("input.txt")
	counter := 0
	for row := range field {
		for col := range field[0] {
			for _, d := range Directions {
				if FindString(row, col, d, &field) {
					counter++
				}
			}
		}
	}
	fmt.Println("result: ", counter)
}
