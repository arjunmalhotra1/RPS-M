# Rock, Paper, Scissors

This is a command-line implementation of the classic game "Rock, Paper, Scissors", with the player facing off against
a computer-controlled opponent.

## Objective
Some parts of the program still need to be implemented, and other parts could use improvement. Can you take care of 
some of the items on the to-do list? Address as many or as few items as you like, and feel free to make
any other improvements!

## Upon Completion
Once you have completed the exercise, zip up the directory and email the zip file back to the team. Please take care if you use Git to not include the `.git` folder (or any other source control folders) when you zip up the project. Our security policy will remove the attachment if it deems any of the contents are hazardous.

## TODO
Here are some items that need to be completed:
1. The function `winnerIs` to determine the winner of each round still needs to be implemented. Unit tests for this function have been provided.
2. There is nothing stopping users from entering invalid values. There could be side effects from using unexpected values.
3. The computer-controlled player has a very weak strategy, always returning 'rock'. This could use improving. This could be as simple as random selection, or [a more sophisticated strategy](https://arstechnica.com/science/2014/05/win-at-rock-paper-scissors-by-knowing-thy-opponent/)
4. The game currently goes on forever. Perhaps the game should continue until one player reaches a target score, or at least give the player a means to elegantly end the game.

If you are new to Golang, here is the quick tutorial for you to pick up the language and syntax: https://go.dev/tour/welcome/1

We have included a Makefile to help you run the unit tests and start the game
- To execute unit tests, run the command:  `make test`
- To start the game, run the command: `make start`
