package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

// func objFunc(x, y, z){
// 	return 2x*x - y*y + 3*z
// }

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
	list := rand.Perm(25)
	for i, _ := range list {
		list[i]++
		fmt.Println(list[i])
	}

}
