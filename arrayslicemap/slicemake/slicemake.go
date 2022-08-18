package main

import "fmt"

func main() {
	// make(tipo, quantidade de elementos)
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	// make(tipo, quantidade de elementos, quantidade de posições do array interno)
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))
}
