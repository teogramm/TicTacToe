package utilities

import (
	"testing"
)

func TestCheckIfOccupied(t *testing.T) {
	level := InitializeBoard()
	t.Log("Checking CheckIfOccupied")
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			if CheckIfOccupied(level,i,j){
				t.Error("Error when checking if occupied")
			}
		}
	}
	for i:=0;i<3;i++{
		tempsign := player1
		for j:=0;j<3;j++{
			level[i][j]=tempsign
			if !CheckIfOccupied(level,i,j){
				t.Error("Error when checking if occupied")
			}
			tempsign = player2
		}
	}
}

func TestShowBoard(t *testing.T) {
	level := InitializeBoard()
	ShowBoard(level)
	// Output:
	// 1 2 3
	// 4 5 6
	// 7 8 9
}

func TestDisplayWinner(t *testing.T) {
	DisplayWinner(0,9)
	// Output: Draw
	DisplayWinner(1,7)
	// Output: Player 1 won!
	DisplayWinner(2,7)
	// Output: Player 2 won!
}

func TestCheckComplete(t *testing.T) {
	t.Log("Checking game completion...")
	level := InitializeBoard()
	t.Log("Checking rows...")
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			level[i][j] = player1
		}
		if temp:=CheckComplete(level);temp!=1{
			t.Error("Error when checking if game is complete")
		}
	}
	t.Log("Checking columns")
	level = InitializeBoard()
	for j:=0;j<3;j++{
		for i:=0;i<3;i++{
			level[i][j] = player2
		}
		if temp:=CheckComplete(level);temp!=2{
			t.Error("Error when checking if game is complete")
		}
	}
	level = InitializeBoard()
	t.Log("Checking diagonal")
	for i:=0;i<3;i++{
		level[i][i]=player1
	}
	if temp:=CheckComplete(level);temp!=1{
		t.Error("Error when checking if game is complete")
	}
	level = InitializeBoard()
	for i:=0;i<3;i++{
		level[i][2-i]=player2
	}
	if temp:=CheckComplete(level);temp!=2{
		t.Error("Error when checking if game is complete")
	}
	level = InitializeBoard()
	if temp:=CheckComplete(level);temp!=0{
		t.Error("Error when checking if game is complete")
	}
}
