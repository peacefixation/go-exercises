## diff

Compare two files to find lines that differ.

Use the [LCS](https://en.wikipedia.org/wiki/Longest_common_subsequence_problem) dynamic programming algorithm to find the difference between the two files.

This algorithm breaks the problem (files with a lot of lines) down into smaller sub-problems (individual lines) and populates a matrix with the length of the longest common subsequence at each stage of the diff. Tracing a path through the matrix from the end to the start (backtracking) reveals the solution.

The recursive approach is potentially too memory intensive to use on large files but has the benefit of naturally printing the backtrack in reverse order. 

An iterative approach is also included but uses extra memory to store the diff so it can be printed in reverse order.