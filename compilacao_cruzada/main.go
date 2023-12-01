package main

import (
    "fmt"
    "runtime"
)

// GOOS=macos GOARCH=amd64 go build main.go 
// GOOS=windows GOARCH=amd64 go build main.go
 

func main () {
        fmt.Println("Essa é uma compilação cruzada. Foi compilada num linux/amd64 e agora esta rodando num sistema: ", runtime.GOARCH, runtime.GOOS)
}
