/**
 * Created by IntelliJ IDEA.
 * User: felixtioh
 * Date: 11/10/13
 * Time: 3:34 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import (
	"encoding/json"
	"fmt"
)
type Maze struct {
	Dim int
	Treasures []*Treasure
	Players []*Player
}

func NewMaze(n, m  int) *Maze {
    maze := &Maze{
        Treasures: make([]*Treasure, 0),
        Players: make([]*Player, 0),
        Dim: n,
    }
    for i := 0; i < m; i++ {
        treasure := NewTreasure(i, n)
        maze.Treasures = append(maze.Treasures, treasure)
    }
    return maze
}

func (maze *Maze) AddPlayer(player *Player) *Maze{
	maze.Players = append(maze.Players, player)
	return maze
}

func (maze *Maze) ToJSON() string {
	jsonByteArray, _ := json.Marshal(maze)
	return string(jsonByteArray[:])
}

func (maze *Maze) Print() {
	fmt.Println("============> Printing Maze:")	
	printed := false
	for y:=0; y<maze.Dim; y++ {
		for x:=0; x<maze.Dim; x++ {
			printed = false
			for _, player := range maze.Players {
				if player.X == x && player.Y == y {
					fmt.Printf("[%d] ", player.Id)
					printed = true
					break	
				}
			}
			if printed {
				continue
			}
			for _, treasure := range maze.Treasures {
				if treasure.X == x && treasure.Y == y {
					fmt.Printf("<%d> ", treasure.Id)
					printed = true
					break
				}
			}
			if !printed {
				fmt.Printf("- ")
			}
			
		}
		fmt.Println("")
	}
	fmt.Println("========> End Printing Maze:")
}

func MazeFromJSON(str string) *Maze {
	maze := new(Maze)
	slice := []byte(str)
	json.Unmarshal(slice, &maze)
	return maze
}

