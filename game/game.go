package game

import "slices"

type Move string

const WinningScore int = 7

const (
	Rock     Move = "rock"
	Paper    Move = "paper"
	Scissors Move = "scissors"
)

type Game struct {
	HumanPlayerScore    int
	ComputerPlayerScore int
}

type Play struct {
	Name string
	Move Move
}

var ValidMoves = []string{string(Rock), string(Paper), string(Scissors)}

func ValidateMove(m string) bool {
	return slices.Contains(ValidMoves, m)
}

// BeatMove returns the move that beats the 'm' input move.
func BeatMove(m Move) Move {
	switch m {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	default:
		return Move("")
	}
}

// WinnerIs determines the name of the winning player, given each player's Move. In the case of a tie, an empty string
// is returned.
func WinnerIs(play1, play2 Play) string {
	if play1.Move == play2.Move {
		return "draw"
	}

	if BeatMove(play1.Move) == play2.Move {
		return play2.Name
	}

	return play1.Name
}
