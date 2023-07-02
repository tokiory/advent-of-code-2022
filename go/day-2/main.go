package main

import (
	"fmt"
	"main/cli"
	"main/game"
	"main/parser"
)

func main() {
	fmt.Println("Hello, World!")
	path := cli.GetPath()
	content := parser.ReadInput(path)
	rounds := parser.Parse(content)

	var score int

	for _, round := range rounds {
		score += game.GetRoundScore(round)
	}

	fmt.Println("Total score is:", score)
}
