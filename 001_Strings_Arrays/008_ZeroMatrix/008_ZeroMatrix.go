package ch01

/**
 * Given a string, compress it into a readable format
 */

func zeroMatrix(mat [][]int) [][]int {
	n := len(mat[0])
	if n < 1 {
		return mat
	}

	zeroArr := make([]Coordinate, 0)

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				zeroArr = append(zeroArr, Coordinate {x: i, y: j})
			}
		}
	}

	for i:= 0; i < len(zeroArr); i++ {
		mat = zeroOutRowAndCol(mat, zeroArr[i])
	}

	return mat
}

func zeroOutRowAndCol(mat [][]int, coord Coordinate) [][]int {
	for i:=0; i < len(mat); i++ {
		mat[coord.x][i] *= 0
		mat[i][coord.y] *= 0
	}
	return mat
}

type Coordinate struct {
	x int
	y int
}