package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertStringsToInt(input []string) []int {
	var numbers []int
	for _, s := range input {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, i)
	}
	return numbers
}
func readinputfile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	return rows
}

func calcdiff(a int, b int) int {
	dif := a - b
	if dif < 0 {
		dif = dif * -1
	}
	return dif
}

func asc(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i-1] > report[i] {
			return false
		}
		dif := calcdiff(report[i], report[i-1])
		if dif < 1 || dif > 3 {
			return false
		}
	}
	return true
}

func des(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i] > report[i-1] {
			return false
		}

		dif := calcdiff(report[i], report[i-1])
		if dif < 1 || dif > 3 {
			return false
		}
	}
	return true
}

func isSave(report []int) bool {
	if len(report) < 2 {
		panic("report is to short < 2")
	}
	if report[0] < report[1] {
		return asc(report)
	}
	return des(report)
}

func main() {
	rows := readinputfile("./test.txt")
	var numbers [][]int
	for _, line := range rows {
		parts := strings.Split(line, " ")
		numbers = append(numbers, convertStringsToInt(parts))
	}
	if len(numbers) == 0 {
		panic("there has it what")
	}
	save_counter := 0
	for _, number := range numbers {
		if isSave(number) {
			save_counter++
		}
		fmt.Println(number, " ", isSave(number))
	}

	fmt.Println(save_counter)
}
