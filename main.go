package main

import "fmt"

// MaxSumIS returns the largest possible sum of a strictly increasing
// subsequence. dp[i] is the best sum of an increasing subsequence ending
// exactly at i; each element looks back at all smaller predecessors.
func MaxSumIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	best := nums[0]
	for i, v := range nums {
		dp[i] = v
		for j := 0; j < i; j++ {
			if nums[j] < v && dp[j]+v > dp[i] {
				dp[i] = dp[j] + v
			}
		}
		if dp[i] > best {
			best = dp[i]
		}
	}
	return best
}

func main() {
	fmt.Println(MaxSumIS([]int{1, 101, 2, 3, 100, 4, 5}))
}
