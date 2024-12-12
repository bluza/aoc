package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	path := "test.txt"
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	var temp_f [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		t := strings.Split(line, "")
		temp_f = append(temp_f, t)
		if guard_idx := slices.Index(t, "^"); guard_idx != -1 {
			fmt.Println("found the guard at", guard_idx, ":", len(temp_f)-1)
		}
	}
	field := Field{temp_f}

	fmt.Println(field)
}
