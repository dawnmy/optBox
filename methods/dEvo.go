package methods

import (
	"math/rand"
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
	n := SetParameter.NumV
	p := SetParameter.NumI
	xl := SetParameter.XLower
	xu := SetParameter.XUpper
	xall := make(float64Mat, p, n)

	for _, i := range rand.Perm(p) {
		xall[i] = xl + rand.Float32()*(xu-xl)
	}
	return (xall)

	// for _, i := range rand.Perm(p) {
	// 	fmt.Println(i)
	// }
}

func Mutation() {

}

func Crossover() {

}

func Selection() {

}

func ObjFuc() {
}

func DiffEvo() string {

	return "Test"

	//_ = xUpper
	//_ =
}
