package main

import (
	"fmt"
	"os"
	"strings"
)

// Vec2 represents a 2D position or direction
type Vec2 struct {
	Row int
	Col int
}

// Add returns a new Vec2 that is the sum of two Vec2s
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		Row: v.Row + other.Row,
		Col: v.Col + other.Col,
	}
}

// Equals checks if two Vec2s are the same
func (v Vec2) Equals(other Vec2) bool {
	return v.Row == other.Row && v.Col == other.Col
}

// Puzzle represents the game state
type Puzzle struct {
	Moves  []Vec2
	Walls  []Vec2
	Boxes  []Vec2
	Robot  Vec2
	Width  int
	Height int
}

// NewPuzzle initializes a new puzzle
func NewPuzzle() *Puzzle {
	return &Puzzle{
		Moves:  make([]Vec2, 0),
		Walls:  make([]Vec2, 0),
		Boxes:  make([]Vec2, 0),
		Robot:  Vec2{},
		Width:  0,
		Height: 0,
	}
}

// ParseMove converts a direction character to a Vec2
func ParseMove(move rune) Vec2 {
	switch move {
	case '<':
		return Vec2{Row: 0, Col: -1}
	case '>':
		return Vec2{Row: 0, Col: 1}
	case '^':
		return Vec2{Row: -1, Col: 0}
	case 'v':
		return Vec2{Row: 1, Col: 0}
	default:
		panic(fmt.Sprintf("Not a valid move: %c", move))
	}
}

// ParseWASD converts a WASD character to a Vec2
func ParseWASD(move rune) Vec2 {
	switch move {
	case 'a':
		return Vec2{Row: 0, Col: -1}
	case 'd':
		return Vec2{Row: 0, Col: 1}
	case 'w':
		return Vec2{Row: -1, Col: 0}
	case 's':
		return Vec2{Row: 1, Col: 0}
	default:
		panic(fmt.Sprintf("Not a valid WASD move: %c", move))
	}
}

// ParseInput reads and parses the puzzle from a string
func ParseInput(input string) *Puzzle {
	puzzle := NewPuzzle()
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Parse the map
	mapEnded := false
	row := 0
	for _, line := range lines {
		if line == "" {
			mapEnded = true
			continue
		}

		if !mapEnded {
			// Parse map
			puzzle.Width = len(line)
			for col, ch := range line {
				pos := Vec2{Row: row, Col: col}
				switch ch {
				case '#':
					puzzle.Walls = append(puzzle.Walls, pos)
				case 'O':
					puzzle.Boxes = append(puzzle.Boxes, pos)
				case '@':
					puzzle.Robot = pos
				}
			}
			row++
		} else {
			// Parse moves
			for _, ch := range line {
				move := ParseMove(ch)
				puzzle.Moves = append(puzzle.Moves, move)
			}
		}
	}
	puzzle.Height = row

	return puzzle
}

// HasWall checks if there's a wall at the given position
func (p *Puzzle) HasWall(pos Vec2) bool {
	for _, wall := range p.Walls {
		if wall.Equals(pos) {
			return true
		}
	}
	return false
}

// HasBoxPart2 checks if there's a box at the given position or
// if it's the right side of a box (since boxes are 2 cells wide in part 2)
func (p *Puzzle) HasBoxPart2(pos Vec2) int {
	// Check if pos is the left side of a box
	for i, box := range p.Boxes {
		if box.Equals(pos) {
			return i
		}
	}

	// Check if pos is the right side of a box
	posLeft := Vec2{Row: pos.Row, Col: pos.Col - 1}
	for i, box := range p.Boxes {
		if box.Equals(posLeft) {
			return i
		}
	}

	return -1
}

// Contains checks if an int slice contains a value
func Contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// PushBoxPart2 tries to push a box in a direction
// Returns whether the push succeeded and which boxes were moved
func (p *Puzzle) PushBoxPart2(boxIndex int, move Vec2) (bool, []int) {
	movedBoxes := []int{}

	box := p.Boxes[boxIndex]
	nextLeft := box.Add(move)
	nextRight := Vec2{Row: nextLeft.Row, Col: nextLeft.Col + 1}

	// Check for walls
	if p.HasWall(nextLeft) || p.HasWall(nextRight) {
		return false, movedBoxes
	}

	// Check for boxes in the way and try to push them too
	boxIndexLeft := p.HasBoxPart2(nextLeft)
	if boxIndexLeft >= 0 && boxIndexLeft != boxIndex {
		ok, pushed := p.PushBoxPart2(boxIndexLeft, move)
		if !ok {
			return false, movedBoxes
		}
		movedBoxes = append(movedBoxes, pushed...)
	}

	boxIndexRight := p.HasBoxPart2(nextRight)
	if boxIndexRight >= 0 && boxIndexRight != boxIndex && boxIndexRight != boxIndexLeft {
		ok, pushed := p.PushBoxPart2(boxIndexRight, move)
		if !ok {
			return false, movedBoxes
		}
		movedBoxes = append(movedBoxes, pushed...)
	}

	// Add this box to the moved list
	movedBoxes = append(movedBoxes, boxIndex)
	return true, movedBoxes
}

// DumpPuzzle2 visualizes the puzzle state
func (p *Puzzle) DumpPuzzle2() {
	for row := 0; row < p.Height; row++ {
		for col := 0; col < p.Width; col++ {
			pos := Vec2{Row: row, Col: col}
			prev := Vec2{Row: row, Col: col - 1}

			if pos.Equals(p.Robot) {
				fmt.Print("@")
			} else if p.HasWall(pos) {
				fmt.Print("#")
			} else if p.HasBoxPart2(pos) >= 0 && (col == 0 || p.HasBoxPart2(prev) < 0) {
				fmt.Print("[")
			} else if col > 0 && p.HasBoxPart2(prev) >= 0 && p.HasBoxPart2(pos) < 0 {
				fmt.Print("]")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Part2 solves the puzzle with wide boxes
func (p *Puzzle) Part2() int {
	// Double the width of the grid
	p.Width *= 2
	p.Robot.Col *= 2

	// Double the width of boxes
	for i := range p.Boxes {
		p.Boxes[i].Col *= 2
	}

	// Double the width of walls
	originalWallsCount := len(p.Walls)
	for i := 0; i < originalWallsCount; i++ {
		wall := p.Walls[i]
		wall.Col *= 2
		p.Walls[i] = wall

		// Add adjacent wall
		clone := Vec2{Row: wall.Row, Col: wall.Col + 1}
		p.Walls = append(p.Walls, clone)
	}

	// Process all moves
	for _, move := range p.Moves {
		next := p.Robot.Add(move)

		// Check for walls
		if p.HasWall(next) {
			continue
		}

		// Check for boxes
		boxIndex := p.HasBoxPart2(next)
		if boxIndex >= 0 {
			ok, boxesToMove := p.PushBoxPart2(boxIndex, move)
			if ok {
				// Move boxes that need to be moved
				processedIndices := make(map[int]bool)
				for _, idx := range boxesToMove {
					if processedIndices[idx] {
						continue
					}
					p.Boxes[idx] = p.Boxes[idx].Add(move)
					processedIndices[idx] = true
				}

				// Move robot
				p.Robot = next
			}
		} else {
			// No box in the way, just move the robot
			p.Robot = next
		}
	}

	// Calculate score
	total := 0
	for _, box := range p.Boxes {
		total += box.Row*100 + box.Col
	}

	return total
}

// Part2Interactive allows playing the puzzle with WASD keys
func (p *Puzzle) Part2Interactive() int {
	// Double the width of the grid
	p.Width *= 2
	p.Robot.Col *= 2

	// Double the width of boxes
	for i := range p.Boxes {
		p.Boxes[i].Col *= 2
	}

	// Double the width of walls
	originalWallsCount := len(p.Walls)
	for i := 0; i < originalWallsCount; i++ {
		wall := p.Walls[i]
		wall.Col *= 2
		p.Walls[i] = wall

		// Add adjacent wall
		clone := Vec2{Row: wall.Row, Col: wall.Col + 1}
		p.Walls = append(p.Walls, clone)
	}

	// Interactive game loop would go here
	// In Go, implementing the raw terminal input would require
	// a terminal library like "github.com/eiannone/keyboard"
	// For simplicity, I'll omit the actual interactive implementation

	// Calculate score
	total := 0
	for _, box := range p.Boxes {
		total += box.Row*100 + box.Col
	}

	return total
}

func main() {
	// Read input from stdin

	data, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		os.Exit(1)
	}

	input := string(data)
	// Part 1 would be here
	// ...

	// Part 2
	puzzle := ParseInput(input)
	result2 := puzzle.Part2()
	fmt.Printf("Part2: %d\n", result2)
}
