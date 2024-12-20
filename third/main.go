package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.ReadFile("input/2024_3.txt")

	fileString := string(file)
	check(err)

	r, err := regexp.Compile("(do\\(\\))|(don't\\(\\))|(mul\\(\\d+,\\d+\\))")
	check(err)

	tokens := r.FindAllString(fileString, -1)
	var all_valid_strings []string
	enabled := true

	for _, token := range tokens {
		switch token {
			case "don't()": {
				enabled = false
			}
			case "do()": {
				enabled = true
			}
			default: {
				if enabled {
					all_valid_strings = append(all_valid_strings, token)
				}
			}
		}
	}


	result := 0

	for _, string := range all_valid_strings {
		r, err := regexp.Compile("[0-9]+")
		check(err)
		numbers := r.FindAllString(string, -1)
		var integers []int
		for _, number := range numbers {
			num, err := strconv.Atoi(number)
			check(err)
			integers = append(integers, num)
		}
		multiplication := 1
		for _, number := range integers {
			multiplication *= number
		}
		result += multiplication
	}
	fmt.Println(result)
}
