package main

import (
	"testing"
)

func TestRotate(t *testing.T) {
	cube := NewCube()

	cube.data[0] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	want := []int{7, 4, 1, 8, 5, 2, 9, 6, 3}

	cube.Rotate(0)

	for i := 0; i < len(cube.data[0]); i++ {
		if cube.data[0][i] != want[i] {
			t.Errorf("expected %v, got %v", want, cube.data[0])
		}
	}
}

func TestRotateInverse(t *testing.T) {
	cube := NewCube()

	cube.data[0] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	want := []int{3, 6, 9, 2, 5, 8, 1, 4, 7}

	cube.RotateInverse(0)

	for i := 0; i < len(cube.data[0]); i++ {
		if cube.data[0][i] != want[i] {
			t.Errorf("expected %v, got %v", want, cube.data[0])
		}
	}
}

func TestFront(t *testing.T) {
	cube := NewCube()

	want := [][]int{
		{BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE},
		{YELLOW, RED, RED, YELLOW, RED, RED, YELLOW, RED, RED},
		{ORANGE, ORANGE, WHITE, ORANGE, ORANGE, WHITE, ORANGE, ORANGE, WHITE},
		{YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, ORANGE, ORANGE, ORANGE},
		{GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN},
		{RED, RED, RED, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},
	}

	cube.Front()

	for index, face := range cube.data {
		for i := 0; i < len(cube.data); i++ {
			if want[index][i] != face[i] {
				t.Errorf("Front() expected %v, got %v", want[index], face)
			}
		}
	}

}

func TestFrontPrime(t *testing.T) {
	cube := NewCube()

	want := [][]int{
		{BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE},
		{WHITE, RED, RED, WHITE, RED, RED, WHITE, RED, RED},
		{ORANGE, ORANGE, YELLOW, ORANGE, ORANGE, YELLOW, ORANGE, ORANGE, YELLOW},
		{YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, RED, RED, RED},
		{GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN},
		{ORANGE, ORANGE, ORANGE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},
	}

	cube.FrontPrime()

	for index, face := range cube.data {
		for i := 0; i < len(cube.data); i++ {
			if want[index][i] != face[i] {
				t.Errorf("FrontPrime() expected %v, got %v", want, cube.data)
			}
		}
	}

}

func TestRight(t *testing.T) {
	cube := NewCube()

	want := [][]int{
		{BLUE, BLUE, WHITE, BLUE, BLUE, WHITE, BLUE, BLUE, WHITE},
		{RED, RED, RED, RED, RED, RED, RED, RED, RED},
		{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE},
		{YELLOW, YELLOW, BLUE, YELLOW, YELLOW, BLUE, YELLOW, YELLOW, BLUE},
		{YELLOW, GREEN, GREEN, YELLOW, GREEN, GREEN, YELLOW, GREEN, GREEN},
		{WHITE, WHITE, GREEN, WHITE, WHITE, GREEN, WHITE, WHITE, GREEN},
	}

	cube.Right()

	for index, face := range cube.data {
		for i := 0; i < len(cube.data); i++ {
			if want[index][i] != face[i] {
				t.Errorf("index = %v, i = %v", index, i)
				t.Errorf("%v != %v", want[index][i], face[i])
				t.Errorf("Right() expected %v, got %v", want[index], face)
			}
		}
	}

}

func TestRightPrime(t *testing.T) {
	cube := NewCube()

	want := [][]int{
		{BLUE, BLUE, YELLOW, BLUE, BLUE, YELLOW, BLUE, BLUE, YELLOW},
		{RED, RED, RED, RED, RED, RED, RED, RED, RED},
		{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE},
		{YELLOW, YELLOW, GREEN, YELLOW, YELLOW, GREEN, YELLOW, YELLOW, GREEN},
		{WHITE, GREEN, GREEN, WHITE, GREEN, GREEN, WHITE, GREEN, GREEN},
		{WHITE, WHITE, BLUE, WHITE, WHITE, BLUE, WHITE, WHITE, BLUE},
	}

	cube.RightPrime()

	for index, face := range cube.data {
		for i := 0; i < len(cube.data); i++ {
			if want[index][i] != face[i] {
				t.Errorf("index = %v, i = %v", index, i)
				t.Errorf("%v != %v", want[index][i], face[i])
				t.Errorf("RightPrime() expected %v, got %v", want[index], face)
			}
		}
	}

}

func TestLeft(t *testing.T) {
	cube := NewCube()

	want := [][]int{
		{YELLOW, BLUE, BLUE, YELLOW, BLUE, BLUE, YELLOW, BLUE, BLUE},
		{RED, RED, RED, RED, RED, RED, RED, RED, RED},
		{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE},
		{GREEN, YELLOW, YELLOW, GREEN, YELLOW, YELLOW, GREEN, YELLOW, YELLOW},
		{GREEN, GREEN, WHITE, GREEN, GREEN, WHITE, GREEN, GREEN, WHITE},
		{BLUE, WHITE, WHITE, BLUE, WHITE, WHITE, BLUE, WHITE, WHITE},
	}

	cube.Left()

	for index, face := range cube.data {
		for i := 0; i < len(cube.data); i++ {
			if want[index][i] != face[i] {
				t.Errorf("index = %v, i = %v", index, i)
				t.Errorf("%v != %v", want[index][i], face[i])
				t.Errorf("Left() expected %v, got %v", want[index], face)
			}
		}
	}

}

func TestLeftPrime(t *testing.T) {
	cube := NewCube()

	want := [][]int{
		{WHITE, BLUE, BLUE, WHITE, BLUE, BLUE, WHITE, BLUE, BLUE},
		{RED, RED, RED, RED, RED, RED, RED, RED, RED},
		{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE},
		{BLUE, YELLOW, YELLOW, BLUE, YELLOW, YELLOW, BLUE, YELLOW, YELLOW},
		{GREEN, GREEN, YELLOW, GREEN, GREEN, YELLOW, GREEN, GREEN, YELLOW},
		{GREEN, WHITE, WHITE, GREEN, WHITE, WHITE, GREEN, WHITE, WHITE},
	}

	cube.LeftPrime()

	for index, face := range cube.data {
		for i := 0; i < len(cube.data); i++ {
			if want[index][i] != face[i] {
				t.Errorf("index = %v, i = %v", index, i)
				t.Errorf("%v != %v", want[index][i], face[i])
				t.Errorf("LeftPrime() expected %v, got %v", want[index], face)
			}
		}
	}

}

func TestBack(t *testing.T) {
	cube := NewCube()

	want := [][]int{
		{BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE},
		{RED, RED, WHITE, RED, RED, WHITE, RED, RED, WHITE},
		{YELLOW, ORANGE, ORANGE, YELLOW, ORANGE, ORANGE, YELLOW, ORANGE, ORANGE},
		{RED, RED, RED, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW},
		{GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN},
		{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, ORANGE, ORANGE, ORANGE},
	}

	cube.Back()

	for index, face := range cube.data {
		for i := 0; i < len(cube.data); i++ {
			if want[index][i] != face[i] {
				t.Errorf("index = %v, i = %v", index, i)
				t.Errorf("%v != %v", want[index][i], face[i])
				t.Errorf("Back() expected %v, got %v", want[index], face)
			}
		}
	}
}

func TestBackPrime(t *testing.T) {
	cube := NewCube()

	want := [][]int{
		{BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE},
		{RED, RED, YELLOW, RED, RED, YELLOW, RED, RED, YELLOW},
		{WHITE, ORANGE, ORANGE, WHITE, ORANGE, ORANGE, WHITE, ORANGE, ORANGE},
		{ORANGE, ORANGE, ORANGE, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW},
		{GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN},
		{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, RED, RED, RED},
	}

	cube.BackPrime()

	for index, face := range cube.data {
		for i := 0; i < len(cube.data); i++ {
			if want[index][i] != face[i] {
				t.Errorf("index = %v, i = %v", index, i)
				t.Errorf("%v != %v", want[index][i], face[i])
				t.Errorf("BackPrime() expected %v, got %v", want[index], face)
			}
		}
	}
}

func TestUp(t *testing.T) {
	cube := NewCube()

	want := [][]int{
		{RED, RED, RED, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE},
		{GREEN, GREEN, GREEN, RED, RED, RED, RED, RED, RED},
		{BLUE, BLUE, BLUE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE},
		{YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW},
		{ORANGE, ORANGE, ORANGE, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN},
		{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, RED, RED, RED},
	}

	cube.Up()

	for index, face := range cube.data {
		for i := 0; i < len(cube.data); i++ {
			if want[index][i] != face[i] {
				t.Errorf("index = %v, i = %v", index, i)
				t.Errorf("%v != %v", want[index][i], face[i])
				t.Errorf("Up() expected %v, got %v", want[index], face)
			}
		}
	}
}

func TesUpPrime(t *testing.T) {
	cube := NewCube()

	want := [][]int{
		{RED, RED, RED, BLUE, BLUE, BLUE, BLUE, BLUE, BLUE},
		{GREEN, GREEN, GREEN, RED, RED, RED, RED, RED, RED},
		{BLUE, BLUE, BLUE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE},
		{YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW},
		{ORANGE, ORANGE, ORANGE, GREEN, GREEN, GREEN, GREEN, GREEN, GREEN},
		{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, RED, RED, RED},
	}

	cube.UpPrime()

	for index, face := range cube.data {
		for i := 0; i < len(cube.data); i++ {
			if want[index][i] != face[i] {
				t.Errorf("index = %v, i = %v", index, i)
				t.Errorf("%v != %v", want[index][i], face[i])
				t.Errorf("UpPrime() expected %v, got %v", want[index], face)
			}
		}
	}
}

func TestDown(t *testing.T) {
	cube := NewCube()

	want := [][]int{
		{BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, ORANGE, ORANGE, ORANGE},
		{RED, RED, RED, RED, RED, RED, BLUE, BLUE, BLUE},
		{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, GREEN, GREEN, GREEN},
		{YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW},
		{GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, RED, RED, RED},
		{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},
	}

	cube.Down()

	for index, face := range cube.data {
		for i := 0; i < len(cube.data); i++ {
			if want[index][i] != face[i] {
				t.Errorf("index = %v, i = %v", index, i)
				t.Errorf("%v != %v", want[index][i], face[i])
				t.Errorf("Down() expected %v, got %v", want[index], face)
			}
		}
	}
}

func TestDownPrime(t *testing.T) {
	cube := NewCube()

	want := [][]int{
		{BLUE, BLUE, BLUE, BLUE, BLUE, BLUE, RED, RED, RED},
		{RED, RED, RED, RED, RED, RED, GREEN, GREEN, GREEN},
		{ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, ORANGE, BLUE, BLUE, BLUE},
		{YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW, YELLOW},
		{GREEN, GREEN, GREEN, GREEN, GREEN, GREEN, ORANGE, ORANGE, ORANGE},
		{WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},
	}

	cube.DownPrime()

	for index, face := range cube.data {
		for i := 0; i < len(cube.data); i++ {
			if want[index][i] != face[i] {
				t.Errorf("index = %v, i = %v", index, i)
				t.Errorf("%v != %v", want[index][i], face[i])
				t.Errorf("DownPrime() expected %v, got %v", want[index], face)
			}
		}
	}
}
