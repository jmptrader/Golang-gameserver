package main

import (
    "fmt"
	"time"
)

func dirIsValid(dir string) bool {
	return dir == "S" || dir == "N" || dir == "E" || dir == "W"
}

func moveIsValid(move *Move) bool {
	switch {
	case move.player.Y == n-1 && move.dir == "S":
		return false
	case move.player.Y == 0 && move.dir == "N":
		return false
	case move.player.X == 0 && move.dir == "W":
		return false
	case move.player.X == n-1 && move.dir == "E":
		return false
	}
	
	return true
}

func movePlayerLoc(move *Move) {
	switch move.dir {
	case "S":
		move.player.Y++
	case "N":
		move.player.Y--
	case "W":
		move.player.X--
	case "E":
		move.player.X++
	}
}

func claimTreasure(move *Move){
	for _, treasure := range maze.Treasures {
		fmt.Println("move.player.X")
		fmt.Println(move.player.X)
		fmt.Println("move.player.Y")
		fmt.Println(move.player.Y)
		fmt.Println("treasure.X")
		fmt.Println(treasure.X)
		fmt.Println("treasure.Y")
		fmt.Println(treasure.Y)
		if treasure.OwnedBy == -1 {
			if treasure.X == move.player.X && treasure.Y == move.player.Y {
				treasure.OwnedBy = move.player.Id
				move.player.TreasuresOwned ++
				move.client.echoer <- "you found a treasure!"	
			} else {
				//move.client.echoer <- "treasure not found"
			}
		} else {
			//move.client.echoer <- "you moved into a treasure cell but it was owned by other player."
		}
	}
}

func ProcessMove(move *Move){
	fmt.Println(move)
	fmt.Println(move.player)
	fmt.Println(n)
	if moveIsValid(move) {
		//move.client.echoer <- "moving player location because it is valid"
		movePlayerLoc(move)
		claimTreasure(move)
		move.client.outgoing <- maze.ToJSON()
	} else {
		//move.client.echoer <- "unable to move player location because it is invalid. Please try another move."		
	}
}

func GameEngine() int {
	for {
		select{
		case m := <- moves:
			fmt.Println("Start processing a move. Other move should not come in")
			ProcessMove(m) 
			fmt.Println("Processed a move. Other moves can be processed now")
		case <- firstUserTrigger:
			fmt.Println("first user arrived")
			fmt.Println("starting count down 20 seconds")
			<- time.After(20 * time.Second)
			fmt.Println("count down ends. Game Started")
			gameStarted = true
			playRoom.BroadcastInitialState()			
		}
		
	}
}