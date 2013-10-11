package main
 
import (
    "net"
    "fmt"
	"strconv"
	"strings"
	"os"
	"bufio"
)

var playRoom *PlayRoom = NewPlayRoom()
var moves chan *Move = make(chan *Move)
var firstUserTrigger chan bool = make(chan bool)
var gameStarted bool = false
var n int
var m int
var maze *Maze

func main() {
	
    // get game config
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter N")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    n, _ = strconv.Atoi(input)
    fmt.Println("Enter your M")
    input, _ = reader.ReadString('\n')
    input = strings.TrimSpace(input)
    m, _ = strconv.Atoi(input)
    fmt.Printf("N: %d, M: %d \n", n, m)

    // build maze
    maze = NewMaze(n, m)

	listener, _ := net.Listen("tcp", ":6666")

	go GameEngine()
	
	for {
		conn, _ := listener.Accept()
		fmt.Println("new connection established.")
		if gameStarted {
			fmt.Println("no longer accepting new clients")
			conn.Close()
		} else {
			playRoom.joins <- conn
		}
	}
}
