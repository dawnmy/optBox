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
	par := methods.Parameter{20, 200, 0.5, 0.7, xl, xu}
	fmt.Println("Final genration: ", methods.DiffEvo(par))
}
