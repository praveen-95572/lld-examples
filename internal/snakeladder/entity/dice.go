package entity

import (
	"math/rand"
	"time"
)

type Dice struct {
	Sides int
}

func NewDice(sides int) *Dice {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Dice{Sides: sides}
}

func (d *Dice) Roll() int {
	return rand.Intn(d.Sides) + 1
}
