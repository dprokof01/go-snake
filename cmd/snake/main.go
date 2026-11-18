package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dprokof01/go-snake/internal/config"
)

// TODO: extract it into a separate package internal/snake m.b?
type Point struct {
	X int
	Y int
}

// TODO: extract it into a separate package internal/snake m.b?
func NewSnake(cfg *config.Config) Point {
	return Point{
		X: cfg.Size.Col / 2,
		Y: cfg.Size.Row / 2,
	}
}

func main() {
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	snake := NewSnake(cfg)

	var lose bool

	for !lose {
		snake.X++

		fmt.Println("\033[H")

		switch {
		case cfg.Size.Col-1 == snake.X || cfg.Size.Row-1 == snake.Y || snake.X == 0 || snake.Y == 0:
			lose = true
			fmt.Println("You lose!")
		default:
			Render(cfg, snake)
		}

		time.Sleep(200 * time.Millisecond)
	}

	// TODO: add Raw Terminal Mode for handling inputs without enter
	// use termios

	// TODO: add game loop

	// TODO: async handling inputs with Goroutines and channels

	// TODO: control the speed of the game with time.Ticker

}

func Render(cfg *config.Config, snake Point) {
	row, col := cfg.Size.Row, cfg.Size.Col
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			switch {
			case i == 0 || i == row-1 || j == 0 || j == col-1:
				fmt.Print("#")
			case i == snake.Y && j == snake.X:
				fmt.Print("@")
			default:
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}
