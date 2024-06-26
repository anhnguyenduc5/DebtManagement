package main

import (
	"testing"
)

func TestCalculateFinalMoney(t *testing.T) {
	players = []*Player{
		{"Player1", 100, 5, 0, 0, make(map[string]int)},
		{"Player2", 150, 7, 0, 0, make(map[string]int)},
		{"Player3", 200, 3, 0, 0, make(map[string]int)},
	}

	calculateFinalMoney(len(players))

	for _, player := range players {
		expectedFinalMoney := (capitalOfEach - player.totalMoneyGet) * -1
		if player.finalMoney != expectedFinalMoney {
			t.Errorf("calculateFinalMoney() failed for player %s, got: %d, want: %d", player.name, player.finalMoney, expectedFinalMoney)
		}
	}
}

func TestHandleDebt(t *testing.T) {
	players = []*Player{
		{"Player1", 100, 5, -10, -10, make(map[string]int)},
		{"Player2", 150, 7, 20, 20, make(map[string]int)},
	}

	handleDebt(0, 1)

	// Check if debts were handled correctly
	expectedDebt := 10
	if players[0].debts["Player2"] != expectedDebt {
		t.Errorf("handleDebt() failed, Player1 should owe Player2 %d, got: %d", expectedDebt, players[0].debts["Player2"])
	}

	// Check if finalMoneyTemp updated correctly
	expectedFinalMoney := -10
	if players[0].finalMoney != expectedFinalMoney {
		t.Errorf("handleDebt() failed, Player1 finalMoney should be %d, got: %d", expectedFinalMoney, players[0].finalMoney)
	}

	if players[1].finalMoney != 20 {
		t.Errorf("handleDebt() failed, Player2 finalMoney should be 10, got: %d", players[1].finalMoney)
	}
}
