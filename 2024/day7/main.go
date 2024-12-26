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
	value     int
	operation string
}

func main() {
	part1()
}

func walk(equation_elements []int, curr int, operation string, result int, path []Equation) bool {

	//pre
	// En los valores revisados aun no existe un valor
	if len(path) == 0 {
		path = append(path, Equation{value: curr, operation: operation})
	}

	// LLegamos al final y las equaciones no dan el resultado
	if len(equation_elements) == 0 {
		local_counter := 0
		// Validamos si las operaciones dan el resultado
		for _, v := range path {
			if v.operation == "+" {
				local_counter += v.value
			}

			if v.operation == "*" {
				if local_counter == 0 {
					local_counter = 1
				}
				local_counter = local_counter * v.value
			}
		}

		if local_counter == result {
			return true
		}

		return false
	}

	//recurse

	//post
	return false
}

func part1() {

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	total_sum := 0
	test_value := ""
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, ": ")

		test_value = elements[0]
		equation_numbers := strings.Split(elements[1], " ")
		// head_node := BinaryNode{}
		var equation_numbers_integers []int
		for i, v := range equation_numbers {
			value, _ := strconv.Atoi(v)
			// next_value, _ := strconv.Atoi(equation_numbers[i+1])
			equation_numbers_integers = append(equation_numbers_integers, value)

		}
	}

	println(total_sum)
}

func part2() {
}
