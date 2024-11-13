package main

import "fmt"

func GetName() string {
	name := ""  // or var name string
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("welcome %s, let's play\n", name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint 
	for {
		fmt.Printf("enter your bet or 0 to quit (balance = %d): ", balance)
		fmt.Scan(&bet)
		if bet > balance {
			fmt.Println("bet is too high")
		} else {
			break
		}
	}
	return bet
}

func CheckWin(spin [][]string, multipliers map[string]uint) []uint {
	var lines []uint

	for _, row := range spin {
		win := true
		checkSymbol := row[0]
		for _, symbol := range row[1:] {
			if checkSymbol != symbol {
				win = false
				break
			}
		}
		if win {
			lines = append(lines, multipliers[checkSymbol])
		} else {
			lines = append(lines, 0)
		}
	}
	return lines
}
