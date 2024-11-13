package main

import (
	"fmt"
)

func main()  {
	symbols := map[string]uint{
		"A": 4,
		"B": 7,
		"C": 12,
		"D": 20,
	}
	multipliers := map[string]uint{
		"A": 20,
		"B": 10,
		"C": 5,
		"D": 2,
	}

	symbolArr := GenerateSymbolArray(symbols)

	balance := uint(200)
	GetName()

	for balance > 0 {
		bet := GetBet(balance)
		if bet == 0 {
			break
		}
		balance -= bet
		spin := GetSpin(symbolArr, 3, 3)
		PrintSpin(spin)
		winningLines := CheckWin(spin, multipliers)
		for i, multi := range winningLines {
			win := multi * bet
			balance += win
			if multi > 0 {
				fmt.Printf("you won (%dx) %d$ on line %d\n", multi, win, i+1)
			}
		}
	}
	fmt.Printf("you left with: %d.\n", balance)
}