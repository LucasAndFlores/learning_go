package main

import (
	"fmt"
	"log"
)

func main () {
    result, err := double(-10)
    if (err != nil) {
        log.Panicln("An error happened", err)
    }
    fmt.Println(result)
}

func double (n float64) (float64, error) {
    if n < 0 {
        return 0,fmt.Errorf("Can't execute the function with this value: %v", n)

    }
    return n * 2, nil
}
