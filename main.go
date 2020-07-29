package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	width  = 630
	height = width
	// A Sudoku board is 9x9
	boardDiameter = 9
)

type Game struct {
	Board [][]Tile
}

func NewGame() *Game {
	g := &Game{}
	g.Board = make([][]Tile, boardDiameter)
	for i := 0; i < boardDiameter; i++ {
		g.Board[i] = make([]Tile, boardDiameter)
		for j := range g.Board[i] {
			g.Board[i][j] = NewTile(0)
		}
	}
	return g
}

func (g *Game) DrawBoard(screen *ebiten.Image) {
	slide := float64(tileDiameter - borderWidth)
	xSlide := 0.0
	translation := &ebiten.DrawImageOptions{}
	for _, row := range g.Board {
		for _, tile := range row {
			screen.DrawImage(tile.Image(), translation)
			translation.GeoM.Translate(slide, 0)
			xSlide += slide
		}
		translation.GeoM.Translate(-xSlide, slide)
		xSlide = 0.0
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawBoard(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width - excess, height - excess
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(width-excess, height-excess)
	ebiten.SetWindowTitle("Sudoku")
	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
