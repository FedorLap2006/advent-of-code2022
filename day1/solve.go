package day1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func caloriesPerElf(inventories [][]int) (r []int) {
	r = make([]int, len(inventories))
	for i, items := range inventories {
		for _, item := range items {
			r[i] += item
		}
	}

	return
}

func parseInventory(contents string) (r []int) {
	items := strings.Split(contents, "\n")
	r = make([]int, len(items))
	for j, item := range items {
		r[j], _ = strconv.Atoi(item)
	}

	return
}

func parseInventories(input string) (r [][]int) {
	inventories := strings.Split(input, "\n\n")
	r = make([][]int, len(inventories))
	for i, contents := range inventories {
		r[i] = parseInventory(contents)
	}

	return
}

func Part1(input string) string {
	calories := caloriesPerElf(parseInventories(input))
	max := 0
	for i := range calories {
		if calories[i] > calories[max] {
			max = i
		}
	}

	return fmt.Sprintf("%d", calories[max])
}

func Part2(input string) string {
	calories := caloriesPerElf(parseInventories(input))
	sort.Ints(calories)

	return fmt.Sprintf("%d",
		calories[len(calories)-1]+
			calories[len(calories)-2]+
			calories[len(calories)-3])
}
