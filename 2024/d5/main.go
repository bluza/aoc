package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rules map[int][]int
type Update []int

func convertToInt(raw []string) []int {
	var numbers []int
	for _, s := range raw {
		i, _ := strconv.Atoi(s)
		numbers = append(numbers, i)
	}
	return numbers
}

func updateIsValid(update Update, rules Rules) bool {
	for curr_pos, curr_number := range update {
		for idx := range curr_pos + 1 {
			if slices.Contains(rules[curr_number], update[idx]) {
				return false
			}
		}
	}
	return true
}

func parseFile(path string) (Rules, []Update) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	before := make(map[int][]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := convertToInt(strings.Split(line, "|"))
		_ = parts
		_ = before
		before[parts[0]] = append(before[parts[0]], parts[1])
	}
	// parse updates
	var updates []Update
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		updates = append(updates, convertToInt(parts))
	}
	return before, updates
}

func main() {
	path := "input.txt"
	rules, updates := parseFile(path)
	sum := 0
	for _, update := range updates {
		if updateIsValid(update, rules) {
			sum += update[(len(update)-1)/2]
		}
	}

	fmt.Println(sum)

}
