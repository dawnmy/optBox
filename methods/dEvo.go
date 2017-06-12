package methods

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type float64Data []float64
type float64Mat [][]float64

// Parameter holds the parameters for population, individual and DE
type Parameter struct {
	NumI   int         // Population size (number of individuals for each generation)
	NumG   int         // Number of generations
	CR     float64     // Crossover rate
	F      float64     // Scalling factor
	XUpper float64Data // Upper bounds of variables
	XLower float64Data // Lower bounds of variables

}

// // Hold the population for each generation
// type DE struct {
// 	Parameter
// 	Population float64Mat //
// 	H          float64Data
// 	U          float64Data
// 	M          float64Data
// }

// Initialize the population
func initPop(setParameter Parameter) float64Mat {
	ps := setParameter.NumI
	xl := setParameter.XLower
	xu := setParameter.XUpper
	xall := float64Mat{}
	tmp := float64Data{}
	for i := 0; i < ps; i++ {
		tmp = float64Data{}
		for j := range xl {
			// Randomization seed
			rand.Seed(time.Now().UTC().UnixNano())
			// Generation random individuals in the range of xl and xu
			tmp = append(tmp, xl[j]+rand.Float64()*(xu[j]-xl[j]))
		}
		xall = append(xall, tmp)
	}
	return xall
}

// Mutation for evolution
func mutation(gGeneration float64Mat,
	individual int,
	f float64,
	xLower float64Data,
	xUpper float64Data) float64Data {
	// depletion of the the old current individual
	gWithOutI := make(float64Mat, individual)
	copy(gWithOutI, gGeneration[:individual])
	gWithOutI = append(gWithOutI, gGeneration[individual+1:]...)

	// Shuffle the population
	shuffle(&gWithOutI)

	var tmp float64
	mutatedIndividual := float64Data{}

	for key, value := range gWithOutI[0] {

		// Mutation based on DE formula
		tmp = value + f*(gWithOutI[1][key]-gWithOutI[2][key])

		// Check the bound for variables
		if tmp > xUpper[key] {
			tmp = xUpper[key]
		} else if tmp < xLower[key] {
			tmp = xLower[key]
		} else {

		}

		mutatedIndividual = append(mutatedIndividual, tmp)
	}
	return mutatedIndividual
}

// Recombination of chromosomes of individual
func crossover(individual *float64Data,
	cr float64,
	mutatedIndividual float64Data) float64 {
	n := len(*individual)
	vi := make(float64Data, n)
	for i, j := range mutatedIndividual {
		rand.Seed(time.Now().UTC().UnixNano())
		if v := rand.Float64(); v > cr {
			vi[i] = (*individual)[i]
		} else {
			vi[i] = j
		}
	}

	old := objFunc(*individual)
	new := objFunc(vi)
	if old > new {
		*individual = vi
		return new
	}
	return old
}

// Alternative recombination function with return
func crossover2(individual float64Data,
	cr float64,
	mutatedIndividual float64Data) float64Data {
	n := len(individual)
	vi := make(float64Data, n)
	for i, j := range individual {
		if v := rand.Float64(); v > cr {
			vi[i] = j
		} else {
			vi[i] = mutatedIndividual[i]
		}
	}

	// Selection based on cost function
	if objFunc(individual) > objFunc(vi) {
		individual = vi
	} else {

	}
	return individual
}

// The shuffle function
func shuffle(population *float64Mat) {
	n := len(*population)
	// Set the randomization seed
	rand.Seed(time.Now().UTC().UnixNano())
	for i, j := range rand.Perm(n) {
		(*population)[i], (*population)[j] = (*population)[j], (*population)[i]
	}
}

func shuffle2(population *float64Mat) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < len(*population); i++ {
		j := rand.Intn(i + 1)
		(*population)[i], (*population)[j] = (*population)[j], (*population)[i]
	}
}

// The objective cost funtion for optimization
func objFunc(x float64Data) float64 {
	return 5*math.Pow(x[0], 3) - 3*x[1] + 7*math.Pow(x[2], 3) - 2*x[3]
}

// DiffEvo invokes DE modules
func DiffEvo(setParameter Parameter) (
	minCostFinal float64,
	minIndividualFinal float64Data) {
	// '''
	// NumI   int         // Population size (number of individuals for each generation)
	// NumG   int         // Number of generations
	// CR     float64     // Crossover rate
	// F      float64     // Scalling factor
	// XLower float64Data // Lower bounds of variables
	// XUpper float64Data // Upper bounds of variables
	// '''
	ps := setParameter.NumI
	ng := setParameter.NumG
	cr := setParameter.CR
	f := setParameter.F
	xl := setParameter.XLower
	xu := setParameter.XUpper
	xall := initPop(setParameter)
	var minIndexFinal int
	for g := 0; g < ng; g++ {
		var minCost float64
		var minIndex int

		for i := 0; i < ps; i++ {
			mutatedIndividual := mutation(xall, i, f, xl, xu)
			var x float64Data
			x = xall[i]
			//xall[i] = Crossover2(xall[i], cr, mutatedIndividual)
			cost := crossover(&x, cr, mutatedIndividual)
			xall[i] = x
			if cost < minCost {
				minCost = cost
				minIndex = i
			}
		}
		minIndexFinal = minIndex
		minCostFinal = minCost
		fmt.Println(minIndex, minCost)
	}
	minIndividualFinal = xall[minIndexFinal]
	return minCostFinal, minIndividualFinal
}
