package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dprokof01/go-snake/internal/config"
	"gopkg.in/yaml.v3"
)

func main() {
	// TODO: move it to func LoadConfig
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	cfg := &config.Config{}
	if err = yaml.Unmarshal(data, cfg); err != nil {
		fmt.Printf("err parse: %d", err)
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
