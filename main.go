package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	width         = 630
	height        = width
	boardDiameter = 9
)

type Game struct {
	Board [][]int
}

func NewGame() *Game {
	// A Sudoku board is 9x9
	g := &Game{}
	g.Board = make([][]int, boardDiameter)
	for i := 0; i < 9; i++ {
		g.Board[i] = make([]int, boardDiameter)
	}
	return g
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	gray := color.Gray{100}
	tile, _ := ebiten.NewImageFromImage(
		ColoredTile(50, 1, gray, color.White),
		ebiten.FilterDefault,
	)
	screen.DrawImage(tile, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width / 2, height / 2
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Sudoku")
	// Call ebiten.RunGame to start your game loop.
	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
