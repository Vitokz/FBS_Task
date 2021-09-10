package functions

var StartMatrix = [][]int{
	{1,1},
	{1,0},
}

func CalculateFibonacciNumber(n int) int {
	if n == 0 {
		return 0
	}
	а := 0
	b := 1
	for i:= 2; i < n; i++ {
		с := а + b
		а = b
		b = с
	}
	return а + b
}
