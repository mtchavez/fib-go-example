package fib

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
