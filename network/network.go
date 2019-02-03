package network

import (
	"bufio"
	"fmt"
	"github.com⁄Teogramm⁄TicTacToe/persistence"
	"github.com⁄Teogramm⁄TicTacToe/utilities"
	"io"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	level  [3][3]string
	writer *bufio.Writer
	reader *bufio.Reader
	sign   string
	name   string
	first  bool
}

type Server struct {
	level  [3][3]string
	writer *bufio.Writer
	reader *bufio.Reader
	sign   string
	name   string
	first  bool
}

func getServerIP(){
	var ip net.IP
	//TODO: Not lazy check
	localip := regexp.MustCompile("192\\.168.*|10.*|172.[1-3].*]")
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if localip.MatchString(ip.String()) {
				fmt.Println("Your IP is:", ip)
			}
		}
	}
}

func (server *Server) Initialize() {
	server.level = utilities.InitializeBoard()
	server.name = persistence.Getname()
	ln, err := net.Listen("tcp", ":33988")
	if err != nil {
		panic(err)
	}
	getServerIP()
	fmt.Printf("Waiting for connection...\n")
	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	server.writer = bufio.NewWriter(conn)
	server.reader = bufio.NewReader(conn)
	server.Negotiate()
	server.tradeName()
	server.Game()
}

func (server *Server) tradeName(){
	var n string
	n,_ = server.reader.ReadString('\n')
	persistence.ClearNewLine(&n)
	fmt.Println("Παίζεις με τον παίκτη",n,"!")
	sendString(server.name,server.writer)
}

func (server *Server) Negotiate() {
	rand.Seed(time.Now().UTC().Unix())
	switch n := rand.Intn(1); n {
	case 0:
		server.sign = "X"
		fmt.Println("You are X. You play first.")
		server.first = true
		sendString("O", server.writer)
	default:
		server.sign = "O"
		fmt.Println("You are O. You play second.")
		sendString("X", server.writer)
		server.first = false
	}
}

func (server *Server) Game() {
	var win, n, moves int
	var str string
	win = 0
	moves = 0
	for {
		if server.first == true {
			utilities.ShowBoard(server.level)
			n = utilities.UserTurn(&server.level, server.sign)
			win = utilities.CheckComplete(server.level)
			moves++
			if win != 0 || moves == 9 {
				sendEnd(win, moves, server.writer)
				break
			}
			sendString(strconv.Itoa(n), server.writer)
		}
		server.first = true
		str, _ = server.reader.ReadString('\n')
		persistence.ClearNewLine(&str)
		n, _ := strconv.Atoi(str)
		server.level[n/3][n%3] = "O"
		moves++
		win = utilities.CheckComplete(server.level)
		if win != 0 || moves == 9 {
			sendEnd(win, moves, server.writer)
			break
		}
	}
}

func isValidIP(ip string) bool {
	regex, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	return regex.MatchString(ip)
}

func getIP() string {
	reader := bufio.NewReader(os.Stdin)
	var input string
	for ; ; {
		fmt.Printf("Enter an IP address: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if isValidIP(input) {
			break
		} else {
			fmt.Printf("Wrong IP!\n")
		}
	}
	return input
}

func sendString(s string, writer *bufio.Writer) {
	temp := []string{s, "\n"}
	_, _ = writer.WriteString(strings.Join(temp, ""))
	writer.Flush()
}

func sendEnd(win, moves int, writer *bufio.Writer) {
	if win == 1 {
		sendString("-1", writer)
		utilities.DisplayWinner(win, moves)
	} else if win == 2 {
		sendString("-2", writer)
		utilities.DisplayWinner(win, moves)
	}
	if moves == 9 {
		sendString("-3", writer)
		utilities.DisplayWinner(win, moves)
	}
}

func (client *Client) Connect() {
	var err error
	client.level = utilities.InitializeBoard()
	ip := getIP()
	client.name = persistence.Getname()
	s := []string{ip, ":33988"}
	ip = strings.Join(s, "")
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		panic(err)
	}
	client.writer = bufio.NewWriter(conn)
	client.reader = bufio.NewReader(conn)
	client.getSign()
	client.tradeName()
	client.game()
	conn.Close()
}

func (client *Client) tradeName(){
	var n string
	sendString(client.name,client.writer)
	n,_ = client.reader.ReadString('\n')
	persistence.ClearNewLine(&n)
	fmt.Println("Παίζεις με τον παίκτη",n,"!")
}

func (client *Client) getSign() {
	s, _ := client.reader.ReadString('\n')
	persistence.ClearNewLine(&s)
	client.sign = s
	if s == "O" {
		client.first = false
	} else {
		client.first = true
	}
	fmt.Println("You have", client.sign, "!")
}

func (client *Client) game() {
	var n int
	var str string
	var err error
	for {
		if client.first {
			utilities.ShowBoard(client.level)
			n = utilities.UserTurn(&client.level, "O")
			sendString(strconv.Itoa(n), client.writer)
		}
		client.first = true
		str, err = client.reader.ReadString('\n')
		if err==io.EOF{
			fmt.Println("Network error!")
			break
		}
		persistence.ClearNewLine(&str)
		n, _ = strconv.Atoi(str)
		if n == -1 {
			fmt.Printf("You lost!")
		}
		if n == -2 {
			fmt.Printf("You won!")
		}
		if n == -3 {
			fmt.Printf("Draw!")
		}
		if n < 0 {
			os.Exit(0)
		}
		client.level[n/3][n%3] = "X"
	}
}

func ServerFlow() {
	var server Server
	server.Initialize()
}

func ClientFlow() {
	var client Client
	client.Connect()
}
