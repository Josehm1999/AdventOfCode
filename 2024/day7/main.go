package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type BinaryNode struct {
	value     int
	operation string
	leftNode  *BinaryNode
	rightNode *BinaryNode
}

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

	// println(current_result, result, current_result == result)
	concat_part1 := strconv.Itoa(current_result)
	concat_part2 := strconv.Itoa(equation_elements[curr+1])
	concatenation_val, _ := strconv.Atoi(concat_part1 + concat_part2)

	sum := walk_recursive(equation_elements, curr+1, current_result+equation_elements[curr+1], result)
	mul := walk_recursive(equation_elements, curr+1, current_result*equation_elements[curr+1], result)
	concatenation := walk_recursive(equation_elements, curr+1, concatenation_val, result)

	return sum || mul || concatenation
}

func part1() {

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	total_sum := 0
	test_value := 0
	var equations []Equation

	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, ": ")

		test_value, _ = strconv.Atoi(elements[0])
		equation_numbers := strings.Split(elements[1], " ")

		var equation_numbers_integers []int
		for _, v := range equation_numbers {
			value, _ := strconv.Atoi(v)
			equation_numbers_integers = append(equation_numbers_integers, value)

		}
		curr_equation := Equation{value: test_value, operands: equation_numbers_integers}
		equations = append(equations, curr_equation)
	}

	total_sum = walk(equations)
	println(equations[0].operands[0], total_sum)
}

func part2() {
}
