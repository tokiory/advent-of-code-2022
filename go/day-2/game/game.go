package game

import "fmt"

type Round struct {
	Player   GameShapeCharacter
	Opponent GameShapeCharacter
}

const asciiDifference = 23

func GetRoundScore(round Round) int {
	playerScore := int(ShapeMap[round.Player])

	// Draw
	if rune(round.Player[0])-asciiDifference == rune(round.Opponent[0]) {
		return playerScore + int(DRAW)
	}

	gameState := LOSE

	if (round.Player == PLAYER_ROCK_SHAPE && round.Opponent == OPPONENT_SCISSORS_SHAPE) ||
		round.Player == PLAYER_PAPER_SHAPE && round.Opponent == OPPONENT_ROCK_SHAPE ||
		round.Player == PLAYER_SCISSORS_SHAPE && round.Opponent == OPPONENT_PAPER_SHAPE {
		fmt.Println(round.Player, "over", round.Opponent, "- win")
		gameState = WIN
	}

	return int(gameState) + playerScore
}
