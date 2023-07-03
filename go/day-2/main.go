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

    fmt.Println("\nPart 2: Alternative strategy")

    var alternativeScore int

    for _, round := range rounds {
        alternativeScore += game.GetAlternativeRoundScore(
            game.RoundAlternative{
                Opponent: round.Opponent,
                Strategy: game.GameStateCharacter(round.Player),
            },
        )
    }

    fmt.Println("Alternative score is:", alternativeScore)
}
