package game

import "slices"

type Move string

const (
	Rock     Move = "rock"
	Paper    Move = "paper"
	Scissors Move = "scissors"
	//Draw     Move = "draw"
)

type Game struct {
	HumanPlayerScore    int
	ComputerPlayerScore int
	// HumanPreviousMove    Move
	// ComputerPreviousMove Move
	// h HumanPlayer
	// c computerPlayer
}

type Play struct {
	Name string
	Move Move
}

// func (g Game) GetHumanPreviousMove() Move {
// 	return g.HumanPreviousMove
// }

// func (g Game) GetComputerPreviousMove() Move {
// 	return g.ComputerPreviousMove
// }

var ValidMoves = []string{string(Rock), string(Paper), string(Scissors)}

func ValidateMove(m string) bool {
	return slices.Contains(ValidMoves, m)
}

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

// winnerIs determines the name of the winning player, given each player's Move. In the case of a tie, an empty string
// is returned.
func WinnerIs(play1, play2 Play) string {
	if play1.Move == play2.Move {
		return "draw"
	}

	switch play1.Move {
	case "rock":
		switch play2.Move {
		case "paper":
			return play2.Name
		case "scissors":
			return play1.Name
		}
	case "paper":
		switch play2.Move {
		case "scissors":
			return play2.Name
		case "rock":
			return play1.Name
		}
	case "scissors":
		switch play2.Move {
		case "rock":
			return play2.Name
		case "paper":
			return play1.Name
		}
	}
	return "draw"
}
