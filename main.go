package main

import (
	"bufio"
	"fmt"
	"os"
	"rock-paper-scissors/game"
	"rock-paper-scissors/player"
	"rock-paper-scissors/player/computer"
	"rock-paper-scissors/player/human"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Let's play Rock, Paper, Scissors. What is your name?\n> ")
	scanner.Scan()
	playerName := scanner.Text()

	currentGame := &game.Game{
		HumanPlayerScore:    0,
		ComputerPlayerScore: 0,
	}

	var humanPlayer player.Player = &human.HumanPlayer{
		Name:    playerName,
		Scanner: scanner,
	}

	var computerPlayer = &computer.ComputerPlayer{
		Name: "Robot",
	}

	for {
		fmt.Println("")

		humanPlayerPlay := game.Play{
			Name: humanPlayer.GetName(),
			Move: humanPlayer.GetNextMove(),
		}

		computerPlayerPlay := game.Play{
			Name: computerPlayer.GetName(),
			Move: computerPlayer.GetNextMove(),
		}

		fmt.Printf("%s throws %s\n", humanPlayerPlay.Name, humanPlayerPlay.Move)
		fmt.Printf("%s throws %s\n", computerPlayerPlay.Name, computerPlayerPlay.Move)

		winner := game.WinnerIs(humanPlayerPlay, computerPlayerPlay)
		switch winner {
		case humanPlayer.GetName():
			fmt.Printf("%s beats %s, %s wins!\n", humanPlayerPlay.Move, computerPlayerPlay.Move, humanPlayerPlay.Name)
			currentGame.HumanPlayerScore++
			computerPlayer.SetPrevBattleResult("lost")
		case computerPlayer.GetName():
			fmt.Printf("%s beats %s, %s wins!\n", computerPlayerPlay.Move, humanPlayerPlay.Move, computerPlayerPlay.Name)
			currentGame.ComputerPlayerScore++
			computerPlayer.SetPrevBattleResult("win")
		default:
			fmt.Println("It's a draw!")
		}

		computerPlayer.SetPrevHumanMove(humanPlayerPlay.Move)
		computerPlayer.SetPrevMove(computerPlayerPlay.Move)

	}
}
