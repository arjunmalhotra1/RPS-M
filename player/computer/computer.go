package computer

import (
	"math/rand"
	"rock-paper-scissors/game"
	"time"
)

type ComputerPlayer struct {
	Name             string
	PrevBattleResult string
	prevHumanMove    game.Move
	prevMove         game.Move
	//currGame *game
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

/*
Therefore, this is the best way to win at rock-paper-scissors: if you lose the first round, switch to the thing that beats the thing your opponent just played. If you win, don't keep playing the same thing, but instead switch to the thing that would beat the thing that you just played. In other words, play the hand your losing opponent just played. To wit: you win a round with rock against someone else's scissors. They are about to switch to paper. You should switch to scissors. Got it? Good.
*/
func (p *ComputerPlayer) GetNextMove() game.Move {
	switch p.GetPrevBattleResult() {
	case "win":
		return game.BeatMove(p.GetPrevMove())
	case "lose":
		return game.BeatMove(p.GetPrevHumanMove())
	default:
		r := rand.New(rand.NewSource(time.Now().Unix()))
		return game.Move(game.ValidMoves[r.Intn(len(game.ValidMoves))])
	}

}
