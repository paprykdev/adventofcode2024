package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func string_to_integers(lines []string) []int {
	var integers []int
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		check(err)
		integers = append(integers, n)
	}
	return integers
}

func is_safe(nums []int) bool {
	increasing := nums[1] > nums[0]
	if increasing {
		for i := 1; i < len(nums); i++ {
			diff := nums[i] - nums[i-1]
			if !(1 <= diff && diff <= 3) {
				return false
			}
		}
		return true
	} else {
		for i := 1; i < len(nums); i++ {
			diff := nums[i] - nums[i-1]
			if !(-3 <= diff && diff <= -1) {
				return false
			}
		}
		return true
	}
}

func is_really_safe(nums []int) bool {
	if is_safe(nums) {
		return true
	}
	for i := range len(nums) {
		slice := make([]int, len(nums))
		copy(slice, nums)
		slice = append(slice[:i], slice[i+1:]...)
		if is_safe(slice) {
			return true
		}
	}
	return false
}

func main() {
	data, err := os.ReadFile("input/2024_2.txt")
	check(err)
	lines := strings.Split(string(data), "\n")
	var safe_reports int = 0
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}

		line := strings.Split(line, " ")
		nums := string_to_integers(line)

		if is_really_safe(nums) {
			safe_reports++
		}
	}
	fmt.Println(safe_reports)
}
