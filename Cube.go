// Package cube provides ...
package main

import "fmt"

const (
	RED = iota
	BLUE
	ORANGE
	YELLOW
	GREEN
	WHITE
)

func printBlue() {
	fmt.Print("\033[1;44m  \033[0m")
}

func printGreen() {
	fmt.Print("\033[1;42m  \033[0m")
}

func printRed() {
	fmt.Print("\033[1;41m  \033[0m")
}

func printOrange() {
	fmt.Print("\033[1;45m  \033[0m")
}

func printYellow() {
	fmt.Print("\033[1;43m  \033[0m")
}

func printWhite() {
	fmt.Print("\033[1;47m  \033[0m")
}

func getColorMap() map[int]interface{} {

	colorMap := map[int]interface{}{
		BLUE:   printBlue,
		RED:    printRed,
		ORANGE: printOrange,
		YELLOW: printYellow,
		GREEN:  printGreen,
		WHITE:  printWhite,
	}

	return colorMap
}

type Cube struct {
	length    int
	data      [][]int
	facesName map[int]string
	facesMap  map[int]int
}

func (cube Cube) getMovementsMap() map[string]interface{} {
	movementsMap := map[string]interface{}{
		"F":  cube.Front,
		"F'": cube.FrontPrime,
		"B":  cube.Back,
		"B'": cube.BackPrime,
		"R":  cube.Right,
		"R'": cube.RightPrime,
		"L":  cube.Left,
		"L'": cube.LeftPrime,
		"U":  cube.Up,
		"U'": cube.UpPrime,
		"D":  cube.Down,
		"D'": cube.DownPrime,
		"F2": cube.FrontDouble,
		"B2": cube.BackDouble,
		"R2": cube.RightDouble,
		"L2": cube.LeftDouble,
		"U2": cube.UpDouble,
		"D2": cube.DownDouble,
	}
	return movementsMap
}

func NewCube() Cube {

	length := 3

	// 3x3x3 Rubik's cube
	data := [][]int{
		{BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE},                   // 0 Front
		{RED, RED, RED, RED, RED, RED, RED, RED, RED},                            // 1 Right
		{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE}, // 2 Left
		{YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW}, // 3 Top
		{GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN},          // 4 Back
		{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},          // 5 Bottom
	}

	facesName := map[int]string{
		0: "Front",
		1: "Right",
		2: "Left",
		3: "Top",
		4: "Back",
		5: "Bottom",
	}

	facesMap := map[int]int{
		0: BLUE,
		1: RED,
		2: ORANGE,
		3: YELLOW,
		4: GREEN,
		5: WHITE,
	}

	return Cube{length, data, facesName, facesMap}
}

func (cube Cube) Rotate(index int) {
	slice := cube.data[index]

	tmp := make([]int, len(slice))

	cubeRowLength := len(slice) / 3

	for i := 0; i < len(slice); i++ {
		y := i / cubeRowLength
		x := i % cubeRowLength

		xPrime := 2 - y
		yPrime := x

		tmp[xPrime+yPrime*cubeRowLength] = slice[x+y*cubeRowLength]
	}
	cube.data[index] = tmp
}

func (cube Cube) RotateInverse(index int) {
	slice := cube.data[index]

	tmp := make([]int, len(slice))

	cubeRowLength := len(slice) / 3

	for i := 0; i < len(slice); i++ {
		y := i / cubeRowLength
		x := i % cubeRowLength

		xPrime := y
		yPrime := 2 - x

		tmp[xPrime+yPrime*cubeRowLength] = slice[x+y*cubeRowLength]
	}
	cube.data[index] = tmp

}

func (cube Cube) getFrontFace() []int {
	return cube.data[0]
}

func (cube Cube) getBackFace() []int {
	return cube.data[4]
}

func (cube Cube) getRightFace() []int {
	return cube.data[1]
}

func (cube Cube) getLeftFace() []int {
	return cube.data[2]
}

func (cube Cube) getTopFace() []int {
	return cube.data[3]
}

func (cube Cube) getBottomFace() []int {
	return cube.data[5]
}

func (cube Cube) getFaceColumn(face []int, index int) []int {
	length := len(face) / cube.length
	values := make([]int, length)

	for i := 0; i < length; i++ {
		values[i] = face[index+cube.length*i]
	}
	return values
}

func (cube Cube) getFaceRow(face []int, index int) []int {
	length := len(face) / cube.length
	values := make([]int, length)

	for i := 0; i < length; i++ {
		values[i] = face[index*cube.length+i]
	}
	return values
}

func (cube Cube) setFaceColumn(face []int, index int, values []int) {
	for i, value := range values {
		face[index+cube.length*i] = value
	}
}

func (cube Cube) setFaceRow(face []int, index int, values []int) {
	for i, value := range values {
		face[index*cube.length+i] = value
	}
}

func (cube Cube) printCube() {
	colorMap := getColorMap()

	fmt.Print("------------------------\n")
	for index, face := range cube.data {
		fmt.Printf("%s :", cube.facesName[index])
		for i := 0; i < len(face); i++ {
			if i%3 == 0 {
				fmt.Print("\n\n")
			}
			colorMap[face[i]].(func())()
			fmt.Print(" ")
		}
		fmt.Print("\n\n")
	}
	fmt.Print("------------------------\n")
}

func (cube Cube) Front() {
	cube.Rotate(0)

	leftFace := cube.getLeftFace()
	bottomFace := cube.getBottomFace()
	rightFace := cube.getRightFace()
	topFace := cube.getTopFace()

	tmp := cube.getFaceColumn(leftFace, 2)

	cube.setFaceColumn(leftFace, 2, cube.getFaceRow(bottomFace, 0))
	cube.setFaceRow(bottomFace, 0, cube.getFaceColumn(rightFace, 0))
	cube.setFaceColumn(rightFace, 0, cube.getFaceRow(topFace, 2))
	cube.setFaceRow(topFace, 2, tmp)
}

func (cube Cube) FrontPrime() {
	cube.RotateInverse(0)

	leftFace := cube.getLeftFace()
	bottomFace := cube.getBottomFace()
	rightFace := cube.getRightFace()
	topFace := cube.getTopFace()

	tmp := cube.getFaceColumn(leftFace, 2)

	cube.setFaceColumn(leftFace, 2, cube.getFaceRow(topFace, 2))
	cube.setFaceRow(topFace, 2, cube.getFaceColumn(rightFace, 0))
	cube.setFaceColumn(rightFace, 0, cube.getFaceRow(bottomFace, 0))
	cube.setFaceRow(bottomFace, 0, tmp)
}

func (cube Cube) Back() {
	cube.Rotate(4)

	leftFace := cube.getLeftFace()
	bottomFace := cube.getBottomFace()
	rightFace := cube.getRightFace()
	topFace := cube.getTopFace()

	tmp := cube.getFaceColumn(leftFace, 0)

	cube.setFaceColumn(leftFace, 0, cube.getFaceRow(topFace, 0))
	cube.setFaceRow(topFace, 0, cube.getFaceColumn(rightFace, 2))
	cube.setFaceColumn(rightFace, 2, cube.getFaceRow(bottomFace, 2))
	cube.setFaceRow(bottomFace, 2, tmp)
}

func (cube Cube) BackPrime() {
	cube.RotateInverse(4)

	leftFace := cube.getLeftFace()
	bottomFace := cube.getBottomFace()
	rightFace := cube.getRightFace()
	topFace := cube.getTopFace()

	tmp := cube.getFaceColumn(leftFace, 0)

	cube.setFaceColumn(leftFace, 0, cube.getFaceRow(bottomFace, 2))
	cube.setFaceRow(bottomFace, 2, cube.getFaceColumn(rightFace, 2))
	cube.setFaceColumn(rightFace, 2, cube.getFaceRow(topFace, 0))
	cube.setFaceRow(topFace, 0, tmp)
}

func (cube Cube) Right() {
	cube.Rotate(1)

	backFace := cube.getBackFace()
	bottomFace := cube.getBottomFace()
	frontFace := cube.getFrontFace()
	topFace := cube.getTopFace()

	tmp := cube.getFaceColumn(frontFace, 2)

	cube.setFaceColumn(frontFace, 2, cube.getFaceColumn(bottomFace, 2))
	cube.setFaceColumn(bottomFace, 2, cube.getFaceColumn(backFace, 0))
	cube.setFaceColumn(backFace, 0, cube.getFaceColumn(topFace, 2))
	cube.setFaceColumn(topFace, 2, tmp)
}

func (cube Cube) RightPrime() {
	cube.RotateInverse(1)

	backFace := cube.getBackFace()
	bottomFace := cube.getBottomFace()
	frontFace := cube.getFrontFace()
	topFace := cube.getTopFace()

	tmp := cube.getFaceColumn(frontFace, 2)

	cube.setFaceColumn(frontFace, 2, cube.getFaceColumn(topFace, 2))
	cube.setFaceColumn(topFace, 2, cube.getFaceColumn(backFace, 0))
	cube.setFaceColumn(backFace, 0, cube.getFaceColumn(bottomFace, 2))
	cube.setFaceColumn(bottomFace, 2, tmp)
}

func (cube Cube) Left() {
	cube.Rotate(2)

	backFace := cube.getBackFace()
	bottomFace := cube.getBottomFace()
	frontFace := cube.getFrontFace()
	topFace := cube.getTopFace()

	tmp := cube.getFaceColumn(frontFace, 0)

	cube.setFaceColumn(frontFace, 0, cube.getFaceColumn(topFace, 0))
	cube.setFaceColumn(topFace, 0, cube.getFaceColumn(backFace, 2))
	cube.setFaceColumn(backFace, 2, cube.getFaceColumn(bottomFace, 0))
	cube.setFaceColumn(bottomFace, 0, tmp)

}

func (cube Cube) LeftPrime() {
	cube.RotateInverse(2)

	backFace := cube.getBackFace()
	bottomFace := cube.getBottomFace()
	frontFace := cube.getFrontFace()
	topFace := cube.getTopFace()

	tmp := cube.getFaceColumn(frontFace, 0)

	cube.setFaceColumn(frontFace, 0, cube.getFaceColumn(bottomFace, 0))
	cube.setFaceColumn(bottomFace, 0, cube.getFaceColumn(backFace, 2))
	cube.setFaceColumn(backFace, 2, cube.getFaceColumn(topFace, 0))
	cube.setFaceColumn(topFace, 0, tmp)

}

func (cube Cube) Up() {
	cube.Rotate(3)

	backFace := cube.getBackFace()
	leftFace := cube.getLeftFace()
	frontFace := cube.getFrontFace()
	rightFace := cube.getRightFace()

	tmp := cube.getFaceRow(frontFace, 0)

	cube.setFaceRow(frontFace, 0, cube.getFaceRow(rightFace, 0))
	cube.setFaceRow(rightFace, 0, cube.getFaceRow(backFace, 0))
	cube.setFaceRow(backFace, 0, cube.getFaceRow(leftFace, 0))
	cube.setFaceRow(leftFace, 0, tmp)
}

func (cube Cube) UpPrime() {
	cube.RotateInverse(3)

	backFace := cube.getBackFace()
	leftFace := cube.getLeftFace()
	frontFace := cube.getFrontFace()
	rightFace := cube.getRightFace()

	tmp := cube.getFaceRow(frontFace, 0)

	cube.setFaceRow(frontFace, 0, cube.getFaceRow(leftFace, 0))
	cube.setFaceRow(leftFace, 0, cube.getFaceRow(backFace, 0))
	cube.setFaceRow(backFace, 0, cube.getFaceRow(rightFace, 0))
	cube.setFaceRow(rightFace, 0, tmp)
}

func (cube Cube) Down() {
	cube.Rotate(5)

	backFace := cube.getBackFace()
	leftFace := cube.getLeftFace()
	frontFace := cube.getFrontFace()
	rightFace := cube.getRightFace()

	tmp := cube.getFaceRow(frontFace, 2)

	cube.setFaceRow(frontFace, 2, cube.getFaceRow(leftFace, 2))
	cube.setFaceRow(leftFace, 2, cube.getFaceRow(backFace, 2))
	cube.setFaceRow(backFace, 2, cube.getFaceRow(rightFace, 2))
	cube.setFaceRow(rightFace, 2, tmp)
}

func (cube Cube) DownPrime() {
	cube.RotateInverse(5)

	backFace := cube.getBackFace()
	leftFace := cube.getLeftFace()
	frontFace := cube.getFrontFace()
	rightFace := cube.getRightFace()

	tmp := cube.getFaceRow(frontFace, 2)

	cube.setFaceRow(frontFace, 2, cube.getFaceRow(rightFace, 2))
	cube.setFaceRow(rightFace, 2, cube.getFaceRow(backFace, 2))
	cube.setFaceRow(backFace, 2, cube.getFaceRow(leftFace, 2))
	cube.setFaceRow(leftFace, 2, tmp)
}

func (cube Cube) FrontDouble() {
	cube.Front()
	cube.Front()
}

func (cube Cube) BackDouble() {
	cube.Back()
	cube.Back()
}

func (cube Cube) RightDouble() {
	cube.Right()
	cube.Right()
}

func (cube Cube) LeftDouble() {
	cube.Left()
	cube.Left()
}

func (cube Cube) UpDouble() {
	cube.Up()
	cube.Up()
}

func (cube Cube) DownDouble() {
	cube.Down()
	cube.Down()
}

func (cube Cube) Shuffle(movements []string) {

	movementsMap := cube.getMovementsMap()

	for _, value := range movements {
		if DEBUG {
			cube.printCube()
		}
		movementsMap[value].(func())()
	}
}

func (cube Cube) isSolved() bool {
	for index, face := range cube.data {
		for i := 0; i < len(face); i++ {
			if face[i] != cube.facesMap[index] {
				return false
			}
		}
	}
	return true
}
