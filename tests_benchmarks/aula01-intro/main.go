package main

import "fmt"

func main () {

    result := soma(1,2,3,4)

    fmt.Println(result)
    
}

func soma (num ...int) int {
    total := 0

    for _, v := range num {
        total +=  v
    }

    return total
}
