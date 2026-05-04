package main

func rotate(matrix [][]int)  {
	arr := make([][]int, len(matrix))
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {

		for j := 0; j < len(matrix); j++ {
			arr[j][len(matrix) - i - 1] = matrix[i][j]
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[i][j] = arr[i][j]
		}
	}
}