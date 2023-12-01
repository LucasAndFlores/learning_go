package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main () {
    contagem := 0

    numGoRoutines := 10

    var waitGroup sync.WaitGroup

    waitGroup.Add(10)

    for i := 0; i < numGoRoutines; i++ {
        go func () {
            tmp := contagem

            runtime.Gosched()

            tmp++

            contagem = tmp

            waitGroup.Done()
        }()

        fmt.Println("Total de go routines", numGoRoutines)

        fmt.Println("Total do contador", contagem)
    }

    waitGroup.Wait()
}
