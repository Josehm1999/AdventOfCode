package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	value    int
	operands []int
}

func main() {
	part1()
}

func walk(equation_elements []Equation) int {
	counter := 0
	for i := 0; i < len(equation_elements); i++ {
		is_valid := walk_recursive(equation_elements[i].operands, 0, equation_elements[i].operands[0], equation_elements[i].value)

		if is_valid {
			counter += equation_elements[i].value
			println(counter)
		}

	}
	return counter
}

func walk_recursive(equation_elements []int, curr int, current_result int, result int) bool {

	// println(len(equation_elements), curr, current_result, result, current_result == result)
	if current_result == result {
		return true
	}

	if len(equation_elements)-1 <= curr {
		return false
	}

	// Concatenation is part 2
	concat_part1 := strconv.Itoa(current_result)
	concat_part2 := strconv.Itoa(equation_elements[curr+1])
	concatenation_val, _ := strconv.Atoi(concat_part1 + concat_part2)

	sum := walk_recursive(equation_elements, curr+1, current_result+equation_elements[curr+1], result)
	mul := walk_recursive(equation_elements, curr+1, current_result*equation_elements[curr+1], result)
	concatenation := walk_recursive(equation_elements, curr+1, concatenation_val, result)

	return sum || mul || concatenation
}

type Seen struct {
	x int
	y int
}

func part1() {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	// line := 0

	temp := strings.Split(file, "\n")

	sum := 0
	var seen []Seen

	for i := 0; i < len(temp)-1; i++ {
		fmt.Println(temp[i])
		isPossibleUp := false
		isPossibleDown := false

		if (len(temp)-2)-i > 2 {
			isPossibleDown = true
		}
		if i > 2 {
			isPossibleUp = true
		}

		letters := strings.Split(temp[i], "")
		println(isPossibleUp, isPossibleDown)
		for j, v := range letters {
			println(j, v)
			seen := append(seen, Seen{x: j, y: i})
		}

	}
	println(sum, temp)
}
