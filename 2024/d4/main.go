package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFileNOTWORKING(field *[][]byte, path string) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err, "could not find file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		*field = append(*field, line)
	}
}

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err, "could not find file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func outOfBounds(row int, col int, side_length int) bool {
	if col < 0 || col > side_length-1 {
		return true
	}
	if row < 0 || row > side_length-1 {
		return true
	}
	return false
}

func findString(row int, col int, dir Direction, field *[]string) bool {
	tofind := "XMAS"
	tfield := *field

	current_row := row
	current_col := col
	for _, c := range tofind {
		if outOfBounds(current_row, current_col, len(tfield)) {
			return false
		}
		if tfield[current_row][current_col] != byte(c) {
			return false
		}
		current_row += dir.y
		current_col += dir.x
	}
	return true
}

func main() {

	var fieldNOTWORKING [][]byte
	readFileNOTWORKING(&fieldNOTWORKING, "input.txt")

	field := readFile("input.txt")
	counter := 0
	for row := range field {
		for col := range field[0] {
			for _, dir := range Directions {
				if findString(row, col, dir, &field) {
					counter++
				}
			}
		}
	}
	fmt.Println("result: ", counter)
}
