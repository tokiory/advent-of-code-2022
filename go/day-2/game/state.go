package game

type GameState int
type GameStateCharacter string

const (
    LOSE GameState = iota * 3
    DRAW
    WIN
)

const (
    LOSE_CHARACTER GameStateCharacter = "X"
    DRAW_CHARACTER GameStateCharacter = "Y"
    WIN_CHARACTER  GameStateCharacter = "Z"
)

var GameStateMap = map[GameStateCharacter]GameState{
    LOSE_CHARACTER: LOSE,
    DRAW_CHARACTER: DRAW,
    WIN_CHARACTER:  WIN,
}
