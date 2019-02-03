package persistence

import "testing"

func TestPlayer_NewPlayer(t *testing.T) {
	var player Player
	player.NewPlayer("Test")
	if player.Name != "Test"{
		t.Error("Error setting name")
	}
	if player.Wins!=0||player.Draws!=0||player.Losses!=0{
		t.Error("Error setting default values for struct player")
	}
}

func TestPlayer_PlayerStats(t *testing.T) {
	var player Player
	player.NewPlayer("Test")
	player.Wins = 5
	player.Losses = 6
	player.Draws = 1
	player.PlayerStats()
	// Output:
	// Player: Test
	// Wins: 5
	// Draws: 1
	// Losses: 6
}