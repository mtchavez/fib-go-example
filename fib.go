package main

import (
	"fmt"
	"os"
	"strconv"
)

type FibStruct struct {
	num    int
	result int
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func (f *FibStruct) Calculate(num int) {
	f.num = num
	f.result = fib(num)
}

func main() {
	num := 10
	if len(os.Args) > 1 {
		num, _ = strconv.Atoi(os.Args[1])
	}
	f := &FibStruct{}
	f.Calculate(num)
	fmt.Printf("Fibonaci of %v is %v\n", f.num, f.result)
}
