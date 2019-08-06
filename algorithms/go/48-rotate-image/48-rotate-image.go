package image

/*
 * clockwise rotate
 * first reverse up to down, then swap the symmetry
 * 1 2 3     7 8 9     7 4 1
 * 4 5 6  => 4 5 6  => 8 5 2
 * 7 8 9     1 2 3     9 6 3
*/
func rotate(matrix [][]int) {
	reverse(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return
}

func reverse(matrix [][]int) {
	lm := len(matrix)
	for i := 0; i < lm / 2; i++ {
		matrix[i], matrix[lm-1-i] = matrix[lm-1-i], matrix[i]
	}
}

/*
 * anticlockwise rotate
 * first reverse left to right, then swap the symmetry
 * 1 2 3     3 2 1     3 6 9
 * 4 5 6  => 6 5 4  => 2 5 8
 * 7 8 9     9 8 7     1 4 7
*/
func anti_rotate(matrix [][]int)  {
	anti_reverse(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return
}

func anti_reverse(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		lm := len(matrix[i])
		for j := 0; j < lm / 2; j++ {
			matrix[i][j], matrix[i][lm-1-j] = matrix[i][lm-1-j], matrix[i][j]
		}
	}
}
