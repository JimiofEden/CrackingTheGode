package ch01



/**
 * Given a string, compress it into a readable format
 */

func rotateMatrix(image [][]int) [][]int {
	if len(image[0]) < 2 {
		return image
	}

	n := len(image[0]);

	for i:=0; i < n/2; i++ {
		for j:=0; j < n-i-1; j++ {
			image[j][n-1-i], image[i][j], image[n-j-1][i], image[n-i-1][n-j-1] = image[i][j], image[n-1-j][i], image[n-i-1][n-j-1], image[j][n-i-1]
		}
	}

	return image
}