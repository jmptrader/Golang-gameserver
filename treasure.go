/**
 * Created by IntelliJ IDEA.
 * User: felixtioh
 * Date: 11/10/13
 * Time: 3:34 PM
 * To change this template use File | Settings | File Templates.
 */
package main

import "math/rand"

type Treasure struct {
	Id int
	OwnedBy int // -1 means not owned by anyone
	X int
	Y int
}

func NewTreasure(id, dim int) *Treasure {
    return &Treasure{
        Id: id,
        OwnedBy: -1,
        X: rand.Intn(dim),
        Y: rand.Intn(dim),
    }
}
