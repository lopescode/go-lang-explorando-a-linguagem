package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva":   12412.33,
			"Guilherme Santos": 1200.0,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 2325.45,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
