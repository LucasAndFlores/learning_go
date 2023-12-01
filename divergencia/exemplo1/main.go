package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main () {
    canal1 := make(chan int)
    canal2 := make(chan int)

    go lansaABraba(20, canal1)
    go outra(canal1,canal2)

    for v := range canal2 {
        fmt.Println(v)
    }
}

func lansaABraba(n int, canal chan int) {
    for i := 0; i < n; i++ {
        canal <- i
    }

    close(canal)
}

func outra(canal1, canal2 chan int) {
    var wg sync.WaitGroup


    for v := range canal1 {
    wg.Add(1)
        go func (n int) {
            canal2 <- trabalho(n)
            wg.Done()
        }(v)
    }

    wg.Wait()

    close(canal2)
}

func trabalho(n int) int {
    time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))

    return n
}




//- Dois canais.
//- Uma func manda X números ao primeiro canal.
//- Outra func faz um range deste canal, e para cada ítem lança uma go func que poe o retorno de trabalho() no canal dois.
//- Trabalho() é um timer aleatório pra simular workload.
