package lcs

import "fmt"

// ComputeLCS compute the LCS length using dynamic programming algorithm
// return the LCS length matrix that can be backtracked to determine diff of xs and ys
func ComputeLCS(xs, ys []string) [][]int {
	c := make([][]int, len(xs)+1)
	for i := 0; i < len(c); i++ {
		c[i] = make([]int, len(ys)+1)
	}

	for i := 0; i <= len(xs); i++ {
		for j := 0; j <= len(ys); j++ {
			if i == 0 || j == 0 { // first (extra) row/column init to 0
				c[i][j] = 0
			} else if xs[i-1] == ys[j-1] {
				c[i][j] = c[i-1][j-1] + 1
			} else {
				c[i][j] = max(c[i][j-1], c[i-1][j])
			}
		}
	}

	return c
}

// PrintLCS print the LCS matrix for inspection
func PrintLCS(c [][]int) {
	for i, rows := range c {
		for j := range rows {
			fmt.Printf("%4d", c[i][j])
		}
		fmt.Println()
	}
}

// max find the max of 2 integers
func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

// PrintDiffRecursive print the diff of xs and ys using recursion
// naturally prints in reverse order when call stack unwinds (desirable, as we are backtracking from last to first)
// i and j should be initially given as len(xs) and len(ys)
func PrintDiffRecursive(c [][]int, xs, ys []string, i, j int) {
	if i > 0 && j > 0 && xs[i-1] == ys[j-1] {
		PrintDiffRecursive(c, xs, ys, i-1, j-1)
		fmt.Printf("  %s\n", xs[i-1])
	} else if j > 0 && (i == 0 || c[i][j-1] > c[i-1][j]) {
		PrintDiffRecursive(c, xs, ys, i, j-1)
		fmt.Printf("+ %s\n", ys[j-1])
	} else if i > 0 && (j == 0 || c[i][j-1] <= c[i-1][j]) {
		PrintDiffRecursive(c, xs, ys, i-1, j)
		fmt.Printf("- %s\n", xs[i-1])
	}
}

// PrintDiffIterative print the diff of xs and ys using iteration
// store diff in memory to print in reverse order
func PrintDiffIterative(c [][]int, xs, ys []string) {
	i := len(xs)
	j := len(ys)

	// store the strings so they can be printed in reverse order
	strs := make([]string, 0)

	for i > 0 && j > 0 {
		if xs[i-1] == ys[j-1] {
			strs = append(strs, fmt.Sprintf("  %s", xs[i-1]))
			i--
			j--
		} else if c[i-1][j] > c[i][j-1] {
			strs = append(strs, fmt.Sprintf("- %s", xs[i-1]))
			i--
		} else {
			strs = append(strs, fmt.Sprintf("+ %s", ys[j-1]))
			j--
		}
	}

	// print the strings in reverse order
	for i := len(strs) - 1; i >= 0; i-- {
		fmt.Println(strs[i])
	}
}
