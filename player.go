/**
 * Created by IntelliJ IDEA.
 * User: felixtioh
 * Date: 11/10/13
 * Time: 3:34 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"math/rand"
	"fmt"
)

type Player struct {
	Id int
	X int
	Y int
	TreasuresOwned int
}

func NewPlayer(id, dim int) *Player {
	fmt.Printf("Creating player with ID %d and dimension %d. \n", id, dim)
	var randX int
	var randY int
	uniqLoc := true
	for {
		uniqLoc = true
		randX = rand.Intn(dim)
		randY = rand.Intn(dim)
		for _, player := range maze.Players {
			fmt.Println(player.X)
			fmt.Println(randX)
			fmt.Println(player.Y)
			fmt.Println(randY)
			fmt.Printf("\n\n")
			if player.X == randX && player.Y == randY {
				fmt.Println("cannot occupied another players location, regenerating again.")
				continue
			}
		}
		if uniqLoc {
			break
		}
	}
    player := &Player{
        Id: id,
        TreasuresOwned: 0,
        X: randX,
        Y: randY,
    }
	fmt.Println(player)
	return player
}
