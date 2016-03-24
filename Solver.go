package main

import "fmt"

type Solver struct {
	moveCounter int
}
type getFaceEdgeFn func([]int) int
type getFaceOppositeFn func() int
type defaultMoveFn func()
type movePatternFn func()
type patternFn func(cube Cube)
type hasEdgesOfColorFn func(cube Cube, color int) bool

func (s Solver) PrintNumberMoves() {
	fmt.Printf("%v movement(s)\n", s.moveCounter)
}

func areEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (s Solver) isCrossAligned(cube Cube) bool {

	faces := make([][]int, 4)

	faces[0] = cube.getFrontFace()
	faces[1] = cube.getBackFace()
	faces[2] = cube.getLeftFace()
	faces[3] = cube.getRightFace()

	for i := 0; i < 4; i++ {
		column := cube.getFaceColumn(faces[i], 1)

		if column[0] != column[1] {
			return false
		}
	}
	return true
}

func (s Solver) hasCross(cube Cube) bool {

	row := cube.getFaceRow(cube.getTopFace(), 1)
	column := cube.getFaceColumn(cube.getTopFace(), 1)

	return areEqual(row, column)
}

func (s *Solver) EdgeSwapPattern(a, b, c, d, e movePatternFn) {
	a()
	b()
	c()
	d()
	e()
	s.moveCounter += 5
}

func (s *Solver) RightEdgeSwapPattern(cube Cube) {
	s.EdgeSwapPattern(
		cube.RightDouble,
		cube.DownPrime,
		cube.FrontPrime,
		cube.Right,
		cube.Front,
	)
}

func (s *Solver) LeftEdgeSwapPattern(cube Cube) {
	s.EdgeSwapPattern(
		cube.LeftDouble,
		cube.Down,
		cube.Front,
		cube.LeftPrime,
		cube.FrontPrime,
	)
}

func (s *Solver) TopEdgeSwapPattern(cube Cube) {
	s.EdgeSwapPattern(
		cube.BackDouble,
		cube.Down,
		cube.Left,
		cube.BackPrime,
		cube.LeftPrime,
	)
}

func (s *Solver) BottomEdgeSwapPattern(cube Cube) {
	s.EdgeSwapPattern(
		cube.FrontDouble,
		cube.Down,
		cube.Right,
		cube.FrontPrime,
		cube.RightPrime,
	)
}

func (s *Solver) makeCross(cube Cube) {
	for !s.hasCross(cube) {
		// s.ResolveTopFaceEdges(cube, cube.getFaceRightEdge, cube.getTopFaceRightEdgeOpposite, cube.Right, s.RightEdgeSwapPattern, s.hasRightSideEdgesOfColor)
		// s.ResolveTopFaceEdges(cube, cube.getFaceLeftEdge, cube.getTopFaceLeftEdgeOpposite, cube.Left, s.LeftEdgeSwapPattern, s.hasLeftSideEdgesOfColor)
		// s.ResolveTopFaceEdges(cube, cube.getFaceTopEdge, cube.getTopFaceTopEdgeOpposite, cube.Front, s.TopEdgeSwapPattern, s.hasFrontSideEdgesOfColor)
		// s.ResolveTopFaceEdges(cube, cube.getFaceBottomEdge, cube.getTopFaceBottomEdgeOpposite, cube.Back, s.BottomEdgeSwapPattern, s.hasBackSideEdgesOfColor)
	}
}

func (s *Solver) Resolve(cube Cube) {
	if DEBUG {
		fmt.Print("*** Resolve begin ***\n")
	}

	s.makeCross(cube)

	if DEBUG {
		fmt.Print("*** Resolve end   ***\n")
	}
}
