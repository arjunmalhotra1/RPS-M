package computer

import (
	"rock-paper-scissors/game"
	"testing"
)

func TestGetNextMove(t *testing.T) {
	// type args struct {
	// }
	tests := []struct {
		name string
		cp   ComputerPlayer
		want game.Move
	}{
		{
			name: "prevBattle-win",
			cp: ComputerPlayer{
				PrevBattleResult: "win",
				prevHumanMove:    game.Move("paper"),
				prevMove:         game.Move("scissors"),
			},
			want: game.Move("rock"),
		},
		{
			name: "prevBattle-lost",
			cp: ComputerPlayer{
				PrevBattleResult: "lost",
				prevHumanMove:    game.Move("paper"),
				prevMove:         game.Move("rock"),
			},
			want: game.Move("scissors"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cp.GetNextMove(); got != tt.want {
				t.Errorf("GetNextMove() = %s, want %s", got, tt.want)
			}
		})
	}

}
