package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(leftArray, rightArray []int) int {
	total := 0

	sort.Slice(leftArray, func(i, j int) bool {
		return leftArray[i] < leftArray[j]
	})
	sort.Slice(rightArray, func(i, j int) bool {
		return rightArray[i] < rightArray[j]
	})

	for i := 0; i < len(leftArray); i++ {
		distance := math.Abs(float64(leftArray[i]) - float64(rightArray[i]))
		total += int(distance)
	}

	return total
}

func part2(leftArray, rightArray []int) int {
	total := 0

	for i := range len(leftArray) {
		count := 0
		for j := range len(rightArray) {
			if rightArray[j] == leftArray[i] {
				count++
			}
		}
		total += leftArray[i] * count
	}

	return total
}

func main() {
	data, err := os.ReadFile("input/2024_1.txt")
	check(err)
	lines := strings.Split(string(data), "\n")
	leftArray, rightArray := make([]int, 0), make([]int, 0)
	for _, line := range lines {
		line := strings.Split(string(line), "   ")

		if len(line) == 2 {
			n1, err := strconv.Atoi(line[0])
			check(err)

			n2, err := strconv.Atoi(line[1])
			check(err)

			leftArray = append(leftArray, n1)
			rightArray = append(rightArray, n2)
		}
	}

	total1 := part1(leftArray, rightArray)
	total2 := part2(leftArray, rightArray)
	fmt.Printf("Part 1: %d\n", total1)
	fmt.Printf("Part 2: %d\n", total2)
}
