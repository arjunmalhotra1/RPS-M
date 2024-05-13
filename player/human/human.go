package human

import (
	"bufio"
	"fmt"
	"rock-paper-scissors/game"
)

type HumanPlayer struct {
	Name        string
	Scanner     *bufio.Scanner
	CurrentMove game.Move
}

// // getName implements player.Player.
// func (r *HumanPlayer) getName() string {
// 	panic("unimplemented")
// }

// // getNextMove implements player.Player.
// func (r *HumanPlayer) getNextMove() game.Move {
// 	panic("unimplemented")
// }

func (r *HumanPlayer) GetName() string {
	return r.Name
}

func (r *HumanPlayer) GetCurrentMove() game.Move {
	return r.CurrentMove
}

func (r *HumanPlayer) GetNextMove() game.Move {
	//fmt.Print("What do you want to throw?\n> ")
	var isValid bool
	var input string
	for !isValid {
		fmt.Printf("What do you want to throw?\nPlease enter one of %+v \n>", game.ValidMoves)
		r.Scanner.Scan()
		input = r.Scanner.Text()
		isValid = game.ValidateMove(input)
	}
	r.CurrentMove = game.Move(input)
	return game.Move(input)
}
