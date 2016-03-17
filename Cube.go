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

type Cube struct {
	length    int
	data      [][]int
	facesName map[int]string
}

func NewCube() Cube {

	length := 3

	// 3x3 Rubik's cube
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

	return Cube{length, data, facesName}
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
	for index, face := range cube.data {
		fmt.Printf("%s\n%v\n%v\n%v\n\n", cube.facesName[index], face[:3], face[3:6], face[6:])
	}
}

func (cube Cube) Front() {
	cube.Rotate(0)

	leftFace := cube.getLeftFace()
	bottomFace := cube.getBottomFace()
	rightFace := cube.getRightFace()
	topFace := cube.getTopFace()

	tmp := cube.getFaceColumn(leftFace, 2)

	fmt.Printf("tmp = %v\n", tmp)

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
