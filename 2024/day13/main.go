package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	part1()
	// part2()
}

type Button struct {
	row int
	col int
}

func part1() {

	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	temp := strings.Split(file, "\n\n")

	for _, crawl_machine := range temp {
		// fmt.Println(crawl_machine)

		individual_lines := strings.Split(crawl_machine, "\n")
		// fmt.Println(individual_lines)

		fmt.Println(individual_lines[0], individual_lines[1], individual_lines[2])
	}
	// fmt.Println(temp[0])
}

func part2() {
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	file := string(data)

	temp := strings.Split(file, "\n")
	fmt.Println(temp)
}
