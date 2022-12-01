package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"aoc/day1"
)

var days = [][2]func(s string) string{
	{day1.Part1, day1.Part2},
}

func main() {
	if len(os.Args) < 4 {
		log.Fatal("expected day, part and input file")
	}

	day, _ := strconv.Atoi(os.Args[1])
	if day > 25 || day < 1 {
		log.Fatal("invalid day")
	}
	day--
	if day > len(days) {
		log.Fatal("patience, padawan")
	}

	part, _ := strconv.Atoi(os.Args[2])

	if part < 1 || part > 2 {
		log.Fatal("invalid part")
	}
	part--
	if days[day][part] == nil {
		log.Fatal("not implemented")
	}

	input, err := ioutil.ReadFile(os.Args[3])
	if err != nil {
		log.Fatalf("failed to read input: %s", err)
	}

	fmt.Println(days[day][part](string(input)))
}
