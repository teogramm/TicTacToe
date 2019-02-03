package main

import (
	"testing"
)
import "github.com/teogramm/TicTacToe/utilities"


func TestCheckForThree(t *testing.T){
	var level [3][3]string
	t.Log("Testing rows...")
	for i:=0;i<3;i++{
		level = utilities.InitializeBoard()
		for j:=0;j<2;j++{
			level[i][j] = player2
		}
		if resi,resj:=CheckForThree(level,player2);resi!=i || resj!=2{
			t.Error("Checking for three does not work!")
		}
	}
	t.Log("Testing columns...")
	for j:=0;j<3;j++{
		level = utilities.InitializeBoard()
		for i:=0;i<2;i++{
			level[i][j] = player2
		}
		if resi,resj:=CheckForThree(level,player2);resi!=2 || resj!=j{
			t.Error("Checking for three does not work!")
		}
	}
	t.Log("Testing diagonals...")
	level = utilities.InitializeBoard()
	for i:=0;i<2;i++{
		level[i][i] = player2
	}
	if resi,resj:=CheckForThree(level,player2);resi!=2 || resj!=2{
		t.Error("Checking for three does not work!")
	}
	level = utilities.InitializeBoard()
	for i:=0;i<2;i++{
		level[i][2-i] = player2
	}
	if resi,resj:=CheckForThree(level,player2);resi!=0 || resj!=2{
		t.Error("Checking for three does not work!")
	}
}

func TestAITurn(t *testing.T){
	level := utilities.InitializeBoard()
	level[1][1]=player1
	ai_turn(&level,1)
	t.Log("Testing AI...")
	if !(level[0][0]==player2||level[0][2]==player2||level[2][2]==player2 ||level[2][0]==player2){
		t.Error("AI Error: Not playing correctly when player places in center")
	}
	level = utilities.InitializeBoard()
	ai_turn(&level,1)
	if level[1][1]!=player2{
		t.Error("AI Error: Not putting it in the center when it is free!")
	}
	level = utilities.InitializeBoard()
	level[0][0] = player1
	level[1][1] = player2
	level[1][2] = player1
	ai_turn(&level,3)
	if !(level[1][0]==player2 || level[1][2]==player2||level[0][1]==player2||level[2][1]==player2){
		t.Error("AI Error: Not blocking properly (case 4)")
	}
	level = utilities.InitializeBoard()
	level[1][1] = player2
	level[0][1] = player2
	ai_turn(&level,5)
	if level[2][1] != player2{
		t.Error("AI Error: Not attacking properly")
	}
	level = utilities.InitializeBoard()
	level[1][1] = player1
	level[0][1] = player1
	ai_turn(&level,3)
	if level[2][1] != player2{
		t.Error("AI Error: Not blocking properly (simple case)")
	}
}
