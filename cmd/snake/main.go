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
func NewSnake(row, col int) Point {
	return Point{
		X: col / 2,
		Y: row / 2,
	}
}

func main() {
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	row, col := cfg.Size.Row, cfg.Size.Col
	snake := NewSnake(row, col)
	dir := Point{
		X: 1,
		Y: 0,
	}

	for {

		switch {
		case dir.X == 1 && snake.X >= col-2:
			dir.X = 0
			dir.Y = 1
		case dir.Y == 1 && snake.Y >= row-2:
			dir.X = -1
			dir.Y = 0
		case dir.X == -1 && snake.X <= 1:
			dir.X = 0
			dir.Y = -1
		case dir.Y == -1 && snake.Y <= 1:
			dir.X = 1
			dir.Y = 0
		}

		snake.X += dir.X
		snake.Y += dir.Y

		fmt.Print("\033[H\033[2J")

		if col-1 <= snake.X || row-1 <= snake.Y || snake.X <= 0 || snake.Y <= 0 {
			fmt.Println("You lose!")
			return
		}

		Render(row, col, snake)
		time.Sleep(200 * time.Millisecond)

	}

	// TODO: add Raw Terminal Mode for handling inputs without enter
	// use termios

	// TODO: async handling inputs with Goroutines and channels

	// TODO: control the speed of the game with time.Ticker

}

func Render(row, col int, snake Point) {
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
