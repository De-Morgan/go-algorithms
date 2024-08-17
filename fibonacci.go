package main

/*Fibonacci Numbers

The Fibonacci numbers are named after a 13th century Italian mathematician known as Fibonacci.

The two first Fibonacci numbers are 0 and 1, and the next Fibonacci number is always the sum of the two previous numbers, so we get 0, 1, 1, 2, 3, 5, 8, 13, 21, ...

*/

//1. Implementation Using a For Loop

// fibonacciLoop returns the fibonacci sequence up to nth fibonacci using for loop
func fibonacciLoop(n int) []uint64 {
	fib := make([]uint64, 0, n)

	f0 := uint64(0)
	f1 := uint64(1)
	fib = append(fib, f0, f1)

	for i := 2; i < n; i++ {
		fn := f1 + f0
		fib = append(fib, fn)
		f1, f0 = fn, f1
	}

	return fib
}

//2. Implementation Using Recursion

// fibonacciRecursion returns the fibonacci sequence up to nth fibonacci using recursion

func fibonacciRecursion(n int) []uint64 {
	fib := make([]uint64, 0, n)

	for i := 0; i < n; i++ {
		fib = append(fib, fibonacciN(i))
	}

	return fib
}

//3. Finding The nth Fibonacci Number Using Recursion

// fibonacciN returns the nth Fibonacci Numbe
func fibonacciN(n int) uint64 {

	if n <= 1 {
		return uint64(n)
	}
	return fibonacciN(n-1) + fibonacciN(n-2)
}
