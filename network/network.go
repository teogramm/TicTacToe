package network

import (
	"TicTacToe/persistence"
	"TicTacToe/utilities"
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isValidIP(ip string) bool{
	regex,_ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	return regex.MatchString(ip)
}

func getIP() string{
	reader := bufio.NewReader(os.Stdin)
	var input string
	for ;;{
		fmt.Printf("Enter an IP address: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if isValidIP(input){
			break
		}else{
			fmt.Printf("Wrong IP!\n")
		}
	}
	return input
}

func sendString(s string, writer *bufio.Writer){
	temp := []string{s,"\n"}
	writer.WriteString(strings.Join(temp,""))
}

func sendEnd(win,moves int){
	if win==1{
		sendString("-1",writer)
		utilities.DisplayWinner(win,moves)
	}else if win==2{
		sendString("-2",writer)
		utilities.DisplayWinner(win,moves)
	}
	if moves==9{
		sendString("-3",writer)
		utilities.DisplayWinner(win,moves)
	}
}

func handleConnection(conn net.Conn){
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)
	var win,moves int
	var str string
	win =0
	moves = 0
	level := utilities.InitializeBoard()
	for{
		str,_ = reader.ReadString('\n')
		persistence.ClearNewLine(&str)
		n,_ := strconv.Atoi(str)
		level[n/3][n%3] = "O"
		moves++
		win = utilities.CheckComplete(level)
		sendEnd(win,moves)
		utilities.ShowBoard(level)
		n = utilities.UserTurn(&level,"X")
		moves++
		sendEnd(win,moves)
		writer.WriteString(strconv.Itoa(n))
		writer.WriteString("\n")
		writer.Flush()
	}
	conn.Close()
}

func Server(){
	ln, err := net.Listen("tcp",":33988")
	if err!=nil{
		panic(err)
	}
	for{
		fmt.Printf("Waiting for connection...\n")
		conn, err := ln.Accept()
		if err!= nil{
			panic(err)
		}
		handleConnection(conn)
	}
}

func Client(){
	var str string
	var n int
	ip := getIP()
	s := []string{ip,":33988"}
	ip = strings.Join(s,"")
	conn, err := net.Dial("tcp",ip)
	if err!=nil{
		panic(err)
	}
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)
	level := utilities.InitializeBoard()
	for{
		utilities.ShowBoard(level)
		n = utilities.UserTurn(&level,"O")
		writer.WriteString(strconv.Itoa(n))
		writer.WriteString("\n")
		writer.Flush()
		str,_ = reader.ReadString('\n')
		persistence.ClearNewLine(&str)
		n,_ = strconv.Atoi(str)
		if n==-1 {
			fmt.Printf("You lost!")
		}
		if n==-2{
			fmt.Printf("You won!")
		}
		if n==-3{
			fmt.Printf("Draw!")
		}
		if n<0{
			conn.Close()
			os.Exit(0)
		}
		level[n/3][n%3] = "X"
	}
}
