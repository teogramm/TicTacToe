package main

import "testing"
import "github.com⁄Teogramm⁄TicTacToe/utilities"


func TestThree(t *testing.T){
	var level [3][3]string
	t.Log("Testing rows...")
	for i:=0;i<3;i++{
		level = utilities.InitializeBoard()
		for j:=0;j<2;j++{
			level[i][j] = "X"
		}
		if resi,resj:=CheckForThree(level,"X");resi!=i || resj!=2{
			t.Error("Checking for three does not work!")
		}
	}
	t.Log("Testing columns...")
	for j:=0;j<3;j++{
		level = utilities.InitializeBoard()
		for i:=0;i<2;i++{
			level[i][j] = "X"
		}
		if resi,resj:=CheckForThree(level,"X");resi!=2 || resj!=j{
			t.Error("Checking for three does not work!")
		}
	}
}

func TestAI(t *testing.T){
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
}