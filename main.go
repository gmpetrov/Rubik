package main

import (
	"log"
	"os"
	"strings"
)

var DEBUG bool = true

func parse(str string) []string {
	str = strings.Trim(str, " ")
	str = strings.ToUpper(str)
	str = strings.Replace(str, "’", "'", -1)

	movements := strings.Split(str, " ")

	tmpCube := NewCube()

	cubeMovements := tmpCube.getMovementsMap()

	for _, value := range movements {
		if _, ok := cubeMovements[value]; !ok {
			log.Fatalf("Argument not valid : %s", value)
		}
	}
	return movements
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		log.Fatal("[Usage] : ./rubik \"F R U2 B' L' D\"")
	}

	movements := parse(args[0])

	cube := NewCube()

	cube.Shuffle(movements)

	solver := Solver{}

	solver.Resolve(cube)

	solver.PrintNumberMoves()

	cube.printCube()
}
