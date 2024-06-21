package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var reader = bufio.NewReader(os.Stdin)
var players []*Player
var capitalOfEach int

type Player struct {
	name           string
	totalMoneyGet  int
	totalRoundsWon int
	finalMoney     int
	debts          string
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var numberPlayer, betAmount int
	fmt.Println("Input number of player: ")
	fmt.Scan(&numberPlayer)
	fmt.Println("Input amount of bet: ")
	fmt.Scan(&betAmount)
	players = make([]*Player, numberPlayer)
	// input name
	//tinh so tien von moi nguoi can dat capitalOfEach /totalMoneySpent = tong so van choi * betamount
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
		}
	}
	var totalRoundsPlayed int
	for _, player := range players {
		totalRoundsPlayed += player.totalRoundsWon
	}
	// tong so von cua moi nguoi
	capitalOfEach = totalRoundsPlayed * betAmount
	for i := 0; i < numberPlayer; i++ {
		players[i].finalMoney = (capitalOfEach - players[i].totalMoneyGet) * -1
	}
	// loop qua: loop 1: lay finalMoney của player 1
	// loop 2: so sánh giá trị của player 2, nếu ai âm thì phải trả tiền cho người dương mot so tien = duong + am
	for i := 0; i < numberPlayer; i++ {
		for j := i + 1; j < numberPlayer; j++ {
			if players[i].finalMoney < 0 && players[j].finalMoney > 0 {
				debt := min(-players[i].finalMoney, players[j].finalMoney)
				players[i].debts += "\n Player " + players[i].name + " needs to pay Player " + players[j].name + " a total of: $" + strconv.Itoa(debt)
				players[i].finalMoney += debt
				players[j].finalMoney -= debt
			} else if players[i].finalMoney > 0 && players[j].finalMoney < 0 {
				debt := min(players[i].finalMoney, -players[j].finalMoney)
				players[j].debts += "\n Player " + players[j].name + " needs to pay Player " + players[i].name + " a total of: $" + strconv.Itoa(debt)
				players[i].finalMoney -= debt
				players[j].finalMoney += debt
			}
		}
	}
	for i := 0; i < numberPlayer; i++ {
		players[i].finalMoney = (capitalOfEach - players[i].totalMoneyGet) * -1
	}
	printResults()
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func printResults() {
	fmt.Println("\nFinal results:")
	for _, player := range players {
		fmt.Printf("Player %s spent: %d $, gain: %d $, final money: %d$,%s\n",
			player.name, capitalOfEach, player.totalMoneyGet, player.finalMoney, player.debts)
	}
}
