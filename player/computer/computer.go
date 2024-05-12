package computer

import (
	"fmt"
	"math/rand"
	"rock-paper-scissors/game"
	"time"
)

type ComputerPlayer struct {
	Name             string
	PrevBattleResult string
	prevHumanMove    game.Move
	prevMove         game.Move
}

func (p *ComputerPlayer) GetName() string {
	return p.Name
}

func (p *ComputerPlayer) SetPrevHumanMove(m game.Move) {
	p.prevHumanMove = m
}

func (p *ComputerPlayer) GetPrevHumanMove() game.Move {
	return p.prevHumanMove
}

func (p *ComputerPlayer) SetPrevMove(m game.Move) {
	p.prevMove = m
}

func (p *ComputerPlayer) GetPrevMove() game.Move {
	return p.prevMove
}
func (p *ComputerPlayer) SetPrevBattleResult(result string) {
	p.PrevBattleResult = result
}

func (p *ComputerPlayer) GetPrevBattleResult() string {
	return p.PrevBattleResult
}

// GetNextMove() returns the next move of the ComputerPlayer
func (p *ComputerPlayer) GetNextMove() game.Move {
	switch p.GetPrevBattleResult() {
	case "win":
		return game.BeatMove(p.GetPrevMove())
	case "lost":
		return game.BeatMove(p.GetPrevHumanMove())
	default:
		fmt.Println("In default")
		r := rand.New(rand.NewSource(time.Now().Unix()))
		return game.Move(game.ValidMoves[r.Intn(len(game.ValidMoves))])
	}

}
