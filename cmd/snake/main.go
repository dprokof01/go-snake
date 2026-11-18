package main

import (
	"fmt"
	"log"

	"github.com/dprokof01/go-snake/internal/config"
)

// TODO: extract it into a separate package internal/snake m.b?
type Point struct {
	X int
	Y int
}

// TODO: extract it into a separate package internal/snake m.b?
func NewSnake(cfg *config.Config) Point {
	snake := Point{cfg.Size.Col / 2, cfg.Size.Row / 2}
	return snake
}

func main() {
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	row, col := cfg.Size.Row, cfg.Size.Col
	snake := NewSnake(cfg)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			switch {
			case i == 0 || i == row-1 || j == 0 || j == col-1:
				fmt.Print("#")
			// TODO: move snake, add render and fps
			// ANSI Escape Sequences clean terminal, color and move cursor.
			case i == snake.Y && j == snake.X:
				fmt.Print("@")
			default:
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}

	// TODO: add Raw Terminal Mode for handling inputs without enter
	// use termios

	// TODO: add game loop

	// TODO: async handling inputs with Goroutines and channels

	// TODO: control the speed of the game with time.Ticker

}
