package utilities

import "testing"

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
}

