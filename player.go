/**
 * Created by IntelliJ IDEA.
 * User: felixtioh
 * Date: 11/10/13
 * Time: 3:34 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import "math/rand"


type Player struct {
	Id int
	X int
	Y int
	TreasuresOwned int
}

func NewPlayer(id, dim int) *Player {
    return &Player{
        Id: id,
        TreasuresOwned: 0,
        X: rand.Intn(dim),
        Y: rand.Intn(dim),
    }
}
