package methods

import (
	"fmt"
	"math/rand"
	//"runtime"
	//"math"
	//"github.com/spf13/cobra"
)

type float64Data []float64

// Set the parameters
type Parameter struct {
	NumV   int     // Number of variables to be optimized
	Pop    int     // Population size
	MaxGen int     // Number of generations
	CR     float64 // Crossover rate
	F      float64 // Scalling factor
	//FBest float64 //
	XUpper float64Data // Upper bounds of variables
	XLower float64Data // Lower bounds of variables

}

// Hold the population for each generation
type DE struct {
	Parameter
	Population float64Data //
	FVals      float64Data
	U          float64Data
}

func InitPop(XUpper, XLower float64) float64Data{
	

}

func Mutation(){

}

func CrossOver(){

}

func Selection() {

}

func Rand(Pop int) {
	//R = list(np.random.permutation(P))
	for _, i := range rand.Perm(Pop) {
		fmt.Println(i)

	}
	//     (j, k, l, u, v) = (R[i] for i in range(5))
	//     if j == p:
	//         j = R[6]
	//     elif k == p:
	//         k = R[6]
	//     elif l == p:
	//         l = R[6]
	//     elif u == p:
	//         u = R[6]
	//     elif v == p:
	//         v = R[6]
	//     M = X[j,:] + F*(X[k,:]-X[l,:])
	// return M
}

// func Use(vals ...interface{}) {
// 	for _, val := range vals {
// 		_ = val
// 	}
// }

// //func SetDefault(){
// 	var Dim = 0
// 	//xUpper := make([] int, 100)
// 	var XUpper = FloatArr{}
// 	//xLower := make([] int, 100)
// 	var XLower = FloatArr{}
// 	var Population = FloatArr{}
// 	var FVals = FloatArr{}
// 	var NumFE = 0 //count the total number of function evaluations
// 	var MaxGen = 0 //number of generations
// 	var NumPop = 15
// 	var Cr = 0.90 //crossover probability
// 	var F = 0.90 //Scaling factor
// 	var U = FloatArr{} //trial vector
// 	var FBest = -1
//Use(Dim, XUpper, XLower, Population, FVals, NumFE, MaxGen, NumPop, Cr, F, U, FBest)
//_ = dim,
//_ = xUpper
//_ = xLower, population, fVals, numFE, maxGen, numPop, cr, F, U, fBest
//}

// func makeDE(de DE){
// 	return

// }

func DiffEvo() string {

	return "Test"

	//_ = xUpper
	//_ =

}
