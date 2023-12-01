package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main () {
    var contagem atomic.Int32

    numGoRoutines := 10000

    var waitGroup sync.WaitGroup

    waitGroup.Add(numGoRoutines)

    for i := 0; i < numGoRoutines; i++ {
        go func () {
        contagem.Add(1)

            runtime.Gosched()
            contagem.Load()
            waitGroup.Done()
        }()

    }

    waitGroup.Wait()

    fmt.Println("Total de go routines", numGoRoutines)

    fmt.Println("Total do contador", contagem.Load())
}
