package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

type float64Data []float64
type float64Mat [][]float64

// Set the parameters
type Parameter struct {
	NumV   int         // Number of variables to be optimized
	NumI   int         // Population size (number of individuals for each generation)
	NumG   int         // Number of generations
	CR     float64     // Crossover rate
	F      float64     // Scalling factor
	XUpper float64Data // Upper bounds of variables
	XLower float64Data // Lower bounds of variables
}

// Hold the population for each generation
type DE struct {
	Parameter
	Population float64Mat //
	H          float64Data
	U          float64Data
	M          float64Data
}

func InitPop(SetParameter Parameter) float64Mat {
	//n := SetParameter.NumV
	p := SetParameter.NumI
	xl := SetParameter.XLower
	xu := SetParameter.XUpper
	xall := float64Mat{}
	tmp := float64Data{}
	for i := 0; i < p; i++ {
		tmp = float64Data{}
		for j, _ := range xl {
			tmp = append(tmp, xl[j]+rand.Float64()*(xu[j]-xl[j]))
		}
		xall = append(xall, tmp)
	}
	return xall
}

func Mutation() {

}

func Crossover() {

}

func Selection() {

}

func ObjFuc() {
}

func objFunc(x []float64) float64 {
	return 2*x[0]*x[0] - x[1]*x[1] + 3
}

func DiffEvo() string {

	return "Test"

	//_ = xUpper
	//_ =
}

func main() {
	runtime.GOMAXPROCS(2)
	//fmt.Println(rand.Perm(10))
	//methods.Rand(10)
	//methods.SetDefault()
	// pop1 := methods.DE{
	// 	Parameters{
	// 		0, 0, 15, 0,
	// 		0.9, 0.9, -1},
	// 	XUpper,
	// 	XLower,
	// 	Population,
	// 	FVals,
	// 	U}
	//newX := append(methods.XUpper, 1, 2, 3, 4)
	// fmt.Println(append(newX, 5, 7))
	//fmt.Println(Cr)()
	const n = 2
	xl := []float64{0, 1}
	xu := []float64{10, 10}
	//xall := [][]float64{}
	// tmp := []float64{}
	par := Parameter{n, 10, 100, 0.5, 0.5, xl, xu}
	xall := InitPop(par)
	fmt.Println(xall)
}
