package game

import "fmt"

type Round struct {
    Player   GameShapeCharacter
    Opponent GameShapeCharacter
}

type RoundAlternative struct {
    Opponent GameShapeCharacter
    Strategy GameStateCharacter
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
        gameState = WIN
    }

    return int(gameState) + playerScore
}

// Part 2
func GetAlternativeRoundScore(round RoundAlternative) int {
    var playerShape GameShapeCharacter

    switch round.Strategy {
    case DRAW_CHARACTER:
        playerShape = getDrawCharacter(round.Opponent)
    case LOSE_CHARACTER:
        playerShape = getLooseCharacter(round.Opponent)
    default:
        playerShape = getWinCharacter(round.Opponent)
    }

    fmt.Println("---")
    fmt.Println("Strategy is", round.Strategy)
    fmt.Println("Player shape is", playerShape)
    fmt.Println("Opponent shape is", round.Opponent)
    fmt.Println("---")

    return int(GameStateMap[round.Strategy]) + int(ShapeMap[playerShape])
}

func getLooseCharacter(opponent GameShapeCharacter) GameShapeCharacter {
    switch opponent {
    case OPPONENT_ROCK_SHAPE:
        return PLAYER_SCISSORS_SHAPE
    case OPPONENT_SCISSORS_SHAPE:
        return PLAYER_PAPER_SHAPE
    default:
        return PLAYER_ROCK_SHAPE
    }
}

func getWinCharacter(opponent GameShapeCharacter) GameShapeCharacter {
    switch opponent {
    case OPPONENT_ROCK_SHAPE:
        return PLAYER_PAPER_SHAPE
    case OPPONENT_SCISSORS_SHAPE:
        return PLAYER_ROCK_SHAPE
    default:
        return PLAYER_SCISSORS_SHAPE
    }
}

func getDrawCharacter(opponent GameShapeCharacter) GameShapeCharacter {
    switch opponent {
    case OPPONENT_ROCK_SHAPE:
        return PLAYER_ROCK_SHAPE
    case OPPONENT_SCISSORS_SHAPE:
        return PLAYER_SCISSORS_SHAPE
    default:
        return PLAYER_PAPER_SHAPE
    }
}
