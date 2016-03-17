package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	cube := NewCube()

	cube.Front()

	cube.printCube()

	fmt.Printf("cube.data[0] = %v\n", cube.data[0])

	// // initialize a 3 element float64 slice
	// dx := make([]float64, 3)

	// // set the elements
	// dx[0] = 2
	// dx[1] = 2
	// dx[2] = 3

	// // pass the slice dx as data to the matrix x
	// x := mat64.NewDense(3, 1, dx)

	// fmt.Println(x)

	if len(args) != 1 {
		log.Fatal("[Usage] : ./rubik \"F R U2 B' L' D\"")
	}
}
