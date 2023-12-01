package main

import ("fmt"
"example.com/greetings"
)

func main () {
    fmt.Println("This is inside otherpackage")
    greetings.SayHi()
}


