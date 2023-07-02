package parser

import (
	"main/game"
	"os"
	"strings"
)

func ReadInput(path string) string {
	byteContent, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(byteContent)
}

func Parse(rawContent string) []game.Round {
	lines := strings.Split(rawContent, "\n")
	rounds := make([]game.Round, 0, len(lines))

	for _, line := range lines {
		rawCharacters := strings.Split(line, " ")

		round := game.Round{
			Opponent: game.GameShapeCharacter(rawCharacters[0]),
			Player:   game.GameShapeCharacter(rawCharacters[1]),
		}

		rounds = append(rounds, round)
	}

	return rounds
}
