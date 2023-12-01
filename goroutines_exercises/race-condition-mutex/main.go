package main

import (
	"fmt"
	"runtime"
	"sync"
)


func main () {
    contagem := 0

    numGoRoutines := 100

    var waitGroup sync.WaitGroup
    var mu sync.Mutex


    waitGroup.Add(numGoRoutines)
    for i := 0; i < numGoRoutines; i++ {
        go func () {
            mu.Lock()
            tmp := contagem


            runtime.Gosched()

            tmp++


            contagem = tmp
            
            mu.Unlock()
            waitGroup.Done()
        }()

    }
    waitGroup.Wait()
    fmt.Println("Total de go routines", numGoRoutines)

    fmt.Println("Total do contador", contagem)

}
