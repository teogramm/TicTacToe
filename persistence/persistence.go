package persistence

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	Name                string
	Wins, Draws, Losses int
}

func ClearNewLine(temp *string){
	*temp = strings.TrimRight(*temp,"\r\n")
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
	ClearNewLine(&buffer)
	return buffer
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
		wins,_ := strconv.Atoi(line[1])
		draws,_ := strconv.Atoi(line[2])
		losses,_ := strconv.Atoi(line[3])
		temp = append(temp,Player{Name: line[0], Wins:wins, Draws:draws, Losses:losses})
	}
	return temp
}

func NewPlayer(name string) Player{
	var newp Player
	newp.Draws =0
	newp.Losses = 0
	newp.Wins = 0
	newp.Name = name
	return newp
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