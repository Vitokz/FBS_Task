package handler

import (
	"context"
	"github.com/Vitokz/Task/Rest/models"
	"github.com/pkg/errors"
	"math"
	"strconv"
)
var StartMatrix = [][]int{
	{1,1},
	{1,0},
}

func (h *Handler) Fibonacci(from , to string,ctx context.Context) (*models.Response,error) {
	fromInt, err:=strconv.Atoi(from)
	if err != nil {
		return nil, errors.New("\"from\" param is not number")
	}else if fromInt < 0 {
		return nil, errors.New("\"from\" param is minus")
	}
	toInt, err := strconv.Atoi(to)
	if err != nil {
		return nil, errors.New("\"to\" param is not number")
	}else if toInt < 0 {
		return nil, errors.New("\"to\" param is minus")
	}else if toInt < fromInt {
		return nil, errors.New("\"to\" param is less than \"from\"")
	}

	result:=new(models.Response)
	for ;fromInt<=toInt;fromInt++ {
		val,err:=h.Db.GetValue(ctx,strconv.Itoa(fromInt))
		if err != nil {
			if fromInt == toInt {
				fib:=strconv.Itoa(calculateFibonacciNumber(fromInt))
				result.Numbers += fib
				if err:=h.Db.SetValue(ctx,strconv.Itoa(fromInt),fib);err != nil {
					return nil,err
				}
			} else {
				fib:=strconv.Itoa(calculateFibonacciNumber(fromInt))
				result.Numbers += fib + ", "
				if err:=h.Db.SetValue(ctx,strconv.Itoa(fromInt),fib);err != nil {
					return nil,err
				}
			}
		}else{
			if fromInt == toInt {
				result.Numbers += val
			} else {
				result.Numbers += val + ", "
			}
		}
	}
	return result,nil
}

func calculateFibonacciNumber(n int) int {
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

