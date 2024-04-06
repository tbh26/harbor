package main

import (
	"fmt"
	"github.com/tbh26/harbor/concurrent_go/manning_book/chapter_6/barrier"
	"math/rand"
)

// const matrixSize = 3
const (
	matrixSize     = 7
	middleRowIndex = 3
)

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
			if i == middleRowIndex {
				fmt.Printf("%3v  *  %3v  =>  %3v \n", matrixA[i], matrixB[i], result[i])
			} else {
				fmt.Printf("%3v     %3v      %3v \n", matrixA[i], matrixB[i], result[i])
			}
		}
		fmt.Println()
	}
}

////

/*
func generateRandMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] = rand.Intn(10) - 5
		}
	}
}
*/

func rowMultiply(matrixA, matrixB, result *[matrixSize][matrixSize]int,
	row int, barrier *barrier.Barrier) {
	for {
		barrier.Wait()
		for col := 0; col < matrixSize; col++ {
			sum := 0
			for i := 0; i < matrixSize; i++ {
				sum += matrixA[row][i] * matrixB[i][col]
			}
			result[row][col] = sum
		}
		barrier.Wait()
	}
}

func matrixBarrierDemo() {
	var matrixA, matrixB, result [matrixSize][matrixSize]int
	barrier := barrier.NewBarrier(matrixSize + 1)
	for row := 0; row < matrixSize; row++ {
		go rowMultiply(&matrixA, &matrixB, &result, row, barrier)
	}

	for i := 0; i < 4; i++ {
		generateRandMatrix(&matrixA)
		generateRandMatrix(&matrixB)
		barrier.Wait()
		barrier.Wait()
		for i := 0; i < matrixSize; i++ {
			//fmt.Println(matrixA[i], matrixB[i], result[i])
			if i == middleRowIndex {
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

	fmt.Println("\n\t\t\t\t\t=-=-=-=-=-=-=-=\n")

	matrixBarrierDemo()
}

func main() {
	fmt.Println(" =-=  matrix & barrier demo =-=")

	demo()
}
