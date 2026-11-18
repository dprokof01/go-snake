package main

import (
	"fmt"
	"log"

	"github.com/dprokof01/go-snake/internal/config"
)

func main() {
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var row, col = cfg.Size.Row, cfg.Size.Col

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 || i == row-1 || j == 0 || j == col-1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}

	// TODO: add snake and move it, add render and fps
	// ANSI Escape Sequences clean terminal, color and move cursor.

	// TODO: add Raw Terminal Mode for handling inputs without enter
	// use termios

	// TODO: add game loop

	// TODO: async handling inputs with Goroutines and channels

	// TODO: control the speed of the game with time.Ticker

}
