package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// TODO:
// 1. find 1 - n digits followed by a ","
// 2. find 1 -n digits followed by a ")"
// every other char or char combination unvalidates the instruction

func findValidInstructions(reader *bufio.Reader) [2]int {
	var digits [2]int

	var delimiter = []byte{',', ')'}
	for i, d := range delimiter {
		// fmt.Println(string(d))
		t, err := reader.ReadString(d)
		if err != nil {
			log.Fatal(err)
			// do hom ma nix gfunden un miasn returnen
		}
		t_digit := t[:len(t)-1]

		if vl, err := strconv.Atoi(t_digit); err == nil {
			// fmt.Println(t)
			digits[i] = vl
		} else {
			fmt.Println("do hods wos")
			digits[0] = -1
			digits[1] = -1
			return digits
		}
	}
	return digits
}

func main() {

	file, err := os.Open("./test.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	const Instruction = "mul("
	var to_calc [][2]int
	for {
		peak_test, _ := reader.Peek(len(Instruction))
		// fmt.Println(string(peak_test))

		if string(peak_test) != Instruction {
			reader.Discard(1)
		} else {
			// fmt.Println("match")
			reader.Discard(len(Instruction))
			digits := findValidInstructions(reader)
			if digits[0] != -1 {
				to_calc = append(to_calc, digits)
			}
			fmt.Println(digits)
		}

	}

	fmt.Println(to_calc)

	// peak_size := make([]byte, 3)
	// _ = peak_size
	// size, _ := reader.Read(peak_size)

	// fmt.Println(string(peak_size))
	// fmt.Println(size)

	// peak_test, _ := reader.Peek(len(Instruction))
	// if string(peak_test) != Instruction {
	// 	reader.Discard(1)
	// } else {
	// 	fmt.Println("match")
	// }
	// fmt.Println("second peak")
	// peak_test, _ = reader.Peek(len(Instruction))
	// fmt.Println(string(peak_test))
}
