package fib

import (
	"fmt"
)

func Example() {
	var output int
	output = fib(10)
	fmt.Printf("Fibonacci of %v is %v\n", 10, output)
	output = fib(20)
	fmt.Printf("Fibonacci of %v is %v\n", 20, output)
	// Output:
	// Fibonacci of 10 is 55
	// Fibonacci of 20 is 6765
}

func ExampleFibStruct_Calculate() {
	f := FibStruct{}
	f.Calculate(10)
	fmt.Printf("Fibonacci of %v is %v\n", 10, f.result)

	f.Calculate(20)
	fmt.Printf("Fibonacci of %v is %v\n", 20, f.result)
	// Output:
	// Fibonacci of 10 is 55
	// Fibonacci of 20 is 6765
}
