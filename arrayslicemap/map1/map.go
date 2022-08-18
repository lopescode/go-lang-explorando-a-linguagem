package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[1234567898] = "Maria"
	aprovados[9876543210] = "Pedro"
	aprovados[9124795400] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[9876543210])
	delete(aprovados, 9876543210)
	fmt.Println(aprovados[9876543210])

	// acessar map pela mesma chave reatribui o valor
	aprovados[1234567898] = "Nome modificado"
	fmt.Println(aprovados)
}
