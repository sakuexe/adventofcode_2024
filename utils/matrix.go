package utils

import "math"

// rotates the matrix in place
// https://github.com/amarjitdhillon/CP_2021/tree/master/48-rotate-image
func RotateMatrix90Deg(matrix [][]rune) [][]rune {

  // first transformation is for all the rows and half of the elements lying to right hand side fo diagonal
  for r, _ := range matrix {
    for c := r; c < len(matrix[0]); c++ {
      matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
    }
  }

  // second transformation for all the remaining rows
  for r, _ := range matrix {
    // looks weird, but does the same as // division in python
    for c := r; c < int(math.Floor(float64(len(matrix[0])) / 2)); c++ {
      matrix[r][c], matrix[r][len(matrix[0]) - c - 1] = matrix[r][len(matrix[0]) - c - 1], matrix[r][c]
    }
  }

  return matrix
}

// rotates the matrix in place
func RotateMatrix45Deg(matrix [][]rune) [][]rune {

  counter := 0
  for counter < 2 * len(matrix[0]) - 1 {
  }

  return matrix
}
