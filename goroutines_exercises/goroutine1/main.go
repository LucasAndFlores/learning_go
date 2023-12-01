package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main () {

    waitGroup.Add(5)
go routine1()
go routine2()
//    for i := 0; i < 2;i++ {
//        go func () {
//            fmt.Println("Dentro do loop", i) 
//            waitGroup.Done()
//        }()
//    }

    waitGroup.Wait()

    
}

func routine1 () {
    fmt.Println("Go routine 1")
    waitGroup.Done()
}


func routine2 () {
    fmt.Println("Go routine 2")
    waitGroup.Done()
}
