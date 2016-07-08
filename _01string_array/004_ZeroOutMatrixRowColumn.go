package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)
	ZeroOutMatrixRowColums(matrix)
	fmt.Println(matrix)

}

func ZeroOutMatrixRowColums(matrix [][]int) {
	rowMap := make(map[int]bool)
	colMap := make(map[int]bool)

	for row := range matrix {
		for col := range matrix[row] {
			if (matrix[row][col] == 0) {
				rowMap[row] = true
				colMap[col] = true
			}
		}
	}

	for i, _ := range rowMap {
		for j := range matrix[i] {
			matrix[i][j] = 0
		}
	}

	for j, _ := range colMap {
		for i := range matrix {
			matrix[i][j] = 0
		}
	}

}