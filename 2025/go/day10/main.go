package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	lightDiagram        []bool
	buttonWiring        [][]int
	joltageRequirements []int
}

func main() {
	data, err := os.Open("./part1.txt")

	if err != nil {
		panic(err)
	}

	defer data.Close()
	scanner := bufio.NewScanner(data)

	var lines []Line
	for scanner.Scan() {

		line := scanner.Text()
		sections := strings.Split(line, " ")
		// fmt.Println(sections[len(sections)-1])

		var currentLineDiagram []bool
		for _, value := range sections[0] {
			if string(value) == "." {
				currentLineDiagram = append(currentLineDiagram, false)
			}
			if string(value) == "#" {
				currentLineDiagram = append(currentLineDiagram, true)
			}
		}

		var currentButtonWiring [][]int
		for _, bW := range sections[1 : len(sections)-1] {
			var tmp []int
			clean_button := strings.Replace(strings.Replace(bW, "(", "", 1), ")", "", 1)
			// fmt.Println(strings.Split(clean_button, ","))
			currStringButton := strings.SplitSeq(clean_button, ",")

			for i := range currStringButton {
				val, err := strconv.Atoi(i)
				if err != nil {
					panic(err)
				}

				tmp = append(tmp, val)
			}

			currentButtonWiring = append(currentButtonWiring, tmp)
		}

		var currentJoltage []int
		for _, jolt := range sections[len(sections)-1] {

			clean_jolt := strings.Replace(strings.Replace(string(jolt), "{", "", 1), "}", "", 1)

			// fmt.Println(string(jolt),clean_jolt)
			curr := strings.SplitSeq(clean_jolt, ",")

			for i := range curr {
				val, err := strconv.Atoi(i)
				if err != nil {
					// panic(err)
					continue
				}

				currentJoltage = append(currentJoltage, val)
			}
		}

		currentLine := Line{
			lightDiagram:        currentLineDiagram,
			buttonWiring:        currentButtonWiring,
			joltageRequirements: currentJoltage,
		}

		lines = append(lines, currentLine)
	}

	// part1(lines)
	part2(lines)

}

func part1(lines []Line) {
	counter := 0
	for _, line := range lines {
		currSolution := findMinimumCombination(line.lightDiagram, line.buttonWiring)

		counter = counter + len(currSolution)
	}

	fmt.Println(counter)
}

func findMinimumCombination(desired []bool, buttonWiring [][]int) [][]int {

	n := len(buttonWiring)

	minSize := n + 1

	var bestCombination [][]int

	for mask := 1; mask < (1 << n); mask++ {
		var combination [][]int
		for i := range n {
			if mask&(1<<i) != 0 {
				combination = append(combination, buttonWiring[i])
			}
		}

		if matchesDesired(desired, combination) {
			if len(combination) < minSize {
				minSize = len(combination)
				bestCombination = combination
			}
		}
	}
	return bestCombination
}

func matchesDesired(desired []bool, combination [][]int) bool {
	current := make([]bool, len(desired))
	for _, switches := range combination {
		for _, pos := range switches {
			current[pos] = !current[pos]
		}
	}

	for i := range desired {
		if current[i] != desired[i] {
			return false
		}
	}

	return true
}

func part2(lines []Line) {
	counter := 0
	for _, line := range lines {
		currSolution := findMinimumCombinationInt(line.joltageRequirements, line.buttonWiring)

		counter = counter + len(currSolution)
	}

	fmt.Println(counter)
}

func findMinimumCombinationInt(desired []int, buttonWiring [][]int) [][]int {

	n := len(buttonWiring)

	// minSize := n + 1

	var bestCombination [][]int

	for mask := 1; mask < (1 << n); mask++ {
		var combination [][]int
		for i := range n {
			if mask&(1<<i) != 0 {
				combination = append(combination, buttonWiring[i])
			}
		}

		// if matchesDesired(desired, combination) {
		// 	if len(combination) < minSize {
		// 		minSize = len(combination)
		// 		bestCombination = combination
		// 	}
		// }
	}
	return bestCombination
}

func matchesDesiredInt(desired []int, combination [][]int) bool {
	current := make([]int, len(desired))
	for _, switches := range combination {
		for _, pos := range switches {
			current[pos]++
		}
	}

	for i := range desired {
		if current[i] != desired[i] {
			return false
		}
	}

	return true
}
