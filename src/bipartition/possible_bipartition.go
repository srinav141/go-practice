/*

Given a set of N people (numbered 1, 2, ..., N), we would like to split everyone into two groups of any size.

Each person may dislike some other people, and they should not go into the same group.

Formally, if dislikes[i] = [a, b], it means it is not allowed to put the people numbered a and b into the same group.

Return true if and only if it is possible to split everyone into two groups in this way.



Example 1:

Input: N = 4, dislikes = [[1,2],[1,3],[2,4]]
Output: true
Explanation: group1 [1,4], group2 [2,3]
Example 2:

Input: N = 3, dislikes = [[1,2],[1,3],[2,3]]
Output: false
Example 3:

Input: N = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
Output: false


Note:

1 <= N <= 2000
0 <= dislikes.length <= 10000
1 <= dislikes[i][j] <= N
dislikes[i][0] < dislikes[i][1]
There does not exist i != j for which dislikes[i] == dislikes[j].


*/

package main

import "fmt"

func possibleBipartition(N int, dislikes [][]int) bool {

	con := make([][]int, N+1)
	fmt.Println(con)
	for i := 0; i < len(dislikes); i++ {
		con[dislikes[i][0]] = append(con[dislikes[i][0]], dislikes[i][1])
		con[dislikes[i][1]] = append(con[dislikes[i][1]], dislikes[i][0])
	}
	fmt.Println(con)

	seen := make(map[int]int)
	for i := 1; i <= N; i++ {
		if _, ok := seen[i]; !ok {
			if !checkSeen(con, i, 1, seen) {
				return false
			}

		}

	}

	return true
}

func checkSeen(con [][]int, node, seenValue int, seen map[int]int) bool {
fmt.Println(seen)
	if v, ok := seen[node]; ok {
		if v == seenValue {
			return true
		}
		return false
	}
	seen[node] = seenValue
	vert := con[node]
	for _, v := range vert {
		if !checkSeen(con, v, -seenValue, seen) {
			return false
		}

	}
	return true
}

func main() {

	N := 3
	dislikes := [][]int{{4, 7}, {4, 8}, {5, 6}, {1, 6}, {3, 7}, {2, 5}, {5, 8}, {1, 2}, {4, 9}, {6, 10}, {8, 10}, {3, 6}, {2, 10}, {9, 10}, {3, 8}, {2, 3}, {1, 9}, {4, 6}, {5, 7}, {3, 8}, {1, 8}, {1, 7}, {2, 4}}
	dislikes = [][]int{{1, 2}, {1, 3}, {2, 3}}
	res := possibleBipartition(N, dislikes)
	fmt.Println(res)

}
