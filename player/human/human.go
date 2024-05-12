package human

import (
	"bufio"
	"fmt"
	"rock-paper-scissors/game"
)

type HumanPlayer struct {
	Name    string
	Scanner *bufio.Scanner
}

func (r *HumanPlayer) GetName() string {
	return r.Name
}

func (r *HumanPlayer) GetNextMove() game.Move {
	var isValid bool
	var input string
	for !isValid {
		fmt.Printf("What do you want to throw?\nPlease enter one of %+v \n>", game.ValidMoves)
		r.Scanner.Scan()
		input = r.Scanner.Text()
		isValid = game.ValidateMove(input)
	}
	return game.Move(input)
}
