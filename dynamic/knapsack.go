package dynamic

func Knapsack(weights []int, values []int, capacity int) int {
	n := len(weights)
	// Create a 2D array to store the maximum value at each n and capacity
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// Build the dp array
	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if weights[i-1] <= w {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[n][capacity]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}