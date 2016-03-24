// Package main provides ...
package main

import "fmt"

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

	fmt.Printf("F\n")
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
	fmt.Printf("F'\n")
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
	fmt.Printf("B\n")
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
	fmt.Printf("B'\n")
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

	fmt.Printf("R\n")
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
	fmt.Printf("R'\n")
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

	fmt.Printf("L\n")
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
	fmt.Printf("L'\n")

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
	fmt.Printf("U\n")
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
	fmt.Printf("U'\n")
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
	fmt.Printf("D\n")
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
	fmt.Printf("D'\n")
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
