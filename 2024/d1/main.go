package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func convertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func countIntInSlice(s []int, lookup int) int {
	counter := 0
	for _, val := range s {
		if val == lookup {
			counter++
		}
	}
	return counter
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var firstval []int
	var secVal []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := scanner.Text()
		noSpaces := strings.Join(strings.Fields(raw), " ")
		splitted := strings.Split(noSpaces, " ")
		firstval = append(firstval, convertToInt(splitted[0]))
		secVal = append(secVal, convertToInt(splitted[1]))

	}
	sort.Ints(firstval)
	sort.Ints(secVal)

	if len(firstval) != len(secVal) {
		panic("the two lists does not have the same length")
	}
	// solution first star
	sum := 0
	for i := 0; i < len(firstval); i++ {
		distance := secVal[i] - firstval[i]
		if distance < 0 {
			distance = distance * -1
		}
		sum += distance
		// fmt.Printf("%d %d distance: %d \n", firstval[i], secVal[i], distance)
	}
	fmt.Printf("sum: %d \n", sum)

	sum = 0
	for _, value := range firstval {
		sum += value * countIntInSlice(secVal, value)
	}
	fmt.Printf("sum part2 %d", sum)

}
