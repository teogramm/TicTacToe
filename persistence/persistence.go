package persistence

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/teogramm/TicTacToe/utilities"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Player struct {
	Name                string
	Wins, Draws, Losses int
}

func (player *Player) NewPlayer(name string) {
	player.Name = name
	player.Wins = 0
	player.Losses = 0
	player.Draws = 0
}

func SearchName(name string, players *[]Player) int {
	for index, player := range *players {
		if player.Name == name {
			return index
		}
	}
	return -1
}

func Getname() string {
	reader := bufio.NewReader(os.Stdin)
	var buffer string
	fmt.Printf("Εισάγετε το όνομα σας: ")
	buffer,_ = reader.ReadString('\n')
	utilities.ClearNewLine(&buffer)
	return buffer
}

func fixfile(file *os.File) {
	file.Seek(0, 0)
	tempfile, err := os.OpenFile("scores.csv.tmp", os.O_CREATE|os.O_WRONLY, 0777)
	if err!=nil{
		log.Fatalln("Error opening file!")
	}
	defer tempfile.Close()
	writer := bufio.NewWriter(tempfile)
	defer writer.Flush()
	reader := bufio.NewReader(file)
	lineregex := regexp.MustCompile("(?P<name>[\\p{L}]+),(?P<wins>[\\d]+),(?P<draws>[\\d]+),(?P<losses>[\\d]+)")
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}else if err != nil {
			log.Fatalln("Error when reading score file!")
		}
		utilities.ClearNewLine(&line)
		if !lineregex.MatchString(line) {
			//Only bother if the line has a name
			if matched, _ := regexp.MatchString("^[\\p{L}]+,", line); matched {
				temp:=strings.Split(line,",")[0]
				_,_ = fmt.Fprintf(os.Stderr,"Resetting stats for user %s\n",temp)
				_,err = writer.WriteString(temp+",0,0,0\n")
				if err!=nil{
					log.Fatalln("Error writing to score file!")
				}
			}else{
				_,_ = fmt.Fprintf(os.Stderr,"User with no name deleted\n")
			}
		}else{
			_,err = writer.WriteString(line+"\n")
			if err!=nil{
				log.Fatalln("Error writing to score file!")
			}
		}
	}
	err = os.Remove(file.Name())
	if err!=nil{
		log.Fatalln("Error deleting old score file!")
	}
	err = os.Rename(tempfile.Name(),file.Name())
	if err!=nil{
		log.Fatalln("Error creating new score file!")
	}
}

func filerr(err error, file *os.File) {
	if err!=nil{
		_,_ = fmt.Fprintf(os.Stderr,"%s","Corrupted file. Attempting to fix...")
		fixfile(file)
	}
}

func LoadFile() []Player {
	var temp []Player
	file, err := os.OpenFile("scores.csv", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln("Error opening score file!")
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			filerr(err,file)
		}
		wins, err := strconv.Atoi(line[1])
		filerr(err, file)
		draws, err := strconv.Atoi(line[2])
		filerr(err, file)
		losses, err := strconv.Atoi(line[3])
		filerr(err, file)
		temp = append(temp, Player{Name: line[0], Wins: wins, Draws: draws, Losses: losses})
	}
	return temp
}

func SaveFile(players *[]Player) {
	scorefile, err := os.OpenFile("scores.csv", os.O_RDWR, 0666)
	defer scorefile.Close()
	if err != nil {
		log.Fatalln("Error opening score file!")
	}
	err = scorefile.Truncate(0)
	if err != nil {
		log.Fatalln("Error writing new score file!")
	}
	_, err = scorefile.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	scorecsv := csv.NewWriter(scorefile)
	defer scorecsv.Flush()
	for _, player := range *players {
		temp := []string{player.Name, strconv.Itoa(player.Wins), strconv.Itoa(player.Draws), strconv.Itoa(player.Losses)}
		err = scorecsv.Write(temp)
		if err != nil {
			panic(err)
		}
	}
}

func (player *Player) PlayerStats() {
	fmt.Printf("Player: %s\nWins: %d\n\tDraws: %d\n\t\tLosses: %d\n", player.Name, player.Wins, player.Draws,
		player.Losses)
}
