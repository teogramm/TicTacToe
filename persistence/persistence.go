package persistence

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/teogramm/TicTacToe/utilities"
	"io"
	"os"
	"strconv"
)

type Player struct {
	Name                string
	Wins, Draws, Losses int
}


func (player *Player) NewPlayer(name string){
	player.Name = name
	player.Wins = 0
	player.Losses = 0
	player.Draws = 0
}

func SearchName(name string,players *[]Player)int{
	for index, player := range *players{
		if player.Name == name{
			return index
		}
	}
	return -1
}

func Getname() string{
	reader := bufio.NewReader(os.Stdin)
	var buffer string
	fmt.Printf("Εισάγετε το όνομα σας: ")
	buffer,_ = reader.ReadString('\n')
	utilities.ClearNewLine(&buffer)
	return buffer
}

func fixfile(file *os.File){
	file.Seek(0,0)

}

func filerr(err error,file *os.File){
	if err!=nil{
		fmt.Println("Malformed file. Want to try and repair it?(Data may be lost!)")
		if utilities.PromptYesNo(){
			fixfile(file)
		}else {
			os.Exit(5)
		}
	}
}

func LoadFile() []Player{
	var temp []Player
	file, err := os.OpenFile("scores.csv",os.O_RDONLY|os.O_CREATE,0666)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for{
		line, err := reader.Read()
		if err == io.EOF{
			break
		} else if err != nil{
			panic(err)
		}
		wins,err := strconv.Atoi(line[1])
		filerr(err,file)
		draws,err := strconv.Atoi(line[2])
		filerr(err,file)
		losses,_ := strconv.Atoi(line[3])
		filerr(err,file)
		temp = append(temp,Player{Name: line[0], Wins:wins, Draws:draws, Losses:losses})
	}
	return temp
}

func SaveFile(players *[]Player){
	scorefile,err := os.OpenFile("scores.csv",os.O_RDWR, 0666)
	defer scorefile.Close()
	if err!=nil{
		panic(err)
	}
	err = scorefile.Truncate(0)
	if err!= nil{
		panic(err)
	}
	_,err = scorefile.Seek(0,0)
	if err!= nil{
		panic(err)
	}
	scorecsv := csv.NewWriter(scorefile)
	defer scorecsv.Flush()
	for _,player := range *players{
		temp := []string{player.Name,strconv.Itoa(player.Wins),strconv.Itoa(player.Draws),strconv.Itoa(player.Losses)}
		err = scorecsv.Write(temp)
		if err != nil{
			panic(err)
		}
	}
}

func (player *Player) PlayerStats(){
	fmt.Printf("Player: %s\nWins: %d\n\tDraws: %d\n\t\tLosses: %d\n",player.Name,player.Wins,player.Draws,
		player.Losses)
}