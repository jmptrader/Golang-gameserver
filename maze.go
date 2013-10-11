/**
 * Created by IntelliJ IDEA.
 * User: felixtioh
 * Date: 11/10/13
 * Time: 3:34 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import "encoding/json"

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

func MazeFromJSON(str string) *Maze {
	maze := new(Maze)
	slice := []byte(str)
	json.Unmarshal(slice, &maze)
	return maze
}

