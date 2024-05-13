package main

import "testing"

func Test_winnerIs(t *testing.T) {
	type args struct {
		play1 play
		play2 play
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "rock-vs-scissors",
			args: args{
				play1: play{name: "player1", move: rock},
				play2: play{name: "player2", move: scissors},
			},
			want: "player1",
		},
		{
			name: "rock-vs-paper",
			args: args{
				play1: play{name: "player1", move: rock},
				play2: play{name: "player2", move: paper},
			},
			want: "player2",
		},
		{
			name: "rock-vs-rock",
			args: args{
				play1: play{name: "player1", move: rock},
				play2: play{name: "player2", move: rock},
			},
			want: "draw",
		},
		{
			name: "scissors-vs-paper",
			args: args{
				play1: play{name: "player1", move: scissors},
				play2: play{name: "player2", move: paper},
			},
			want: "player1",
		},
		{
			name: "scissors-vs-rock",
			args: args{
				play1: play{name: "player1", move: scissors},
				play2: play{name: "player2", move: rock},
			},
			want: "player2",
		},
		{
			name: "scissors-vs-scissors",
			args: args{
				play1: play{name: "player1", move: scissors},
				play2: play{name: "player2", move: scissors},
			},
			want: "draw",
		},
		{
			name: "paper-vs-rock",
			args: args{
				play1: play{name: "player1", move: paper},
				play2: play{name: "player2", move: rock},
			},
			want: "player1",
		},
		{
			name: "paper-vs-scissors",
			args: args{
				play1: play{name: "player1", move: paper},
				play2: play{name: "player2", move: scissors},
			},
			want: "player2",
		},
		{
			name: "paper-vs-paper",
			args: args{
				play1: play{name: "player1", move: paper},
				play2: play{name: "player2", move: paper},
			},
			want: "draw",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winnerIs(tt.args.play1, tt.args.play2); got != tt.want {
				t.Errorf("winnerIs() = %v, want %v", got, tt.want)
			}
		})
	}
}
