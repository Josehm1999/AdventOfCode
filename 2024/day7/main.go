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

func repeatSlices(slice [][]string, repeate int) [][]string {

	pools := make([][]string, len(slice)*repeate)

	for i := range pools {
		copy(pools[i*len(slice):], slice)
	}

	result := [][]string{}

	for _, pool := range pools {
		newResult := [][]string{}

		if len(result) == 0 {
			for _, y := range pool {
				newResult = append(newResult, []string{y})
				result = newResult
			}
		} else {
			for _, x := range result {
				for _, y := range pool {
					newList := make([]string, len(x)+1)
					copy(newList, x)
					newList[len(x)] = y
					newResult = append(newResult, newList)
				}
			}
		}
		result = newResult
	}

	for i, v := range result {

		for _, value := range v {
			println(i, len(v), value)
		}
	}

	return result
}

func test(counter int, operations []string, equation_numbers []string) int {

	// for i, v := range equation_numbers {
	// 	println(i, v)
	//
	// 	value, _ := strconv.Atoi(equation_numbers[i])
	// 	if operations[i] == "+" {
	// 		counter += value
	// 	} else {
	// 		counter *= value
	// 	}
	// }
	for i := 1; i < len(equation_numbers)+1; i++ {
		// value, _ := strconv.Atoi(equation_numbers[i-1])
		// println(value, len(operations), operations[0])
		// if operations[i-1] == "+" {
		// 	counter += value
		// } else {
		// 	counter *= value
		// }
	}
	return counter
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
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Split(line, ": ")

		test_value, _ = strconv.Atoi(elements[0])
		equation_numbers := strings.Split(elements[1], " ")
		// equations := [][]Equation{}
		operations := [][]string{{"+", "*"}}

		repeated := repeatSlices(operations, len(equation_numbers)-1)
		// ans := 0
		println("New line", repeated, test_value)
		// for _, v := range repeated {
		// 	if test(ans, v, equation_numbers) == test_value {
		// 		println("Works")
		// 		total_sum += test_value
		// 		break
		// 	}
		// }
		// for _, equation := range equation_numbers {
		// 	for _, oper := range operations {
		//
		//
		// 	}
		// }

		// for i, v := range equation_numbers {
		// 	value, _ := strconv.Atoi(v)
		// 	// next_value, _ := strconv.Atoi(equation_numbers[i+1])
		// 	equation_numbers_integers = append(equation_numbers_integers, value)
		//
		// }
	}

	println(total_sum)
}

func part2() {
}
