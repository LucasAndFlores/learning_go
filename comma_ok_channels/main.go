// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go mandaNumeros(par, impar, quit)
	receive(par, impar, quit)
}

func mandaNumeros(par, impar chan int, quit chan bool) {
	total := 100

	for i := 0; i <= total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}

        quit <- true
	close(par)
	close(impar)
}

func receive(par, impar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("O numero ", v, "par")
		case v := <-impar:
			fmt.Println("O numero ", v, "impar")
                case v, ok := <-quit:
                        if !ok {
                        fmt.Println("Erro!", v)
                        } else {
                        fmt.Println("Encerrando", v)
                        return
                        }
		}
	}
}


