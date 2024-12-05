package utils

// https://www.geeksforgeeks.org/rotate-a-matrix-by-90-degree-in-clockwise-direction-without-using-any-extra-space/
func RotateMatrix90Deg(originalMatrix [][]rune) [][]rune {

  // make a deep copy
  matrix := make([][]rune, len(originalMatrix))
  for i := range originalMatrix {
    matrix[i] = make([]rune, len(originalMatrix[i]))
  }

  for i := range len(originalMatrix) {
    for j := range len(originalMatrix[i]) {
      matrix[j][len(originalMatrix) - i - 1] = originalMatrix[i][j]
    }
  }

  return matrix
}

func RotateMatrix45Deg(matrix [][]rune) [][]rune {
  rows := len(matrix)
  cols := len(matrix[0])
  rotatedMatrix := [][]rune{}

  // walk through the matrix diagonally
  // https://www.geeksforgeeks.org/zigzag-or-diagonal-traversal-of-matrix/
  for line := range rows + cols {

    rotatedMatrix = append(rotatedMatrix, []rune{})
    start_col := max(0, line - rows)
    elements := min(line, (cols - start_col), rows)

    for j := range elements {
      character := matrix[min(rows, line) - j - 1][start_col + j]
      rotatedMatrix[line] = append(rotatedMatrix[line], character)
    }
  }

  return rotatedMatrix
}
