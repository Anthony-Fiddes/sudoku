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
				value = -1
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
