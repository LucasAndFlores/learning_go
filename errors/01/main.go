package main

import "fmt"

func main () {
    var answer1 string

    fmt.Print("Name: ")
    _, err := fmt.Scan(&answer1)
    if err != nil {
        panic(err)
    }

    fmt.Println(answer1)
}
