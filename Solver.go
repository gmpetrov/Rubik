package main

type Solver struct {
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

func (s Solver) makeCross(cube Cube) {

}

func (s Solver) Simple(cube Cube) {

	s.makeCross(cube)

}
