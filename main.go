package main

import (
	"fmt"
	"runtime"

	"./methods"
)

func main() {
	runtime.GOMAXPROCS(2)
	xl := []float64{0, 1, 0, 2}
	xu := []float64{5, 6, 8, 4}
	par := methods.Parameter{20, 100, 0.5, 0.7, xl, xu}
	bestCost, bestVariables := methods.DiffEvo(par)
	fmt.Println("Final genration\nbest cost: ",
		bestCost,
		"\nbest variables: ",
		bestVariables)
}
