package dynamicprogramming

// ClimbingStairs calculates the number of distinct ways to climb n stairs,
// where you can take either 1 or 2 steps at a time.
func ClimbingStairs(n int) int {
	if n <= 1 {
		return 1
	}

	// Create a slice to store the number of ways to reach each step
	ways := make([]int, n+1)
	ways[0] = 1 // 1 way to stay at the ground (0 steps)
	ways[1] = 1 // 1 way to reach the first step (1 step)

	for i := 2; i <= n; i++ {
		// The number of ways to reach step i is the sum of the ways to reach
		// step i-1 (taking 1 step) and step i-2 (taking 2 steps)
		ways[i] = ways[i-1] + ways[i-2]
	}

	return ways[n]
}

// ClimbingStairsSpaceOptimized calculates the number of distinct ways to climb n stairs
// using space optimization, only keeping track of the last two steps.
func ClimbingStairsSpaceOptimized(n int) int {
	if n <= 2 {
		return n
	}

	prev2, prev1 := 1, 1 // ways to reach step 0 and step 1

	for i := 2; i <= n; i++ {
		current := prev1 + prev2 // current step is the sum of the last two steps
		prev2 = prev1            // update prev2 to be the previous step
		prev1 = current          // update prev1 to be the current step
	}

	return prev1 // return the number of ways to reach step n
}

// ClimbingStairsMemoization calculates the number of distinct ways to climb n stairs using memoization.
var memoStairs = make([]int, 100)

// ClimbingStairsMemoization calculates the number of distinct ways to climb n stairs using memoization.
func ClimbingStairsMemoization(n int) int {
	if n <= 1 {
		return 1
	}

	if n < len(memoStairs) && memoStairs[n] != 0 {
		return memoStairs[n]
	}

	memoStairs[n] = ClimbingStairsMemoization(n-1) + ClimbingStairsMemoization(n-2)
	return memoStairs[n]
}

// ClimbingStairsTabulation calculates the number of distinct ways to climb n stairs using tabulation.
func ClimbingStairsTabulation(n int) int {
	if n <= 1 {
		return 1
	}

	tab := make([]int, n+1)
	tab[0] = 1 // 1 way to stay at the ground (0 steps)
	tab[1] = 1 // 1 way to reach the first step (1 step)

	for i := 2; i <= n; i++ {
		tab[i] = tab[i-1] + tab[i-2] // sum of ways to reach the last two steps
	}

	return tab[n]
}
