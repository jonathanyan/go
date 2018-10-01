package main

import (
	"fmt"
)

func catMouseGame(graph [][]int) int {
	N := len(graph)
	dp := make([][][]int, N)
	for i := range graph {
		dp[i] = make([][]int, N)
		for j := range graph {
			dp[i][j] = make([]int, N+N)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	return whoWins(graph, 1, 2, 0, dp)
}

func whoWins(graph [][]int, mouse, cat, turn int, dp [][][]int) int {
	if turn >= 2*len(graph) {
		return 0
	}
	if dp[mouse][cat][turn] < 0 {
		if turn%2 == 0 {
			// mouse
			results := make([]int, 0, len(graph[mouse]))
			mouseWin := false
			for _, next := range graph[mouse] {
				if next == cat {
					continue
				}
				if next == 0 {
					mouseWin = true
					dp[mouse][cat][turn] = 1
					break
				}
				results = append(results, whoWins(graph, next, cat, turn+1, dp))
			}
			if !mouseWin {
				dp[mouse][cat][turn] = 2
				for _, r := range results {
					if r == 0 {
						dp[mouse][cat][turn] = 0
						break
					}
				}
				for _, r := range results {
					if r == 1 {
						dp[mouse][cat][turn] = 1
						break
					}
				}
			}
		} else {
			// cat
			results := make([]int, 0, len(graph[cat]))
			catWin := false
			for _, next := range graph[cat] {
				if next == 0 {
					continue
				}
				if next == mouse {
					catWin = true
					dp[mouse][cat][turn] = 2
					break
				}
				results = append(results, whoWins(graph, mouse, next, turn+1, dp))
			}
			if !catWin {
				dp[mouse][cat][turn] = 1
				for _, r := range results {
					if r == 0 {
						dp[mouse][cat][turn] = 0
						break
					}
				}
				for _, r := range results {
					if r == 2 {
						dp[mouse][cat][turn] = 2
						break
					}
				}
			}
		}
	}
	return dp[mouse][cat][turn]
}

func main() {
	fmt.Println("abc")
	graph := [][]int{{2, 5}, {3}, {0, 4, 5}, {1, 4, 5}, {2, 3}, {0, 2, 3}}
	a := catMouseGame(graph)
	fmt.Println(a)
}
