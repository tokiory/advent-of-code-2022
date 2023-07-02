package game

type GameShapeCharacter string
type GameShapeScore int

const (
	ROCK_SCORE GameShapeScore = iota + 1
	PAPER_SCORE
	SCISSORS_SCORE
)

const (
	OPPONENT_ROCK_SHAPE     GameShapeCharacter = "A"
	OPPONENT_PAPER_SHAPE    GameShapeCharacter = "B"
	OPPONENT_SCISSORS_SHAPE GameShapeCharacter = "C"
)

const (
	PLAYER_ROCK_SHAPE     GameShapeCharacter = "X"
	PLAYER_PAPER_SHAPE    GameShapeCharacter = "Y"
	PLAYER_SCISSORS_SHAPE GameShapeCharacter = "Z"
)

var ShapeMap = map[GameShapeCharacter]GameShapeScore{
	PLAYER_ROCK_SHAPE:     ROCK_SCORE,
	PLAYER_PAPER_SHAPE:    PAPER_SCORE,
	PLAYER_SCISSORS_SHAPE: SCISSORS_SCORE,
}
