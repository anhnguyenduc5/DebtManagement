package main

import (
	"fmt"
	"math/rand"
	"time"
)

var players []*Player
var capitalOfEach int

type Player struct {
	name           string
	totalMoneyGet  int
	totalRoundsWon int
	finalMoney     int
	debts          map[string]int
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numberPlayer, betAmount := inputPlayersInfo()
	calculate(betAmount, numberPlayer)
	printResults()
}

func calculate(betAmount int, numberPlayer int) {
	var totalRoundsPlayed int
	for _, player := range players {
		totalRoundsPlayed += player.totalRoundsWon
	}
	capitalOfEach = totalRoundsPlayed * betAmount
	for i := 0; i < numberPlayer; i++ {
		players[i].finalMoney = (capitalOfEach - players[i].totalMoneyGet) * -1
	}
	for i := 0; i < numberPlayer; i++ {
		for j := i + 1; j < numberPlayer; j++ {
			if players[i].finalMoney < 0 && players[j].finalMoney > 0 {
				debt := findMin(-players[i].finalMoney, players[j].finalMoney)
				players[i].debts[players[j].name] = debt
				players[i].finalMoney += debt
				players[j].finalMoney -= debt
			} else if players[i].finalMoney > 0 && players[j].finalMoney < 0 {
				debt := findMin(players[i].finalMoney, -players[j].finalMoney)
				players[j].debts[players[i].name] = debt
				players[i].finalMoney -= debt
				players[j].finalMoney += debt
			}
		}
	}
	for i := 0; i < numberPlayer; i++ {
		players[i].finalMoney = (capitalOfEach - players[i].totalMoneyGet) * -1
	}
}

func inputPlayersInfo() (int, int) {
	var numberPlayer, betAmount int
	fmt.Println("Input number of player: ")
	fmt.Scan(&numberPlayer)
	fmt.Println("Input amount of bet: ")
	fmt.Scan(&betAmount)
	players = make([]*Player, numberPlayer)
	for i := 0; i < numberPlayer; i++ {
		fmt.Printf("Input name of player %d: ", i+1)
		var name string
		fmt.Scan(&name)
		fmt.Printf("Input total rounds player: %s WON \n ", name)
		var roundsWon int
		fmt.Scan(&roundsWon)
		players[i] = &Player{
			name:           name,
			totalRoundsWon: roundsWon,
			totalMoneyGet:  roundsWon * betAmount * numberPlayer,
			debts:          make(map[string]int),
		}
	}
	return numberPlayer, betAmount
}
func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func printResults() {
	fmt.Println("\nFinal results:")
	for _, player := range players {
		fmt.Printf("Player %s spent: %d $, gain: %d $, final money: %d$\n",
			player.name, capitalOfEach, player.totalMoneyGet, player.finalMoney)

		if len(player.debts) != 0 {
			for debtor, amount := range player.debts {
				fmt.Printf(" %s need to pay to %s: %d $,", player.name, debtor, amount)
			}
			fmt.Println()
		}
	}
}
