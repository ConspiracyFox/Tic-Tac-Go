package main

import "fmt"

func main() {

	//the game board -> empty Strings

	ticBoard := [3][3]string{}

	// var to carry current player name

	player := "x"

	for {

		fmt.Println("player", player)
		fmt.Println(ticBoard[0])
		fmt.Println(ticBoard[1])
		fmt.Println(ticBoard[2])

		//read row value

		var row int
		fmt.Println("Enter a number 0, 1 or 2")
		fmt.Scanln(&row)

		if row < 0 || row > 2 {
			fmt.Println("Invalid number, please enter again: 0, 1 or 2")
			continue
		}

		//read column value

		var column int
		fmt.Println("Enter a number 0, 1 or 2")
		fmt.Scanln(&column)

		if column < 0 || column > 2 {
			fmt.Println("Invalid number, please enter again: 0, 1 or 2")
			continue
		}

		//set value into game ticBoard
		if ticBoard[row][column] == "" {
			ticBoard[row][column] = player
		} else {
			//index already has value
			fmt.Println("invalid entry: ", row, column, "value ", ticBoard[row][column])
			continue

		}

		//Display the Board (after each move)

		//chech if someone won
		for i := 0; i < 3; i++ {
			if ticBoard[i][0] == ticBoard[i][1] && ticBoard[i][1] == ticBoard[i][2] && ticBoard[i][2] == player ||
				ticBoard[0][i] == ticBoard[1][i] && ticBoard[1][i] == ticBoard[2][i] && ticBoard[2][i] == player {
				fmt.Println("Game over, winner is player ", player)
				return
			}
		}

		if ticBoard[0][0] == ticBoard[1][1] && ticBoard[1][1] == ticBoard[2][2] && ticBoard[2][2] == player ||
			ticBoard[0][2] == ticBoard[1][1] && ticBoard[1][1] == ticBoard[2][0] && ticBoard[2][0] == player {
			fmt.Println("Game over, winner is player ", player)
		}

		//Swap players

		if player == "x" {
			player = "o"
		} else {
			player = "x"
		}

	}

}
