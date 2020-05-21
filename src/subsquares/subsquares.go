/*
Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.



Example 1:

Input: matrix =
[
  [0,1,1,1],
  [1,1,1,1],
  [0,1,1,1]
]
Output: 15
Explanation:
There are 10 squares of side 1.
There are 4 squares of side 2.
There is  1 square of side 3.
Total number of squares = 10 + 4 + 1 = 15.
Example 2:

Input: matrix =
[
  [1,0,1],
  [1,1,0],
  [1,1,0]
]
Output: 7
Explanation:
There are 6 squares of side 1.
There is 1 square of side 2.
Total number of squares = 6 + 1 = 7.
*/
package main

import "fmt"

func countSquares(matrix [][]int) int {

	count := 0
	m := len(matrix)
	n := len(matrix[0])

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {

			if matrix[i][j] == 0 {
				continue
			}

			matrix[i][j] = min(min(matrix[i-1][j], matrix[i][j-1]), matrix[i-1][j-1]) + 1
		}

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count += matrix[i][j]
		}
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	a := [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}
	a = [][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}

	res := countSquares(a)
	fmt.Println(res)

}
