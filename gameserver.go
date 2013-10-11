package main
 
import (
    "net"
    "fmt"
	"time"
)

var playRoom *PlayRoom = NewPlayRoom()
var moves chan *Move = make(chan *Move)
var firstUserTrigger chan bool = make(chan bool)
var gameStarted bool = false


func ProcessMove(m *Move){
	
}

func GameEngine() int {
	for {
		select{
		case m := <- moves:
			ProcessMove(m) 
		case <- firstUserTrigger:
			fmt.Println("first user arrived")
			fmt.Println("starting count down 20 seconds")
			<- time.After(20 * time.Second)
			fmt.Println("count down ends. Game Started")
			gameStarted = true
			playRoom.StartGame()			
		}
		
	}
}

func main() {

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
