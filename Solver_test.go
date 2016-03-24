package main

import (
	"testing"
)

func TestHasCross(t *testing.T) {
	cube := NewCube()
	solver := Solver{}

	if !solver.hasCross(cube) {
		t.Error("hasCross() : Should be true")
	}

	cube.Right()

	if solver.hasCross(cube) {
		t.Error("hasCross() : Should be false")
	}
}

func TestIsCrossAligned(t *testing.T) {
	cube := NewCube()
	solver := Solver{}

	if !solver.isCrossAligned(cube) {
		t.Error("isCrossAligned() : Should be true")
	}

	cube.Up()

	if solver.isCrossAligned(cube) {
		t.Error("isCrossAligned() : Should be false")
	}
}

func TestRightSwapPattern(t *testing.T) {
	cube := NewCube()

	topFace := cube.getTopFace()
	rightFace := cube.getRightFace()

	topFace[5], rightFace[1] = rightFace[1], topFace[5]

	s := Solver{}

	s.RightEdgeSwapPattern(cube)

	topFace = cube.getTopFace()
	rightFace = cube.getRightFace()

	got1 := topFace[5]
	got2 := rightFace[1]

	if got1 != YELLOW || got2 != RED {
		t.Errorf("RightSwapPattern(): got1 = %v, got2 = %v", got1, got2)
	}

}

func TestLeftSwapPattern(t *testing.T) {
	cube := NewCube()

	topFace := cube.getTopFace()
	leftFace := cube.getLeftFace()

	topFace[3], leftFace[1] = leftFace[1], topFace[3]

	s := Solver{}

	s.LeftEdgeSwapPattern(cube)

	topFace = cube.getTopFace()
	leftFace = cube.getLeftFace()

	got1 := topFace[3]
	got2 := leftFace[1]

	if got1 != YELLOW || got2 != ORANGE {
		t.Errorf("LeftSwapPattern(): got1 = %v, got2 = %v", got1, got2)
	}

}

func TestTopSwapPattern(t *testing.T) {
	cube := NewCube()

	topFace := cube.getTopFace()
	backFace := cube.getBackFace()

	topFace[1], backFace[1] = backFace[1], topFace[1]

	s := Solver{}

	s.TopEdgeSwapPattern(cube)

	topFace = cube.getTopFace()
	backFace = cube.getBackFace()

	got1 := topFace[1]
	got2 := backFace[1]

	if got1 != YELLOW || got2 != GREEN {
		t.Errorf("TopSwapPattern(): got1 = %v, got2 = %v", got1, got2)
	}

}

func TestBottomSwapPattern(t *testing.T) {
	cube := NewCube()

	topFace := cube.getTopFace()
	frontFace := cube.getFrontFace()

	topFace[7], frontFace[1] = frontFace[1], topFace[7]

	s := Solver{}

	s.BottomEdgeSwapPattern(cube)

	topFace = cube.getTopFace()
	frontFace = cube.getFrontFace()

	got1 := topFace[7]
	got2 := frontFace[1]

	if got1 != YELLOW || got2 != BLUE {
		t.Errorf("BottomSwapPattern(): got1 = %v, got2 = %v", got1, got2)
	}

}
