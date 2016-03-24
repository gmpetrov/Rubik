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

const (
	FRONT = iota
	RIGHT
	LEFT
	TOP
	BACK
	BOTTOM
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
	length      int
	data        [][]int
	facesMap    map[int]int
	connections map[int]struct {
		name                     string
		left, right, top, bottom int
	}
}

type Edge struct {
	a, b int
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

	facesMap := map[int]int{
		0: BLUE,
		1: RED,
		2: ORANGE,
		3: YELLOW,
		4: GREEN,
		5: WHITE,
	}

	connections := map[int]struct {
		name                     string
		left, right, top, bottom int
	}{
		FRONT:  {"Front", LEFT, RIGHT, TOP, BOTTOM},
		BACK:   {"Back", RIGHT, LEFT, TOP, BOTTOM},
		LEFT:   {"Left", BACK, FRONT, TOP, BOTTOM},
		RIGHT:  {"Right", FRONT, BACK, TOP, BOTTOM},
		TOP:    {"Top", LEFT, RIGHT, BACK, FRONT},
		BOTTOM: {"Bottom", LEFT, RIGHT, FRONT, BACK},
	}

	return Cube{length, data, facesMap, connections}
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
	return cube.data[FRONT]
}

func (cube Cube) getBackFace() []int {
	return cube.data[BACK]
}

func (cube Cube) getRightFace() []int {
	return cube.data[RIGHT]
}

func (cube Cube) getLeftFace() []int {
	return cube.data[LEFT]
}

func (cube Cube) getTopFace() []int {
	return cube.data[TOP]
}

func (cube Cube) getBottomFace() []int {
	return cube.data[BOTTOM]
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

func (cube Cube) getFaceTopEdge(face []int) int {
	return cube.getFaceRow(face, 0)[1]
}

func (cube Cube) getFaceLeftEdge(face []int) int {
	return cube.getFaceRow(face, 1)[0]
}

func (cube Cube) getFaceRightEdge(face []int) int {
	return cube.getFaceRow(face, 1)[2]
}

func (cube Cube) getFaceBottomEdge(face []int) int {
	return cube.getFaceRow(face, 2)[1]
}

func (cube Cube) getFaceCenter(face []int) int {
	return cube.getFaceRow(face, 1)[1]
}

func (cube Cube) getTopFaceRightEdgeOpposite() int {
	return cube.getFaceRow(cube.getRightFace(), 0)[1]
}

func (cube Cube) getTopFaceLeftEdgeOpposite() int {
	return cube.getFaceRow(cube.getLeftFace(), 0)[1]
}

func (cube Cube) getTopFaceTopEdgeOpposite() int {
	return cube.getFaceRow(cube.getBackFace(), 0)[1]
}

func (cube Cube) getTopFaceBottomEdgeOpposite() int {
	return cube.getFaceRow(cube.getFrontFace(), 0)[1]
}

func (cube Cube) getFaceEdges(face []int) []int {
	edges := make([]int, 4)

	edges[0] = cube.getFaceTopEdge(face)
	edges[1] = cube.getFaceRightEdge(face)
	edges[2] = cube.getFaceBottomEdge(face)
	edges[3] = cube.getFaceLeftEdge(face)

	return edges
}

func (cube Cube) getRightSideEdges() []Edge {

	edges := make([]Edge, 4)

	faceEdges := cube.getFaceEdges(cube.getRightFace())
	sideEdges := make([]int, 4)

	sideEdges[0] = cube.getFaceRow(cube.getTopFace(), 1)[2]
	sideEdges[1] = cube.getFaceRow(cube.getBackFace(), 1)[0]
	sideEdges[2] = cube.getFaceRow(cube.getBottomFace(), 1)[2]
	sideEdges[3] = cube.getFaceRow(cube.getFrontFace(), 1)[2]

	for i := 0; i < 4; i++ {
		edges[i].a, edges[i].b = faceEdges[i], sideEdges[i]
	}
	return edges
}

func (cube Cube) getLeftSideEdges() []Edge {

	edges := make([]Edge, 4)

	faceEdges := cube.getFaceEdges(cube.getLeftFace())
	sideEdges := make([]int, 4)

	sideEdges[0] = cube.getFaceRow(cube.getTopFace(), 1)[0]
	sideEdges[1] = cube.getFaceRow(cube.getFrontFace(), 1)[2]
	sideEdges[2] = cube.getFaceRow(cube.getBottomFace(), 1)[0]
	sideEdges[3] = cube.getFaceRow(cube.getBackFace(), 1)[0]

	for i := 0; i < 4; i++ {
		edges[i].a, edges[i].b = faceEdges[i], sideEdges[i]
	}
	return edges
}

func (cube Cube) getFrontSideEdges() []Edge {

	edges := make([]Edge, 4)

	faceEdges := cube.getFaceEdges(cube.getLeftFace())
	sideEdges := make([]int, 4)

	sideEdges[0] = cube.getFaceRow(cube.getTopFace(), 2)[1]
	sideEdges[1] = cube.getFaceRow(cube.getRightFace(), 1)[0]
	sideEdges[2] = cube.getFaceRow(cube.getBottomFace(), 0)[1]
	sideEdges[3] = cube.getFaceRow(cube.getLeftFace(), 1)[2]

	for i := 0; i < 4; i++ {
		edges[i].a, edges[i].b = faceEdges[i], sideEdges[i]
	}
	return edges
}

func (cube Cube) getBackSideEdges() []Edge {

	edges := make([]Edge, 4)

	faceEdges := cube.getFaceEdges(cube.getLeftFace())
	sideEdges := make([]int, 4)

	sideEdges[0] = cube.getFaceRow(cube.getTopFace(), 0)[1]
	sideEdges[1] = cube.getFaceRow(cube.getLeftFace(), 1)[0]
	sideEdges[2] = cube.getFaceRow(cube.getBottomFace(), 2)[1]
	sideEdges[3] = cube.getFaceRow(cube.getRightFace(), 1)[2]

	for i := 0; i < 4; i++ {
		edges[i].a, edges[i].b = faceEdges[i], sideEdges[i]
	}
	return edges
}

func (cube Cube) printCube() {
	colorMap := getColorMap()

	fmt.Print("------------------------\n")
	for index, face := range cube.data {
		fmt.Printf("%s :", cube.connections[index].name)
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

func (cube Cube) Shuffle(movements []string) {

	movementsMap := cube.getMovementsMap()

	if DEBUG {
		fmt.Print("*** SHUFFLE ***\n")
		cube.printCube()
	}

	for _, value := range movements {
		movementsMap[value].(func())()
		if DEBUG {
			cube.printCube()
		}
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
