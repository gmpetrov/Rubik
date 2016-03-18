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
