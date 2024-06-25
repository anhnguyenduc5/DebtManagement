package main

import (
	"testing"
)

func setup() {
	players = []*Player{
		{name: "A", totalRoundsWon: 54, totalMoneyGet: 2700000, debts: make(map[string]int)},
		{name: "B", totalRoundsWon: 21, totalMoneyGet: 1050000, debts: make(map[string]int)},
		{name: "C", totalRoundsWon: 41, totalMoneyGet: 2050000, debts: make(map[string]int)},
		{name: "D", totalRoundsWon: 32, totalMoneyGet: 1600000, debts: make(map[string]int)},
		{name: "E", totalRoundsWon: 38, totalMoneyGet: 1900000, debts: make(map[string]int)},
	}
	capitalOfEach = 300
}

func TestFindMin(t *testing.T) {
	if findMin(1, 2) != 1 {
		t.Errorf("Expected 1, got %d", findMin(1, 2))
	}
	if findMin(2, 1) != 1 {
		t.Errorf("Expected 1, got %d", findMin(2, 1))
	}
}

func TestCalculate(t *testing.T) {
	setup()
	calculate(10000, 5)

	expectedFinalMoney := []int{840000, -810000, 190000, -260000, 40000}
	for i, player := range players {
		if player.finalMoney != expectedFinalMoney[i] {
			t.Errorf("Expected final money for %s: %d, got: %d", player.name, expectedFinalMoney[i], player.finalMoney)
		}
	}

	expectedDebts := map[string]map[string]int{
		"B": {"A": 810000},
		"D": {"A": 30000, "C": 190000, "E": 40000},
	}
	for _, player := range players {
		for debtor, amount := range player.debts {
			if expectedDebts[player.name][debtor] != amount {
				t.Errorf("Expected debt from %s to %s: %d, got: %d", player.name, debtor, expectedDebts[player.name][debtor], amount)
			}
		}
	}
}
