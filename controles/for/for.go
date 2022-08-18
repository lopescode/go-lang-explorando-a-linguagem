package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	// estrutura while
	for i <= 10 { // não tem while em Go
		fmt.Println(i)
		i++
	}

	// for convencional
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	// for convencional com estrutura aninhada
	fmt.Println("\nMisturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	// laço for infinito
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second * 5)
	}
}
