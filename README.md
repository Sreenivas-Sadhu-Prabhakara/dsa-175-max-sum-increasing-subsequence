# Maximum Sum Increasing Subsequence

> Striver SDE Sheet 175 · Day 25 DP — LIS's recurrence with the objective swapped: maximize the sum ending at each element, not the length. Go, stdlib only, tested.

## Problem

Among all strictly increasing subsequences, find the one with the largest *sum* — which may be short: `[1, 101, 2, 3, 100]`'s best is 1+2+3+100, but a huge early element can beat a long chain.

## Approach

The LIS skeleton (`dp[i] = best over j < i with nums[j] < nums[i]`) with one substitution: dp holds sums, seeded with the element's own value, and the answer is the max over all endings rather than the longest. The O(n log n) tails trick from #171 does **not** transfer — it relies on "longer is never worse", which breaks when the objective is a sum — so O(n²) is the honest complexity here. Seeing which optimizations survive a changed objective, and which silently don't, is the pedagogical payload of this problem.

**Complexity:** O(n²) time, O(n) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
