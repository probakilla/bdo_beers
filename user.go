package main

import (
    "fmt"
    "strconv"
)

func askUserForPatate() int {
    var correct bool = false
    var nbPatates int
    for ! correct {
        fmt.Println("Please enter the number of patates you have")
        var patates string
        fmt.Scanln(&patates)
        var err error
        nbPatates, err = strconv.Atoi(patates)
        if err == nil {
            correct = ! correct
        }
    }
    return nbPatates
}

