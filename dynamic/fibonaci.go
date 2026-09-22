package dynamic

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	// Create an array to store Fibonacci numbers up to n
	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1

	// Calculate Fibonacci numbers iteratively
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}