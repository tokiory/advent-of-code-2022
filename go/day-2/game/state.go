package game

type GameState int

const (
	LOSE GameState = iota * 3
	DRAW
	WIN
)
