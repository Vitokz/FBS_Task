package functions

import (
	"math"
	"strconv"
)

var StartMatrix = [][]int{
	{1,1},
	{1,0},
}

func CalculateFibonacciNumber(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	powers := powerOfTwo(n)

	matrices := make([][][]int, 0)
	for i := range powers {
		matrices = append(matrices, getMatrixPow(StartMatrix, powers[i]))
	}

	for len(matrices) >1 {
		matrix1:=matrices[len(matrices)-1]
		matrix2:=matrices[len(matrices)-2]
		matrices=matrices[:len(matrices)-2]
		matrices=append(matrices,multMatrix(matrix1,matrix2))
	}
	return matrices[0][0][0]
}

func powerOfTwo(n int) []int {
	bin := reverse(strconv.FormatInt(int64(n-1), 2))
	powers := make([]int, 0)
	for i := range bin {
		if bin[i:i+1] == "1" {
			powers = append(powers, int(math.Pow(2, float64(i))))
		}
	}
	return powers
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getMatrixPow(matrix [][]int, power int) [][]int {
	if power == 1 {
		return matrix
	}
	K := getMatrixPow(matrix, power/2)

	resultMatrix := multMatrix(K, K)
	return resultMatrix
}

func multMatrix(matrix1 [][]int, matrix2 [][]int) [][]int {
	dot11 := matrix1[0][0]*matrix2[0][0] + matrix1[0][1]*matrix2[1][0]
	dot12 := matrix1[0][0]*matrix2[0][1] + matrix1[0][1]*matrix2[1][1]
	dot21 := matrix1[1][0]*matrix2[0][0] + matrix1[1][1]*matrix2[1][0]
	dot22 := matrix2[1][0]*matrix2[0][1] + matrix1[1][1]*matrix2[1][1]
	return [][]int{
		{dot11, dot12},
		{dot21, dot22},
	}
}