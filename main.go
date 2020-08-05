package main

import "fmt"

func main() {
    var patates int = askUserForPatate()
    beer := NewBeer(patates)
    fmt.Println(beer.toString())
}

