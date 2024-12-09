package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func FindString(row int, col int, dir Direction, field *[][]byte) bool {

	tofind := "XMAS"
	tfield := *field
	if tfield[row][col] != byte(tofind[0]) {
		return false
	}
	for i, c := range tofind[1:] {
		x, y := dir.Next(i + 1)
		//fmt.Println("move to: ", x, y)
		nxt_row := row + x
		nxt_col := col + y
		if nxt_col < 0 || nxt_col > len(tfield[0])-1 {
			return false
		}
		if nxt_row < 0 || nxt_row > len(tfield)-1 {
			return false
		}
		//fmt.Println(string(tfield[next_row][new_col]))
		if tfield[nxt_row][nxt_col] != byte(c) {
			return false
		}
	}
	fmt.Println("found xmas: ", row, col, dir)
	return true

}

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

func main() {

	field := readFile("./test.txt")
	counter := 0
	for row := 0; row < len(field); row++ {
		for col := 0; col < len(field[0]); col++ {
			for _, d := range Directions {
				if FindString(row, col, d, &field) {
					counter++
				}
			}
		}
	}
	fmt.Println("result: ", counter)
}
