package main

import (
	"rock-paper-scissors/game"
	"testing"
)

func Test_winnerIs(t *testing.T) {
	type args struct {
		play1 game.Play
		play2 game.Play
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "rock-vs-scissors",
			args: args{
				play1: game.Play{Name: "player1", Move: game.Rock},
				play2: game.Play{Name: "player2", Move: game.Scissors},
			},
			want: "player1",
		},
		{
			name: "rock-vs-paper",
			args: args{
				play1: game.Play{Name: "player1", Move: game.Rock},
				play2: game.Play{Name: "player2", Move: game.Paper},
			},
			want: "player2",
		},
		{
			name: "rock-vs-rock",
			args: args{
				play1: game.Play{Name: "player1", Move: game.Rock},
				play2: game.Play{Name: "player2", Move: game.Rock},
			},
			want: "draw",
		},
		{
			name: "scissors-vs-paper",
			args: args{
				play1: game.Play{Name: "player1", Move: game.Scissors},
				play2: game.Play{Name: "player2", Move: game.Paper},
			},
			want: "player1",
		},
		{
			name: "scissors-vs-rock",
			args: args{
				play1: game.Play{Name: "player1", Move: game.Scissors},
				play2: game.Play{Name: "player2", Move: game.Rock},
			},
			want: "player2",
		},
		{
			name: "scissors-vs-scissors",
			args: args{
				play1: game.Play{Name: "player1", Move: game.Scissors},
				play2: game.Play{Name: "player2", Move: game.Scissors},
			},
			want: "draw",
		},
		{
			name: "paper-vs-rock",
			args: args{
				play1: game.Play{Name: "player1", Move: game.Paper},
				play2: game.Play{Name: "player2", Move: game.Rock},
			},
			want: "player1",
		},
		{
			name: "paper-vs-scissors",
			args: args{
				play1: game.Play{Name: "player1", Move: game.Paper},
				play2: game.Play{Name: "player2", Move: game.Scissors},
			},
			want: "player2",
		},
		{
			name: "paper-vs-paper",
			args: args{
				play1: game.Play{Name: "player1", Move: game.Paper},
				play2: game.Play{Name: "player2", Move: game.Paper},
			},
			want: "draw",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := game.WinnerIs(tt.args.play1, tt.args.play2); got != tt.want {
				t.Errorf("winnerIs() = %v, want %v", got, tt.want)
			}
		})
	}
}
