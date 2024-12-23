package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
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
		var equation_numbers_integers []int
		for _, v := range equation_numbers {
			value, _ := strconv.Atoi(v)
			equation_numbers_integers = append(equation_numbers_integers, value)
			println(value)
		}
		sum := 0
		// mul := 1
		mixed := 1
		// for _, v := range equation_numbers_integers {
		// 	sum += v
		// 	mul = mul * v
		// }
		//
		int_test_value, _ := strconv.Atoi(test_value)
		// if sum == int_test_value {
		// 	total_sum += int_test_value
		// }
		//
		// if mul == int_test_value {
		// 	total_sum += int_test_value
		// }

		var can_be_mul []int
		var can_be_sum []int
		for i, _ := range equation_numbers_integers {

			if i == len(equation_numbers_integers) {
				continue
			}

			if int_test_value%equation_numbers_integers[i] == 0 {
				can_be_mul = append(can_be_mul, equation_numbers_integers[i])
				println("Can be mul", equation_numbers_integers[i], sum, mixed)
			} else {
				can_be_sum = append(can_be_sum, equation_numbers_integers[i])
			}
		}

		for _, v := range can_be_mul {
			// println(v)
			mixed = mixed * v
		}

		for _, v := range can_be_sum {
			// println(v)
			mixed += mixed + v
		}

		if int_test_value == mixed {
			total_sum += int_test_value
		}
	}

	println(total_sum)
}

func part2() {
}
