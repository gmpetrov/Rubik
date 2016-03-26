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

const (
	CORNER_TOP_LEFT = iota
	EDGE_TOP
	CORNER_TOP_RIGHT
	EDGE_LEFT
	CENTER
	EDGE_RIGHT
	CORNER_BOTTOM_LEFT
	EDGE_BOTTOM
	CORNER_BOTTOM_RIGHT
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
		name                           string
		left, right, top, bottom, back int
	}
}

type Edge struct {
	face, a, b int
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
		name                           string
		left, right, top, bottom, back int
	}{
		FRONT:  {"Front", LEFT, RIGHT, TOP, BOTTOM, BACK},
		BACK:   {"Back", RIGHT, LEFT, TOP, BOTTOM, FRONT},
		LEFT:   {"Left", BACK, FRONT, TOP, BOTTOM, RIGHT},
		RIGHT:  {"Right", FRONT, BACK, TOP, BOTTOM, LEFT},
		TOP:    {"Top", LEFT, RIGHT, BACK, FRONT, BOTTOM},
		BOTTOM: {"Bottom", LEFT, RIGHT, FRONT, BACK, TOP},
	}

	return Cube{length, data, facesMap, connections}
}

func (cube Cube) getConnectionsHelper() map[int]struct{ rot, rotPrime interface{} } {
	helper := map[int]struct {
		rot, rotPrime interface{}
	}{
		FRONT:  {cube.Front, cube.FrontPrime},
		BACK:   {cube.Back, cube.BackPrime},
		LEFT:   {cube.Left, cube.LeftPrime},
		RIGHT:  {cube.Right, cube.RightPrime},
		TOP:    {cube.Up, cube.UpPrime},
		BOTTOM: {cube.Down, cube.DownPrime},
	}
	return helper
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

func (cube Cube) getFaceId(face []int) int {
	return face[CENTER]
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

func (cube Cube) RunSequence(movements []string) {

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

func isCorner(pos int) bool {
	return pos == CORNER_TOP_LEFT ||
		pos == CORNER_TOP_RIGHT ||
		pos == CORNER_BOTTOM_LEFT ||
		pos == CORNER_BOTTOM_RIGHT
}

func isEdge(pos int) bool {
	return pos == EDGE_TOP ||
		pos == EDGE_LEFT ||
		pos == EDGE_RIGHT ||
		pos == EDGE_BOTTOM
}

func (cube Cube) MoveFacetDown(facePos, facet int) {
	connectionsHelper := cube.getConnectionsHelper()

	if facet <= CORNER_TOP_RIGHT {
		connectionsHelper[facePos].rot.(func())()
		connectionsHelper[facePos].rot.(func())()
	} else if facet <= EDGE_RIGHT {
		if facet == EDGE_LEFT {
			connectionsHelper[facePos].rotPrime.(func())()
		} else if facet == EDGE_RIGHT {
			connectionsHelper[facePos].rot.(func())()
		}
	}
}

// func (cube Cube) TransfertFacet(from, fromPos, to, toPos int) {

// 	connectionsHelper := cube.getConnectionsHelper()

// 	cube.MoveFacetDown(from, fromPos)

// }
