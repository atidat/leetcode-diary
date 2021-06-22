package main

import "fmt"

func main() {
	matrix := [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int)  {
	ml := len(matrix)-1
	for t := (ml)/2; t >= 0; t-- {
		i := t
		for j := i; j < ml-i; j++ {
			fmt.Println("==", matrix[i][j], matrix[j][ml-i], matrix[ml-i][ml-j], matrix[ml-j][j])
			matrix[j][ml-i], matrix[ml-i][ml-j], matrix[ml-j][i], matrix[i][j] =
				matrix[i][j], matrix[j][ml-i], matrix[ml-i][ml-j], matrix[ml-j][i]
			fmt.Println("--", matrix[i][j], matrix[j][ml-i], matrix[ml-i][ml-j], matrix[ml-j][j])
		}

	}
}

