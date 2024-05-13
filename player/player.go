package player

import "rock-paper-scissors/game"

type Player interface {
	GetName() string
	GetNextMove() game.Move
}
