package main

import (
	"fmt"
	"github.com⁄Teogramm⁄TicTacToe/network"
	"github.com⁄Teogramm⁄TicTacToe/persistence"
	"github.com⁄Teogramm⁄TicTacToe/utilities"
	"math/rand"
	"strconv"
	"time"
)

const player1 = "O"
const player2 = "X"

func CheckForThree(level [3][3]string, player string) (i, j int) {
	//Ελέγχει εάν πρόκειται να σχηματιστεί τριάδα από κάποιο σύμβολο
	var opposite string
	var streak int
	if player == player1 {
		opposite = player2
	} else {
		opposite = player1
	}
	for i := 0; i < 3; i++ {
		streak = 0
		for j := 0; j < 3; j++ {
			if level[i][j] == player {
				streak++
			} else if level[i][j] == opposite {
				streak--
			}
		}
		if streak == 2 {
			for j := 0; j < 3; j++ {
				if level[i][j] != player {
					return i, j
				}
			}
		}
	}
	for j := 0; j < 3; j++ {
		streak = 0
		for i := 0; i < 3; i++ {
			if level[i][j] == player {
				streak++
			} else if level[i][j] == opposite {
				streak--
			}
		}
		if streak == 2 {
			for i := 0; i < 3; i++ {
				if level[i][j] != player {
					return i, j
				}
			}
		}
	}
	streak = 0
	for i := 0; i < 3; i++ {
		if level[i][i] == player {
			streak++
		} else if level[i][i] == opposite {
			streak--
		}
	}
	if streak == 2 {
		for j := 0; j < 3; j++ {
			if level[j][j] != player {
				return i, j
			}
		}
	}
	streak = 0
	for i := 0; i < 3; i++ {
		if level[i][2-i] == player {
			streak++
		} else if level[i][2-i] == opposite {
			streak--
		}
	}
	if streak == 2 {
		for j := 0; j < 3; j++ {
			if level[j][2-j] != player {
				return i, j
			}
		}
	}
	return -1, -1
}

func ai_turn(level *[3][3]string, moves int) {
	var i, j int
	if _, err := strconv.Atoi(level[1][1]); err == nil {
		level[1][1] = player2
		return
	} else {
		i, j = CheckForThree(*level, player2)
		if i != -1 {
			fmt.Println(i, j)
			if !utilities.CheckIfOccupied(*level, i, j) {
				level[i][j] = player2
				return
			}
		}
		i, j = CheckForThree(*level, player1)
		if i != -1 {
			if !utilities.CheckIfOccupied(*level, i, j) {
				level[i][j] = player2
				return
			}
		}
		if level[1][1] == player1 && moves == 1 {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			i := r.Intn(2)
			j := r.Intn(2)
			if i == 1 {
				i = 2
			}
			if j == 1 {
				j = 2
			}
			level[i][j] = player2
			return
		}
		if level[1][1] == player1 && moves == 3 {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			i := r.Intn(4)
			switch i {
			case 0:
				level[0][1] = player2
			case 1:
				level[1][0] = player2
			case 2:
				level[1][2] = player2
			default:
				level[2][1] = player2
			}
			return
		}
		for i = 0; i < 3; i++ {
			for j = 0; j < 3; j++ {
				if level[i][j] != player1 && level[i][j] != player2 {
					level[i][j] = player2
					return
				}
			}
		}
	}
}

func menu() int {
	var usersel int
	fmt.Printf("ΤΡΙΛΙΖΑ!!!\n1.Παιχνίδι με τον υπολογιστή\n2.Παινχίδι με 2 παίκτες\n3.Προβολή στατιστικών για " +
		"παίκτη\n4.Έξοδος\n5.Παιχνίδι δικτύου\nΕπιλέξτε: ")
	usersel = utilities.ReadInteger()
	for ; usersel < 1 || usersel > 5; {
		if usersel < 1 || usersel > 5 {
			fmt.Printf("Λάθος επιλογή!\n")
			fmt.Printf("ΤΡΙΛΙΖΑ!!!\n1.Παιχνίδι με τον υπολογιστή\n2.Παινχίδι με 2 παίκτες\n3.Προβολή στατιστικών για " +
				"παίκτη\n4.Έξοδος\nΕπιλέξτε: ")
			usersel = utilities.ReadInteger()
		}
	}
	return usersel
}

func main() {
	var sel int
	sel = 0
	var index2 int
	var players []persistence.Player
	for ; sel != 4; {
		players = persistence.LoadFile()
		sel = menu()
		if sel < 3 { //Αν ο χρήστης επιλέξει παιχνίδι αρχίζω το ταμπλό
			tempname := persistence.Getname()
			index1 := persistence.SearchName(tempname, &players)
			if index1 == -1 {
				players = append(players, persistence.NewPlayer(tempname))
				index1 = len(players) - 1
			}
			if sel == 2 {
				tempname = persistence.Getname()
				index2 = persistence.SearchName(tempname, &players)
				if index2 == -1 {
					players = append(players, persistence.NewPlayer(tempname))
					index2 = len(players) - 1
				}
			}
			level := utilities.InitializeBoard()
			moves := 0
			win := 0
			for ; moves < 9 && win == 0; {
				utilities.ShowBoard(level)
				_ = utilities.UserTurn(&level, player1)
				win = utilities.CheckComplete(level)
				moves++
				if win != 0 || moves > 8 {
					break
				}
				if sel == 2 {
					utilities.ShowBoard(level)
					_ = utilities.UserTurn(&level, player2)
				} else {
					ai_turn(&level, moves)
				}
				win = utilities.CheckComplete(level)
				moves++
				if win != 0 || moves > 8 {
					break
				}
			}
			utilities.ShowBoard(level)
			utilities.DisplayWinner(win, moves)
			if moves == 9 {
				players[index1].Draws++
				if sel == 2 {
					players[index2].Draws++
				}
			} else if win == 1 {
				players[index1].Wins++
				if sel == 2 {
					players[index2].Losses++
				}
			} else if win == 2 {
				players[index1].Losses++
				if sel == 2 {
					players[index2].Wins++
				}
			}
			persistence.SaveFile(&players)
		}
		if sel == 3 {
			temp := persistence.Getname()
			index := persistence.SearchName(temp, &players)
			if index == -1 {
				fmt.Println("Δεν υπάρχει ο παίκτης!")
			} else {
				players[index].PlayerStats()
			}
		}
		if sel == 5 {
			for {
				fmt.Printf("1.Server\n2.Client\nΕπιλέξτε: ")
				sel = utilities.ReadInteger()
				if sel < 1 || sel > 2 {
					continue
				} else {
					break
				}
			}
			if sel == 1 {
				network.ServerFlow()
			} else {
				network.ClientFlow()
			}
		}
	}
}
