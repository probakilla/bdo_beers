package main

import "fmt"

const (
    PATATE int = 5
    LEAVING int = 2
    SUGAR int = 1
    WATER int = 6
)

type Beer struct {
    count int
    patate int
    leaving int
    sugar int
    water int
}

func NewBeer(nbPatates int) Beer {
    var b Beer
    b.count = nbPatates / PATATE
    b.patate = b.count * PATATE
    b.leaving = b.count * LEAVING
    b.sugar = b.count * SUGAR
    b.water = b.count * WATER
    return b
}

func (b Beer) toString() string {
    var str string = fmt.Sprintf("To make %d beers, you need : \n" +
    "Patates : %d\nLeaving : %d\nSugar : %d\nWater : %d\n",
    b.count, b.patate, b.leaving, b.sugar, b.water)
    return str
}

