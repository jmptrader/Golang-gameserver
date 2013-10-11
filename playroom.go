package main
 
import (
    "net"
	"fmt"
	"strings"
)
 
type PlayRoom struct {
	clients []*Client
	joins chan net.Conn
	incoming chan string
	outgoing chan string
}

func (playRoom *PlayRoom) StartGame() {
	for _, client := range playRoom.clients {
		client.outgoing <- "hi all!!"
	}
}

func dirIsValid(dir string) bool {
	return dir == "S" || dir == "N" || dir == "E" || dir == "W"
}

/*
	When new connection is established
*/
func (playRoom *PlayRoom) Join(connection net.Conn) {
	client := NewClient(connection)
	msg := ""
	playRoom.clients = append(playRoom.clients, client)
	if len(playRoom.clients) == 1 {
		firstUserTrigger <- true
	}
	go func() { 
		for { 
			dir := <-client.incoming
			dir = strings.TrimSpace(dir)
			if gameStarted {
				if dirIsValid(dir){
					move := &Move{
						client: client,
						dir: dir,
					} 
					moves <- move		
				} else {
					msg = "Valid moves are: N, S, E, W"
					fmt.Println(msg)
					client.outgoing <- msg	
				}	
			} else {
				msg = "Do not move yet. Game has yet to start"
				fmt.Println(msg)
				client.outgoing <- msg
			}		
		} 
	}()
}

func (playRoom *PlayRoom) Listen() {
	go func() {
		for {
			select {
			case conn := <-playRoom.joins:
				playRoom.Join(conn)
			}
		}
	}()
}

func NewPlayRoom() *PlayRoom {
	playRoom := &PlayRoom{
		clients: make([]*Client, 0),
		joins: make(chan net.Conn),
		incoming: make(chan string),
		outgoing: make(chan string),
	}

	playRoom.Listen()

	return playRoom
}