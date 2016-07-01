package main

import "fmt"


func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	PrintMatrix(matrix)
	RotateMatrix(matrix, 3)
	fmt.Println("Rotated matrix:")
	PrintMatrix(matrix)

}

func RotateMatrix(matrix [][]int, row int) {
	for layer := 0; layer < row / 2; layer++ {
		first := layer
		last := row - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first

			top := matrix[first][i]

			matrix[first][i] = matrix[last - offset][first]
			matrix[last - offset][first] = matrix[last][last - offset]
			matrix[last][last - offset] = matrix[i][last]
			matrix[i][last] = top
		}
	}
}

func PrintMatrix(matrix [][]int)  {
	for i:= range matrix {
		fmt.Println(matrix[i])
	}
}
