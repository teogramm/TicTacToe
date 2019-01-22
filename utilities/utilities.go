package utilities

import(
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"os"
)

const player1 = "O"
const player2 = "X"

func ReadInteger() int{
	reader := bufio.NewReader(os.Stdin)
	var i int
	var err error
	for ;true;
	{
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		i, err = strconv.Atoi(input)
		if err!=nil{
			fmt.Println("Λάθος είσοδος!")
		}else{
			break
		}
	}
	return i
}

func CheckIfOccupied(level [3][3]string, i,j int) bool {
	if _, err := strconv.Atoi(level[i][j]); err == nil {
		return false
	} else {
		return true
	}
}

func UserTurn(level *[3][3]string, sign string) int{
	var usersel int
	fmt.Printf("Επιλέξτε θέση: ")
	usersel = ReadInteger()
	ok := false
	for ; ok == false; {
		usersel-=1
		ok = usersel +1 > 0 && usersel+1 < 10 && !CheckIfOccupied(*level, usersel/3,usersel%3)
		if ok == false {
			fmt.Printf("Λάθος επιλογή!\n")
			fmt.Printf("Επιλέξτε θέση: ")
			usersel = ReadInteger()
		}
	}
	level[usersel/3][usersel%3] = sign
	return usersel
}


func CheckComplete(lev [3][3]string) int{
	//Επιστρέφω 0 αν δεν έχει σηματιστεί τριάδα, 1 αν νίκησε ο player1, 2 αν νίκησε ο player2
	for i:=0;i<3;i++{
		if lev[i][0] == lev[i][1] && lev[i][1] == lev[i][2] {
			if lev[i][0] == player1 {
				return 1
			} else if lev[i][0] == player2 {
				return 2
			}
		}
		if lev[0][i] == lev[1][i] && lev[1][i] == lev[2][i] {
			if lev[0][i] == player1{
				return 1;
			}else if lev[0][i] == player2{
				return 2
			}
		}
	}
	//Κυρια διαγώνιος
	if lev[0][0] == lev[1][1] && lev[1][1] == lev[2][2] {
		if lev[0][0] == player1{
			return 1
		} else if lev[0][0] == player2{
			return 2
		}
	}
	//Ελέγχω δευτερεύουσα διαγώνιο
	if lev[0][2] == lev[1][1] && lev[1][1] == lev[2][0] {
		if lev[0][2] == player1{
			return 1
		} else if lev[0][2] == player2{
			return 2
		}
	}
	return 0
}

func ShowBoard(level [3][3]string){
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			fmt.Printf("%s ",level[i][j])
		}
		fmt.Printf("\n")
	}
}

func DisplayWinner(win,moves int){
	if moves == 9{
		fmt.Println("Ισοπαλία")
	} else if win == 1{
		fmt.Println("Νίκησε ο παίκτης 1!")
	} else if win ==2{
		fmt.Println("Νίκησε ο παίκτης 2!")
	}
}

func InitializeBoard()[3][3]string{
	count := 1
	var level [3][3]string
	for i:=0;i<3;i++ {
		for j:=0;j<3;j++{
			level[i][j] = strconv.Itoa(count)
			count++
		}
	}
	return level
}