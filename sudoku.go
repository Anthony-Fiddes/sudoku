package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Puzzle [][]Tile

func NewPuzzle(length, width int) Puzzle {
	p := make(Puzzle, width)
	for y := 0; y < width; y++ {
		p[y] = make([]Tile, length)
		for x := range p[y] {
			p[y][x] = NewTile(0)
		}
	}
	return p
}

func (p Puzzle) Tile(x, y int) *Tile {
	return &p[y][x]
}

func (p Puzzle) SetTile(x, y int, t *Tile) {
	p[y][x] = *t
}

func (p Puzzle) value(x, y int) int {
	return p[y][x].Value
}

func (p Puzzle) setValue(x, y, val int) {
	p[y][x].Value = val
}

func (p Puzzle) IsValidRow(y int) bool {
	return p.isValidRange(0, y, boardDiameter-1, y)
}

func (p Puzzle) IsValidCol(x int) bool {
	return p.isValidRange(x, 0, x, boardDiameter-1)
}

func (p Puzzle) IsValidSquare(x, y int) bool {
	topLeftX := (x / squareDiameter) * squareDiameter
	topLeftY := (y / squareDiameter) * squareDiameter
	botRightX := topLeftX + squareDiameter - 1
	botRightY := topLeftY + squareDiameter - 1
	return p.isValidRange(topLeftX, topLeftY, botRightX, botRightY)
}

func (p Puzzle) IsValid(x, y int) bool {
	return p.IsValidSquare(x, y) && p.IsValidRow(y) && p.IsValidCol(x)
}

func (p Puzzle) isValidRange(minX, minY, maxX, maxY int) bool {
	existingDigits := make(map[int]struct{})
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			value := p.Tile(x, y).Value
			if value == blank {
				continue
			}
			_, present := existingDigits[value]
			if present {
				return false
			}
			existingDigits[value] = struct{}{}
		}
	}
	return true
}

func (p Puzzle) solve() bool {
	for x := 0; x < boardDiameter; x++ {
		for y := 0; y < boardDiameter; y++ {
			// Skip tiles with assigned values
			if p.value(x, y) != blank {
				continue
			}
			// Try 1-9 in an open tile
			for i := 1; i <= boardDiameter; i++ {
				p.setValue(x, y, i)
				if p.IsValid(x, y) {
					// If we've reached the max x and y, we have found a
					// solution (this assumes a rectangular puzzle)
					if x == len(p[0])-1 && y == len(p)-1 {
						return true
					}
					// For valid guesses, attempt to search further
					if p.solve() {
						return true
					}
				}
				p.setValue(x, y, blank)
			}
			// If the values 1 to 9 are tried with no success, the puzzle is invalid
			return false
		}
	}
	return false
}

func (p Puzzle) String() string {
	var result string
	for _, row := range p {
		for _, tile := range row {
			result += fmt.Sprintf("%d ", tile.Value)
		}
		result += "\n"
	}
	return result
}

// LoadPuzzle reads a QQWing formatted Sudoku board from a text file into a
// Puzzle value in memory
func LoadPuzzle(path string) (Puzzle, error) {
	puzz := NewPuzzle(boardDiameter, boardDiameter)
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return puzz, fmt.Errorf("Fatal error opening QQWing puzzle text file: %w", err)
	}
	lines := bufio.NewScanner(file)
	tileY := 0
	for lines.Scan() {
		line := lines.Bytes()
		if tileY >= len(puzz) {
			return puzz, errors.New("Inputted puzzle has too many rows")
		}
		tileX := 0
		for _, char := range line {
			if tileX >= len(puzz[0]) {
				return puzz, errors.New("Inputted puzzle has too many columns")
			}
			var value int
			if char == '.' {
				value = blank
			} else if char < '0' || char > '9' {
				continue
			} else {
				value, err = strconv.Atoi(string(char))
				if err != nil {
					return puzz, fmt.Errorf(
						"Error converting ASCII character to a valid Sudoku value: %w",
						err,
					)
				}
			}
			tile := NewTile(value)
			puzz.SetTile(tileX, tileY, &tile)
			tileX++
		}
		if tileX > 0 {
			tileY++
		}
	}
	if lines.Err() != nil {
		return puzz, fmt.Errorf("Fatal error reading QQWing puzzle text file: %w", lines.Err())
	}
	return puzz, nil
}
