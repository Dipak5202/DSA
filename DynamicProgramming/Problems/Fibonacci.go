package dynamicprogramming

//Recursive
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

//Dynamic Programming
//memoization FibonacciMemoization calculates the nth Fibonacci number using memoization.
var memo = make([]int, 10)

// FibonacciMemoization calculates the nth Fibonacci number using memoization.
func FibonacciMemoization(n int) int {
	if n <= 1 {
		return n
	}

	if n < len(memo) && memo[n] != 0 {
		return memo[n]
	}

	memo[n] = FibonacciMemoization(n-1) + FibonacciMemoization(n-2)
	return memo[n]
}

// FibonacciTabulation calculates the nth Fibonacci number using tabulation.
func FibonacciTabulation(n int) int {
	if n <= 1 {
		return 1
	}

	tab := make([]int, n+1)
	// Initialize the base cases
	tab[0] = 0
	tab[1] = 1

	for i := 2; i <= n; i++ {
		tab[i] = tab[i-1] + tab[i-2]
	}

	return tab[n]
}

// FibonacciSpaceOptimized calculates the nth Fibonacci number using space optimization.
func FibonacciSpaceOptimized(n int) int {
	if n <= 1 {
		return 1
	}

	prev2, prev1 := 0, 1

	for i := 2; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}
