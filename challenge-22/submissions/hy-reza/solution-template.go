package main

import (
	"fmt"
	"math"
)

func main() {
	// Standard U.S. coin denominations in cents
	denominations := []int{1, 5, 10, 25, 50}

	// Test amounts
	amounts := []int{87, 42, 99, 33, 7}

	for _, amount := range amounts {
		// Find minimum number of coins
		minCoins := MinCoins(amount, denominations)

		// Find coin combination
		coinCombo := CoinCombination(amount, denominations)

		// Print results
		fmt.Printf("Amount: %d cents\n", amount)
		fmt.Printf("Minimum coins needed: %d\n", minCoins)
		fmt.Printf("Coin combination: %v\n", coinCombo)
		fmt.Println("---------------------------")
	}
}

// MinCoins returns the minimum number of coins needed to make the given amount.
// If the amount cannot be made with the given denominations, return -1.
func MinCoins(amount int, denominations []int) int {
	// dp[i] will store the minimum number of coins needed to make amount i.
	dp := make([]int, amount+1)

	// Initialize dp array. dp[0] is 0 as no coins are needed for amount 0.
	// All other amounts are initialized to a large value (infinity) to represent
	// that they are not yet reachable.
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32 // Use MaxInt32 to represent infinity
	}
	dp[0] = 0

	// Iterate through each amount from 1 to the target amount.
	for i := 1; i <= amount; i++ {
		// For each amount, iterate through all available denominations.
		for _, coin := range denominations {
			// If the current amount `i` is greater than or equal to the coin denomination,
			// and if `dp[i-coin]` is not infinity (meaning `i-coin` is reachable),
			// then we can potentially use this coin to reach amount `i`.
			if i >= coin && dp[i-coin] != math.MaxInt32 {
				// Update dp[i] with the minimum of its current value and (1 + dp[i-coin]).
				// 1 represents the current coin being used.
				dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i-coin])))
			}
		}
	}

	// If dp[amount] is still infinity, it means the amount cannot be made with the given denominations.
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// CoinCombination returns a map with the specific combination of coins that gives
// the minimum number. The keys are coin denominations and values are the number of
// coins used for each denomination.
// If the amount cannot be made with the given denominations, return an empty map.
func CoinCombination(amount int, denominations []int) map[int]int {
	// dp[i] will store the minimum number of coins needed to make amount i.
	dp := make([]int, amount+1)
	// prevCoin[i] will store the last coin used to reach amount i with the minimum number of coins.
	prevCoin := make([]int, amount+1)

	// Initialize dp array. dp[0] is 0.
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	// Iterate through each amount from 1 to the target amount.
	for i := 1; i <= amount; i++ {
		// For each amount, iterate through all available denominations.
		for _, coin := range denominations {
			// If the current amount `i` is greater than or equal to the coin denomination,
			// and if `dp[i-coin]` is not infinity,
			// and if using this coin results in a smaller number of coins for amount `i`.
			if i >= coin && dp[i-coin] != math.MaxInt32 && 1+dp[i-coin] < dp[i] {
				// Update dp[i] with the new minimum.
				dp[i] = 1 + dp[i-coin]
				// Record the coin that was used to achieve this minimum.
				prevCoin[i] = coin
			}
		}
	}

	// If dp[amount] is still infinity, the amount cannot be made. Return an empty map.
	if dp[amount] == math.MaxInt32 {
		return make(map[int]int)
	}

	// Reconstruct the coin combination by backtracking using the prevCoin array.
	coinCount := make(map[int]int)
	currentAmount := amount
	for currentAmount > 0 {
		coin := prevCoin[currentAmount]
		coinCount[coin]++
		currentAmount -= coin
	}

	return coinCount
}
