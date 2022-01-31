package main

import "fmt"
import "strings"

func printBoardPlay() {
  ticBoard := [9]string{}
	player := "x"

	for {
		fmt.Println(" ",ticBoard[0], " | ", ticBoard[1], "|", ticBoard[2])
		fmt.Println(" ",ticBoard[3], " | ", ticBoard[4], "|", ticBoard[5])
		fmt.Println(" ",ticBoard[6], " | ", ticBoard[7], "|", ticBoard[8])
		var field int
		fmt.Println("Pick a field")
		fmt.Scanln(&field)

		if field < 0 || field > 8 {
			fmt.Println("Invalid field enter again")
			continue
		}
		if ticBoard[field] == "" {
			ticBoard[field] = strings.Replace(player,"","",100)
		} else {
			fmt.Println("Invalid field enter again")
			continue
		}

		if ticBoard[0] == player {
			if ticBoard[1] == player && ticBoard[2] == player {
				fmt.Println(player, " won")
				return
			}
			if ticBoard[3] == player && ticBoard[6] == player {
				fmt.Println(player, " won")
				return
			}
			if ticBoard[4] == player && ticBoard[8] == player {
				fmt.Println(player, " won")
				return
			}
		} else if ticBoard[8] == player {
			if ticBoard[2] == player && ticBoard[5] == player {
				fmt.Println(player, " won")
				return
			}
			if ticBoard[6] == player && ticBoard[7] == player {
				fmt.Println(player, " won")
				return
			}
		} else if ticBoard[4] == player {
			if ticBoard[1] == player && ticBoard[7] == player {
				fmt.Println(player, " won")
				return
			}
			if ticBoard[3] == player && ticBoard[5] == player {
				fmt.Println(player, " won")
				return
			}
			if ticBoard[2] == player && ticBoard[6] == player {
				fmt.Println(player, " won")
				return
			}
		}
		if player == "x" {
			player = "o"
		} else {
			player = "x"
		}
	}
}

func main() {
	printBoardPlay()

}
