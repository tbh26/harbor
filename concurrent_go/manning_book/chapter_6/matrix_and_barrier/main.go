package main

import (
	"fmt"
	"math/rand"
)

const matrixSize = 3

func generateRandMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] = rand.Intn(10) - 5
		}
	}
}

func matrixMultiply(matrixA, matrixB, result *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			sum := 0
			for i := 0; i < matrixSize; i++ {
				sum += matrixA[row][i] * matrixB[i][col]
			}
			result[row][col] = sum
		}
	}
}

func matrixMultiplyDemo() {
	var matrixA, matrixB, result [matrixSize][matrixSize]int
	for i := 0; i < 4; i++ {
		generateRandMatrix(&matrixA)
		generateRandMatrix(&matrixB)
		matrixMultiply(&matrixA, &matrixB, &result)
		for i := 0; i < matrixSize; i++ {
			if i == 1 {
				fmt.Printf("%3v  *  %3v  =>  %3v \n", matrixA[i], matrixB[i], result[i])
			} else {
				fmt.Printf("%3v     %3v      %3v \n", matrixA[i], matrixB[i], result[i])
			}
		}
		fmt.Println()
	}
}

////

func demo() {
	matrixMultiplyDemo()
}

func main() {
	fmt.Println(" =-=  matrix & barrier demo =-=")

	demo()
}
